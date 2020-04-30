package event

import (
	"errors"
	"fmt"
	"github.com/AsimovNetwork/data-sync/library/common"
	"math/big"
)

type CommitteeProposalStatusChangeEvent struct{}

func (committeeProposalStatusChangeEvent CommitteeProposalStatusChangeEvent) handle(b blockInfo, args map[string]interface{}) error {
	round, ok := args["round"]
	if !ok {
		return committeeProposalStatusChangeEvent.checkArg("round")
	}

	proposalId, ok := args["proposalId"]
	if !ok {
		return committeeProposalStatusChangeEvent.checkArg("proposalId")
	}

	status, ok := args["status"]
	if !ok {
		return committeeProposalStatusChangeEvent.checkArg("status")
	}

	supportRate, ok := args["supportRate"]
	if !ok {
		return committeeProposalStatusChangeEvent.checkArg("supportRate")
	}

	rejectRate, ok := args["rejectRate"]
	if !ok {
		return committeeProposalStatusChangeEvent.checkArg("rejectRate")
	}

	return minerProposalService.ChangeStatus(b.height, b.time, (*round.(**big.Int)).Int64(), (*proposalId.(**big.Int)).Int64(), *status.(*uint8), (*supportRate.(**big.Int)).Int64(), (*rejectRate.(**big.Int)).Int64())
}

func (committeeProposalStatusChangeEvent CommitteeProposalStatusChangeEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("miner proposal status change event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
