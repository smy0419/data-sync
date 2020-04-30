package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"math/big"
)

type StartCommitteeProposalEvent struct{}

func (startCommitteeProposalEvent StartCommitteeProposalEvent) handle(b blockInfo, args map[string]interface{}) error {
	round, ok := args["round"]
	if !ok {
		return startCommitteeProposalEvent.checkArg("round")
	}

	proposalId, ok := args["proposalId"]
	if !ok {
		return startCommitteeProposalEvent.checkArg("proposalId")
	}

	proposer, ok := args["proposer"]
	if !ok {
		return startCommitteeProposalEvent.checkArg("proposer")
	}

	proposalType, ok := args["proposalType"]
	if !ok {
		return startCommitteeProposalEvent.checkArg("proposalType")
	}

	status, ok := args["status"]
	if !ok {
		return startCommitteeProposalEvent.checkArg("status")
	}

	endTime, ok := args["endTime"]
	if !ok {
		return startCommitteeProposalEvent.checkArg("endTime")
	}

	return minerProposalService.Insert(
		(*round.(**big.Int)).Int64(),
		b.height,
		b.time,
		(*endTime.(**big.Int)).Int64(),
		(*proposalId.(**big.Int)).Int64(),
		proposer.(*asimovCommon.Address).Hex(),
		*proposalType.(*uint8),
		*status.(*uint8),
		b.txHash,
	)
}

func (startCommitteeProposalEvent StartCommitteeProposalEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("miner start committee proposal event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
