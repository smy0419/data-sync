package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"math/big"
)

type NewRoundEvent struct{}

func (newRoundEvent NewRoundEvent) handle(b blockInfo, args map[string]interface{}) error {
	round, ok := args["round"]
	if !ok {
		return newRoundEvent.checkArg("round")
	}

	startTime, ok := args["startTime"]
	if !ok {
		return newRoundEvent.checkArg("startTime")
	}

	endTime, ok := args["endTime"]
	if !ok {
		return newRoundEvent.checkArg("endTime")
	}

	validators, ok := args["validators"]
	if !ok {
		return newRoundEvent.checkArg("validators")
	}
	validatorSlice := make([]string, 0)
	for _, v := range *validators.(*[]asimovCommon.Address) {
		validatorSlice = append(validatorSlice, v.Hex())
	}

	err := minerMemberService.Insert(b.height, b.time, (*round.(**big.Int)).Int64(), validatorSlice)
	if err != nil {
		return err
	}

	return minerRoundService.Insert(
		b.height,
		b.time,
		(*round.(**big.Int)).Int64(),
		(*startTime.(**big.Int)).Int64(),
		(*endTime.(**big.Int)).Int64(),
	)
}

func (newRoundEvent NewRoundEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("miner new round event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
