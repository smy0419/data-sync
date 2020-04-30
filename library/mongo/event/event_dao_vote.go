package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"github.com/AsimovNetwork/data-sync/library/mongo/service"
	"math/big"
)

type DAOVoteEvent struct{}

var voteService = service.DaoVoteService{}

func (daoVoteEvent DAOVoteEvent) handle(b blockInfo, args map[string]interface{}) error {
	organizationAddress, ok := args["organizationAddress"]
	if !ok {
		return daoVoteEvent.checkArg("organizationAddress")
	}

	voter, ok := args["voter"]
	if !ok {
		return daoVoteEvent.checkArg("voter")
	}

	voteId, ok := args["voteId"]
	if !ok {
		return daoVoteEvent.checkArg("voteId")
	}

	decision, ok := args["decision"]
	if !ok {
		return daoVoteEvent.checkArg("decision")
	}

	err := voteService.Insert(b.height, b.time, b.txHash, organizationAddress.(*asimovCommon.Address).Hex(), voter.(*asimovCommon.Address).Hex(), (*voteId.(**big.Int)).Int64(), *decision.(*bool))

	err = daoTodoListService.ReleaseById(b.height, organizationAddress.(*asimovCommon.Address).Hex(), voter.(*asimovCommon.Address).Hex(), model.TodoTypeVote, (*voteId.(**big.Int)).Int64())
	if err != nil {
		return err
	}
	return nil
}

func (daoVoteEvent DAOVoteEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("dao vote event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
