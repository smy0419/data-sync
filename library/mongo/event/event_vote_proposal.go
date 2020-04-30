package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"math/big"
)

type VoteProposalEvent struct{}

func (voteProposalEvent VoteProposalEvent) handle(b blockInfo, args map[string]interface{}) error {
	voters, ok := args["voters"]
	if !ok {
		return voteProposalEvent.checkArg("voters")
	}
	votersSlice := make([]string, 0)
	for _, v := range *voters.(*[]asimovCommon.Address) {
		votersSlice = append(votersSlice, v.Hex())
	}

	proposalId, ok := args["proposalId"]
	if !ok {
		return voteProposalEvent.checkArg("proposalId")
	}

	return foundationTodoListService.Insert(
		b.height,
		b.time,
		votersSlice,
		(*proposalId.(**big.Int)).Int64(),
	)
}

func (voteProposalEvent VoteProposalEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("voting proposal event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
