package event

import (
	"errors"
	"fmt"
	"github.com/AsimovNetwork/asimov/chaincfg"
	"github.com/AsimovNetwork/data-sync/library/common"
)

type InitGenesisOrganization struct{}

func (initGenesisOrganization InitGenesisOrganization) handle(b blockInfo, args map[string]interface{}) error {
	// get asimov network
	network := common.Cfg.GetAsimovNet()
	members, ok := chaincfg.NetConstructorArgsMap[network]["genesisCitizens"]
	if !ok {
		return initGenesisOrganization.checkArg("genesisCitizens")
	}
	memberSlice := make([]string, 0)
	for _, v := range members {
		memberSlice = append(memberSlice, v.Hex())
	}

	return foundationMemberService.Insert(b.height, b.time, memberSlice)
}

func (initGenesisOrganization InitGenesisOrganization) checkArg(arg string) error {
	errMsg := fmt.Sprintf("init genesis organization event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
