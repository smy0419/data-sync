package event

import (
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
	"math/big"
)

type UpdateRoundBlockInfoEvent struct{}

func (updateRoundBlockInfoEvent UpdateRoundBlockInfoEvent) handle(b blockInfo, args map[string]interface{}) error {
	round, ok := args["round"]
	if !ok {
		return updateRoundBlockInfoEvent.checkArg("round")
	}

	validatorAddresses, ok := args["validators"]
	if !ok {
		return updateRoundBlockInfoEvent.checkArg("validators")
	}
	validatorsSlice := make([]string, 0)
	for _, v := range *validatorAddresses.(*[]asimovCommon.Address) {
		validatorsSlice = append(validatorsSlice, v.Hex())
	}

	plannedBlocks, ok := args["plannedBlocks"]
	if !ok {
		return updateRoundBlockInfoEvent.checkArg("plannedBlocks")
	}
	plannedBlocksSlice := make([]int64, 0)
	for _, v := range *plannedBlocks.(*[]*big.Int) {
		plannedBlocksSlice = append(plannedBlocksSlice, v.Int64())
	}

	actualBlocks, ok := args["actualBlocks"]
	if !ok {
		return updateRoundBlockInfoEvent.checkArg("actualBlocks")
	}
	actualBlocksSlice := make([]int64, 0)
	for _, v := range *actualBlocks.(*[]*big.Int) {
		actualBlocksSlice = append(actualBlocksSlice, v.Int64())
	}

	err := validatorService.ModifyBlocks(b.height, validatorsSlice, plannedBlocksSlice, actualBlocksSlice)
	if err != nil {
		return err
	}

	err = minerSignUpService.UpdateBlocksAndEfficiency(b.height, (*round.(**big.Int)).Int64(), validatorsSlice, plannedBlocksSlice, actualBlocksSlice)
	if err != nil {
		return err
	}

	err = minerMemberService.UpdateBlocksAndEfficiency(b.height, (*round.(**big.Int)).Int64(), validatorsSlice, plannedBlocksSlice, actualBlocksSlice)
	if err != nil {
		return err
	}

	return nil
}

func (updateRoundBlockInfoEvent UpdateRoundBlockInfoEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("update round block info event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
