package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"math/big"
)

type SignupCommitteeEvent struct{}

func (signupCommitteeEvent SignupCommitteeEvent) handle(b blockInfo, args map[string]interface{}) error {
	round, ok := args["round"]
	if !ok {
		return signupCommitteeEvent.checkArg("round")
	}

	validator, ok := args["validator"]
	if !ok {
		return signupCommitteeEvent.checkArg("validator")
	}

	return minerSignUpService.Insert(
		b.height,
		b.time,
		b.txHash,
		(*round.(**big.Int)).Int64(),
		validator.(*asimovCommon.Address).Hex(),
	)
}

func (signupCommitteeEvent SignupCommitteeEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("miner sign up event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
