package event

import (
	"encoding/hex"
	"errors"
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/asimov/protos"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"math/big"
)

type DonateEvent struct{}

func (donateEvent DonateEvent) handle(b blockInfo, args map[string]interface{}) error {
	donator, ok := args["donator"]
	if !ok {
		return donateEvent.checkArg("donator")
	}

	assetType, ok := args["assetType"]
	if !ok {
		return donateEvent.checkArg("assetType")
	}
	assetBytes := protos.AssetFromInt(*assetType.(**big.Int)).Bytes()

	amount, ok := args["amount"]
	if !ok {
		return donateEvent.checkArg("amount")
	}

	return foundationBalanceSheetService.Insert(
		b.height,
		b.time,
		b.txHash,
		donator.(*asimovCommon.Address).Hex(),
		(*amount.(**big.Int)).Int64(),
		hex.EncodeToString(assetBytes),
		model.TransferTypeIn,
		-1,
	)
}

func (donateEvent DonateEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("donate event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
