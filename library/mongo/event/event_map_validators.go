package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo/service"
)

type MapValidatorsEvent struct{}

var btcMinerService = service.BtcMinerService{}
var validatorRelationService = service.ValidatorRelationService{}

func (mapValidatorsEvent MapValidatorsEvent) handle(b blockInfo, args map[string]interface{}) error {
	btcMinerStr, ok := args["btcAddresses"]
	if !ok {
		return mapValidatorsEvent.checkArg("btcAddresses")
	}
	btcMinerSlice := (*btcMinerStr.(*interface{})).([]string)

	validators, ok := args["asimovAddresses"]
	if !ok {
		return mapValidatorsEvent.checkArg("asimovAddresses")
	}
	validatorSlice := make([]string, 0)
	for _, v := range *validators.(*[]asimovCommon.Address) {
		validatorSlice = append(validatorSlice, v.Hex())
	}

	domains, ok := args["domains"]
	if !ok {
		return mapValidatorsEvent.checkArg("domains")
	}
	domainSlice := (*domains.(*interface{})).([]string)

	if len(btcMinerSlice) != len(validatorSlice) || len(validatorSlice) != len(domainSlice) {
		return errors.New("invalid parameter, length of three parameters should be equivalent")
	}

	for i := 0; i < len(btcMinerSlice); i++ {
		// Save BTC Miner
		err := btcMinerService.Insert(b.height, b.time, btcMinerSlice[i], domainSlice[i])
		if err != nil {
			return err
		}

		// Save Validator
		err = validatorService.Insert(b.height, b.time, validatorSlice[i])
		if err != nil {
			return err
		}

		// Save validator relation
		err = validatorRelationService.Insert(b.height, b.time, btcMinerSlice[i], validatorSlice[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (mapValidatorsEvent MapValidatorsEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("map validators event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
