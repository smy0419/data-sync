package event

import (
	"errors"
	"fmt"
	"github.com/AsimovNetwork/data-sync/library/common"
	"math/big"
)

type MultiAssetProposalEffectHeightEvent struct{}

func (multiAssetProposalEffectHeightEvent MultiAssetProposalEffectHeightEvent) handle(b blockInfo, args map[string]interface{}) error {
	proposalId, ok := args["proposalId"]
	if !ok {
		return multiAssetProposalEffectHeightEvent.checkArg("proposalId")
	}

	workHeight, ok := args["workHeight"]
	if !ok {
		return multiAssetProposalEffectHeightEvent.checkArg("workHeight")
	}

	return minerProposalService.ChangeStatusByWorkHeight(b.height, (*proposalId.(**big.Int)).Int64(), (*workHeight.(**big.Int)).Int64())
}

func (multiAssetProposalEffectHeightEvent MultiAssetProposalEffectHeightEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("multi asset proposal work height event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
