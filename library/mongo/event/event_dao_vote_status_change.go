package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"math/big"
)

type VoteStatusChangeEvent struct{}

func (voteStatusChangeEvent VoteStatusChangeEvent) handle(b blockInfo, args map[string]interface{}) error {
	organizationAddress, ok := args["organizationAddress"]
	if !ok {
		return voteStatusChangeEvent.checkArg("organizationAddress")
	}

	voteId, ok := args["voteId"]
	if !ok {
		return voteStatusChangeEvent.checkArg("voteId")
	}

	status, ok := args["status"]
	if !ok {
		return voteStatusChangeEvent.checkArg("status")
	}
	err := daoProposalService.Update(b.height, organizationAddress.(*asimovCommon.Address).Hex(), (*voteId.(**big.Int)).Int64(), *status.(*uint8))
	if err != nil {
		return err
	}

	return nil
}

func (voteStatusChangeEvent VoteStatusChangeEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("dao start vote event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
