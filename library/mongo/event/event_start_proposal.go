package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"math/big"
)

type StartProposalEvent struct{}

func (startProposalEvent StartProposalEvent) handle(b blockInfo, args map[string]interface{}) error {
	proposalId, ok := args["proposalId"]
	if !ok {
		return startProposalEvent.checkArg("proposalId")
	}

	proposalType, ok := args["proposalType"]
	if !ok {
		return startProposalEvent.checkArg("proposalType")
	}

	proposer, ok := args["proposer"]
	if !ok {
		return startProposalEvent.checkArg("proposer")
	}

	endTime, ok := args["endTime"]
	if !ok {
		return startProposalEvent.checkArg("endTime")
	}

	return foundationProposalService.Insert(
		b.height,
		b.time,
		(*proposalId.(**big.Int)).Int64(),
		*proposalType.(*uint8),
		proposer.(*asimovCommon.Address).Hex(),
		(*endTime.(**big.Int)).Int64(),
		b.txHash,
	)
}

func (startProposalEvent StartProposalEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("start proposal event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
