package event

import (
	"errors"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/asimov/crypto"
	"github.com/AsimovNetwork/asimov/rpcs/rpcjson"
	"github.com/AsimovNetwork/asimov/vm/fvm"
	"github.com/AsimovNetwork/asimov/vm/fvm/abi"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo/service"
	"math/big"
	"strings"
)

type blockInfo struct {
	height int64
	time   int64
	txHash string
}

type EventLog interface {
	handle(b blockInfo, args map[string]interface{}) error
}

var evenLogHandler = map[string]EventLog{
	// validator begin
	"InitValidator":      InitValidatorEvent{},
	"MapValidatorsEvent": MapValidatorsEvent{},
	// validator end

	// foundation event begin
	"InitGenesisOrganization":   InitGenesisOrganization{},
	"ReceiveEvent":              DonateEvent{},
	"TransferAssetEvent":        TransferAssetEvent{},
	"StartProposalEvent":        StartProposalEvent{},
	"VoteEvent":                 VoteEvent{},
	"VoteProposalEvent":         VoteProposalEvent{},
	"ProposalStatusChangeEvent": ProposalStatusChangeEvent{},
	"ElectCitizenEvent":         ElectCitizenEvent{},
	"ImpeachCitizenEvent":       ImpeachCitizenEvent{},
	// foundation event end

	// miner event begin
	"InitValidatorCommittee":              InitValidatorCommittee{},
	"UpdateRoundBlockInfoEvent":           UpdateRoundBlockInfoEvent{},
	"ProposalVotersEvent":                 ProposalVotersEvent{},
	"CommitteeProposalStatusChangeEvent":  CommitteeProposalStatusChangeEvent{},
	"NewRoundEvent":                       NewRoundEvent{},
	"SignupCommitteeEvent":                SignupCommitteeEvent{},
	"StartCommitteeProposalEvent":         StartCommitteeProposalEvent{},
	"VoteForCommitteeEvent":               VoteForCommitteeEvent{},
	"MultiAssetProposalEffectHeightEvent": MultiAssetProposalEffectHeightEvent{},
	// miner event end

	// dao event begin
	"CreateVoteContract":      CreateVoteContract{},
	"CloseOrganizationEvent":  CloseOrganizationEvent{},
	"RenameOrganizationEvent": RenameOrganizationEvent{},
	"RemoveMemberEvent":       RemoveMemberEvent{},

	"CreateAssetEvent":     CreateAssetEvent{},
	"TransferSuccessEvent": TransferSuccessEvent{},

	"InviteNewPresident":  InviteNewPresident{},
	"ConfirmNewPresident": ConfirmNewPresident{},
	"InviteNewMember":     InviteNewMember{},
	"JoinNewMember":       JoinNewMember{},

	"StartVoteEvent":        StartVoteEvent{},
	"DAOVoteEvent":          DAOVoteEvent{},
	"VoteStatusChangeEvent": VoteStatusChangeEvent{},
	"MintAssetEvent":        MintAssetEvent{},
	// dao event end

	"ContractTemplateEvent": ContractTemplateEvent{},
}

var validatorService = service.ValidatorService{}
var earningService = service.EarningService{}
var systemContractService = service.SystemContractService{}
var foundationBalanceSheetService = service.FoundationBalanceSheetService{}
var foundationMemberService = service.FoundationMemberService{}
var foundationProposalService = service.FoundationProposalService{}
var foundationVoteService = service.FoundationVoteService{}
var foundationTodoListService = service.FoundationTodoListService{}
var minerTodoListService = service.MinerTodoListService{}

// miner service start
var minerSignUpService = service.MinerSignUpService{}
var minerProposalService = service.MinerProposalService{}
var minerVoteService = service.MinerVoteService{}
var minerRoundService = service.MinerRoundService{}
var minerMemberService = service.MinerMemberService{}
var blockService = service.BlockService{}

type EventLogService struct{}

