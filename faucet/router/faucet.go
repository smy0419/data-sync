package router

import (
	"github.com/AsimovNetwork/data-sync/faucet/service"
	"github.com/AsimovNetwork/data-sync/library/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type faucetRequest struct {
	Address string `json:"receiver" binding:"required"`
	Amount  int64  `json:"amount" binding:"required"` // 单位：xing
}

var faucetService = service.FaucetService{}

/**
arg: faucetRequest
result: tx_hash
*/
func FaucetTransfer(c *gin.Context) {
	var request faucetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, response.Error(response.BadArguments))
		return
	}

	// 1、通过私钥生成地址
	address, curvePrivate, err := faucetService.GenerateAddressViaPrivate()
	if err != nil {
		c.JSON(http.StatusOK, response.Error(response.SystemError))
		return
	}

	// 2、通过地址拿UTXO
	utxoSlice, err := faucetService.GetUTXO(address.String())
	if err != nil {
		c.JSON(http.StatusOK, response.Error(response.SystemError))
		return
	}

	// 3、构建交易
	redeemTx, err := faucetService.GenerateTransaction(request.Address, request.Amount, utxoSlice, address)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(response.SystemError))
		return
	}
	if redeemTx.TxOut == nil || len(redeemTx.TxOut) == 0 {
		c.JSON(http.StatusOK, response.Error(response.NotEnoughAmount))
		return
	}

	// 4、签名交易
	err = faucetService.SignTransaction(address, redeemTx, curvePrivate)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(response.SystemError))
		return
	}

	// 5、发送交易
	result, err := faucetService.SendTransaction(redeemTx)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(response.SystemError))
		return
	}

	c.JSON(http.StatusOK, response.OK(gin.H{
		"tx_hash": result,
	}))
}
