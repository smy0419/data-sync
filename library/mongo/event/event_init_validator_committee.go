package event

import (
	"errors"
	"fmt"
	"github.com/AsimovNetwork/asimov/chaincfg"
	"github.com/AsimovNetwork/data-sync/library/common"
)

type InitValidatorCommittee struct{}

func (initValidatorCommittee InitValidatorCommittee) handle(b blockInfo, args map[string]interface{}) error {
	var round int64 = 1
	network := common.Cfg.GetAsimovNet()
	validators, ok := chaincfg.NetConstructorArgsMap[network]["_validators"]
	if !ok {
		return initValidatorCommittee.checkArg("_validators")
	}
	validatorSlice := make([]string, 0)
	for _, v := range validators {
		validatorSlice = append(validatorSlice, v.Hex())
	}

	startTime := b.time
	endTime := startTime + 30*24*60*60

	err := minerMemberService.Insert(b.height, b.time, round, validatorSlice)
	if err != nil {
		return err
	}

	return minerRoundService.Insert(b.height, b.time, round, startTime, endTime)
}

func (initValidatorCommittee InitValidatorCommittee) checkArg(arg string) error {
	errMsg := fmt.Sprintf("init validator committee event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
