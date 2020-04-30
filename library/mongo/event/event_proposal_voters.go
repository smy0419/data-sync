package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"math/big"
)

type ProposalVotersEvent struct{}

func (proposalVotersEvent ProposalVotersEvent) handle(b blockInfo, args map[string]interface{}) error {
	round, ok := args["round"]
	if !ok {
		return proposalVotersEvent.checkArg("round")
	}

	proposalId, ok := args["proposalId"]
	if !ok {
		return proposalVotersEvent.checkArg("proposalId")
	}

	voters, ok := args["voters"]
	if !ok {
		return proposalVotersEvent.checkArg("voters")
	}
	votersSlice := make([]string, 0)
	for _, v := range *voters.(*[]asimovCommon.Address) {
		votersSlice = append(votersSlice, v.Hex())
	}

	return minerTodoListService.Insert(
		b.height,
		b.time,
		(*round.(**big.Int)).Int64(),
		votersSlice,
		(*proposalId.(**big.Int)).Int64(),
	)
}

func (proposalVotersEvent ProposalVotersEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("proposal voters event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