func (eventLogService EventLogService) HandleEventLog(height int64, time int64, receipts []*rpcjson.ReceiptResult) error {
	if height == 0 {
		// Get genesis block
		remoteBlock, err := blockService.FetchBlocks(0, 1)
		if err != nil {
			return err
		}

		// Initial validator
		initValidator, ok := evenLogHandler["InitValidator"]
		if !ok {
			errMsg := "InitValidator handler not found"
			common.Logger.Error(errMsg)
			return errors.New(errMsg)
		}
		b := blockInfo{
			height: height,
			time:   time,
			txHash: remoteBlock[0].RawTx[0].Hash,
		}
		err = initValidator.handle(b, nil)
		if err != nil {
			common.Logger.Error("InitValidator handler run failed")
			return err
		}

		// Initial genesis organization
		initGenesisOrganization, ok := evenLogHandler["InitGenesisOrganization"]
		if !ok {
			errMsg := "InitGenesisOrganization handler not found"
			common.Logger.Error(errMsg)
			return errors.New(errMsg)
		}
		err = initGenesisOrganization.handle(b, nil)
		if err != nil {
			common.Logger.Error("InitGenesisOrganization handler run failed")
			return err
		}

		// initial validator committee
		initValidatorCommittee, ok := evenLogHandler["InitValidatorCommittee"]
		if !ok {
			errMsg := "InitValidatorCommittee handler not found"
			common.Logger.Error(errMsg)
			return errors.New(errMsg)
		}
		err = initValidatorCommittee.handle(b, nil)
		if err != nil {
			common.Logger.Error("InitValidatorCommittee handler run failed")
			return err
		}

		return nil
	} else {
		for _, receipt := range receipts {
			for _, log := range receipt.Logs {
				// Contract address
				address := log.Address
				// The 0th topic matches the method signature
				topic := log.Topics[0]
				// RLP encoded Log value
				data := log.Data

				// Get contract address via ABI
				exist, abiStr, err := systemContractService.GetSystemContractAbi(height, address)
				if err != nil {
					return err
				}
				if !exist {
					exist, abiStr, err = systemContractService.GetDaoContractAbiByAddress(address)
					if err != nil {
						return err
					}
					if !exist {
						break
					}
				}

				// Parsing log via ABI
				definition, err := abi.JSON(strings.NewReader(abiStr))
				events := definition.Events
				for _, event := range events {
					signature := make([]string, 0)
					inputName := make([]string, 0)
					signature = append(signature, event.Name, "(")
					for _, input := range event.Inputs {
						signature = append(signature, input.Type.String(), ",")
						inputName = append(inputName, input.Name)
					}
					// if len(event.Inputs) == 0 {
					// 	fmt.Println(">>>>>>>>>>>>>>>>>>>>")
					// 	fmt.Println(event.Name)
					// }
					if len(event.Inputs) > 0 {
						signature = signature[:len(signature)-1]
					}
					signature = append(signature, ")")
					sum := crypto.Keccak256([]byte(strings.Join(signature, "")))
					hash := asimovCommon.BytesToHash(sum)
					if hash.String() == topic {
						// fmt.Println("#############################")
						// fmt.Println(signature)
						result, err := fvm.UnpackEvent(abiStr, event.Name, asimovCommon.Hex2Bytes(data))
						if err != nil {
							panic(err)
						}
						args := make(map[string]interface{})

						switch result.(type) {
						case []interface{}:
							resultSlice := result.([]interface{})
							for i := 0; i < len(inputName); i++ {
								args[inputName[i]] = TypeConvert(resultSlice[i])
								// fmt.Println(fmt.Sprintf("key %s, value %v", inputName[i], resultSlice[i]))
							}
						default:
							if len(inputName) > 0 {
								args[inputName[0]] = TypeConvert(result)
							}
						}

						handler, ok := evenLogHandler[event.Name]
						if !ok {
							errMsg := "log handler not found"
							common.Logger.Error(errMsg)
							// return errors.New(errMsg)
							break
						}
						b := blockInfo{
							height: height,
							time:   time,
							txHash: log.TxHash,
						}
						err = handler.handle(b, args)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}

// Type Conversion
func TypeConvert(param interface{}) interface{} {
	switch param.(type) {
	case bool:
		result := param.(bool)
		return &result
	case string:
		result := param.(string)
		return &result
	case uint8:
		result := param.(uint8)
		return &result
	case uint16:
		result := param.(uint16)
		return &result
	case uint32:
		result := param.(uint32)
		return &result
	case uint64:
		result := param.(uint64)
		return &result
	case int:
		result := param.(int)
		return &result
	case int8:
		result := param.(int8)
		return &result
	case int16:
		result := param.(int16)
		return &result
	case int32:
		result := param.(int32)
		return &result
	case int64:
		result := param.(int64)
		return &result
	case asimovCommon.Address:
		result := param.(asimovCommon.Address)
		return &result
	case [4]byte:
		result := param.([4]byte)
		return &result
	case [32]byte:
		result := param.([32]byte)
		return &result
	case *big.Int:
		result := param.(*big.Int)
		return &result
	case []uint8:
		result := param.([]uint8)
		return &result
	case []uint16:
		result := param.([]uint16)
		return &result
	case []uint32:
		result := param.([]uint32)
		return &result
	case []uint64:
		result := param.([]uint64)
		return &result
	case []*big.Int:
		result := param.([]*big.Int)
		return &result
	case []int8:
		result := param.([]int8)
		return &result
	case []int16:
		result := param.([]int16)
		return &result
	case []int32:
		result := param.([]int32)
		return &result
	case []int64:
		result := param.([]int64)
		return &result
	case []asimovCommon.Address:
		result := param.([]asimovCommon.Address)
		return &result
	default:
		return &param
	}
}
