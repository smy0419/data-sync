package service

import (
	"encoding/json"
	"github.com/AsimovNetwork/data-sync/library/mongo/model"
	"github.com/AsimovNetwork/data-sync/library/redis"
	"strings"
)

type KeyTransactionService struct{}

type TransactionCacheService struct{}

var redisClient = redis.RedisClient{}

const (
	KeyListMaxIndex     int    = 99
	TransactionMaxIndex int    = 999
	TransactionKey      string = "transactions"
	TransactionCountKey string = "transaction_count"
	TenDaySeconds       int    = 10 * 24 * 60 * 60
)

func (keyTransactionService KeyTransactionService) Record(keyMap map[string]model.KeyTransaction) error {
	if len(keyMap) <= 0 {
		return nil
	}

	splitMap := make(map[string][]interface{})
	for k, v := range keyMap {
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			return err
		}

		key := strings.Split(k, "_")[0]
		if _, ok := splitMap[key]; !ok {
			splitMap[key] = []interface{}{string(jsonBytes)}
		} else {
			splitMap[key] = append(splitMap[key], string(jsonBytes))
		}
	}

	command := make([]redis.RedisCommand, 0)
	for k, v := range splitMap {
		lPushArgs := append([]interface{}{k}, v...)
		lPush := redis.RedisCommand{
			Name: "LPUSH",
			Args: lPushArgs,
		}
		ltrim := redis.RedisCommand{
			Name: "LTRIM",
			Args: []interface{}{k, 0, KeyListMaxIndex},
		}
		command = append(command, lPush, ltrim)
	}

	return redisClient.Pipeline(command...)
}

func (transactionCacheService TransactionCacheService) CacheTransaction(transactions []interface{}) error {
	keyTransactions := make([]interface{}, 0)
	for _, v := range transactions {
		tx := v.(*model.Transaction)
		tmp := model.KeyTransaction{
			TxHash: tx.Hash,
			Time:   tx.Time,
			Fee:    tx.Fee,
		}
		jsonBytes, err := json.Marshal(tmp)
		if err != nil {
			return err
		}
		keyTransactions = append(keyTransactions, string(jsonBytes))
	}

	command := make([]redis.RedisCommand, 0)
	lPush := redis.RedisCommand{
		Name: "LPUSH",
		Args: append([]interface{}{TransactionKey}, keyTransactions...),
	}
	ltrim := redis.RedisCommand{
		Name: "LTRIM",
		Args: []interface{}{TransactionKey, 0, TransactionMaxIndex},
	}
	incTxCount := redis.RedisCommand{
		Name: "INCRBY",
		Args: []interface{}{TransactionCountKey, len(keyTransactions)},
	}
	command = append(command, lPush, ltrim, incTxCount)

	for _, v := range transactions {
		cmd := redis.RedisCommand{
			Name: "SETEX",
			Args: []interface{}{v.(*model.Transaction).Hash, TenDaySeconds, v.(*model.Transaction).Height},
		}
		command = append(command, cmd)
	}

	return redisClient.Pipeline(command...)
}

func (transactionCacheService TransactionCacheService) RollbackTxCount(count interface{}) error {
	return redisClient.Increase(TransactionCountKey, count)
}

func (transactionCacheService TransactionCacheService) RollbackTransactionHeight(txHash []interface{}) error {
	return redisClient.Delete(txHash...)
}
