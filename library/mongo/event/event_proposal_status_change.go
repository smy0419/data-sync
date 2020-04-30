package event

import (
	"errors"
	"fmt"
	"github.com/AsimovNetwork/data-sync/library/common"
	"math/big"
)

type ProposalStatusChangeEvent struct{}

func (proposalStatusChangeEvent ProposalStatusChangeEvent) handle(b blockInfo, args map[string]interface{}) error {
	proposalId, ok := args["proposalId"]
	if !ok {
		return proposalStatusChangeEvent.checkArg("proposalId")
	}

	status, ok := args["status"]
	if !ok {
		return proposalStatusChangeEvent.checkArg("status")
	}

	return foundationProposalService.ChangeStatus(b.height, b.time, (*proposalId.(**big.Int)).Int64(), *status.(*uint8))
}

func (proposalStatusChangeEvent ProposalStatusChangeEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("proposal status change event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
