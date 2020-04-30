package event

import (
	"errors"
	"fmt"
	"github.com/AsimovNetwork/data-sync/library/common"
	"math/big"

	asimovCommon "github.com/AsimovNetwork/asimov/common"
)

type VoteEvent struct{}

func (voteEvent VoteEvent) handle(b blockInfo, args map[string]interface{}) error {
	proposalId, ok := args["proposalId"]
	if !ok {
		return voteEvent.checkArg("proposalId")
	}

	voter, ok := args["voter"]
	if !ok {
		return voteEvent.checkArg("voter")
	}

	decision, ok := args["decision"]
	if !ok {
		return voteEvent.checkArg("decision")
	}

	return foundationVoteService.Insert(b.height, b.time, (*proposalId.(**big.Int)).Int64(), voter.(*asimovCommon.Address).Hex(), *decision.(*bool), b.txHash)
}

func (voteEvent VoteEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("vote event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
