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

type TransferAssetEvent struct{}

func (transferAssetEvent TransferAssetEvent) handle(b blockInfo, args map[string]interface{}) error {
	receivingAddress, ok := args["receivingAddress"]
	if !ok {
		return transferAssetEvent.checkArg("receivingAddress")
	}

	proposalType, ok := args["proposalType"]
	if !ok {
		return transferAssetEvent.checkArg("proposalType")
	}

	assetType, ok := args["assetType"]
	if !ok {
		return transferAssetEvent.checkArg("assetType")
	}
	assetBytes := protos.AssetFromInt(*assetType.(**big.Int)).Bytes()

	amount, ok := args["amount"]
	if !ok {
		return transferAssetEvent.checkArg("amount")
	}

	return foundationBalanceSheetService.Insert(
		b.height,
		b.time,
		b.txHash,
		receivingAddress.(*asimovCommon.Address).Hex(),
		(*amount.(**big.Int)).Int64(),
		hex.EncodeToString(assetBytes),
		model.TransferTypeOut,
		int8(*proposalType.(*uint8)),
	)
}

func (transferAssetEvent TransferAssetEvent) checkArg(arg string) error {
	errMsg := fmt.Sprintf("transfer asset event miss arg %s", arg)
	common.Logger.Error(errMsg)
	return errors.New(errMsg)
}
