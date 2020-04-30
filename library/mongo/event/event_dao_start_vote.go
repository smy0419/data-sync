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

type StartVoteEvent struct{}

var daoProposalService = service.DaoProposalService{}

func (startVoteEvent StartVoteEvent) handle(b blockInfo, args map[string]interface{}) error {
	organizationAddress, ok := args["organizationAddress"]
	if !ok {
		return startVoteEvent.checkArg("organizationAddress")
	}

	endTime, ok := args["endTime"]
	if !ok {
		return startVoteEvent.checkArg("endTime")
	}

	voteId, ok := args["voteId"]
	if !ok {
		return startVoteEvent.checkArg("voteId")
	}

	err := daoProposalService.Insert(b.height, b.time, b.txHash, organizationAddress.(*asimovCommon.Address).Hex(), (*endTime.(**big.Int)).Int64(), (*voteId.(**big.Int)).Int64(), model.ProposalTypeIssueAsset)
	if err != nil {
		return err
	}

	err = daoTodoListService.InsertMany(b.height, b.time, organizationAddress.(*asimovCommon.Address).Hex(), model.TodoTypeVote, (*endTime.(**big.Int)).Int64(), (*voteId.(**big.Int)).Int64())
	if err != nil {
		return err
	}
	return nil
}

func (startVoteEvent StartVoteEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("dao start vote event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
