package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"math/big"
)

type VoteForCommitteeEvent struct{}

func (voteForCommitteeEvent VoteForCommitteeEvent) handle(b blockInfo, args map[string]interface{}) error {
	round, ok := args["round"]
	if !ok {
		return voteForCommitteeEvent.checkArg("round")
	}

	proposalId, ok := args["proposalId"]
	if !ok {
		return voteForCommitteeEvent.checkArg("proposalId")
	}

	voter, ok := args["voter"]
	if !ok {
		return voteForCommitteeEvent.checkArg("voter")
	}

	decision, ok := args["decision"]
	if !ok {
		return voteForCommitteeEvent.checkArg("decision")
	}

	return minerVoteService.Insert(
		(*round.(**big.Int)).Int64(),
		b.height,
		b.time,
		(*proposalId.(**big.Int)).Int64(),
		voter.(*asimovCommon.Address).Hex(),
		*decision.(*bool),
		b.txHash,
	)
}

func (voteForCommitteeEvent VoteForCommitteeEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("miner vote for committee event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
