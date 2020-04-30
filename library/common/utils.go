package common

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/asimov/crypto"
)

func ToStruct(source interface{}, destination interface{}) error {
	jsonStr, err := json.Marshal(source)
	if err != nil {
		Logger.Errorf("marshal  failed. source: %v, err: %s", source, err)
		return err
	}

	err = json.Unmarshal(jsonStr, destination)
	if err != nil {
		Logger.Errorf("unmarshal failed. source: %s, err: %s", string(jsonStr), err)
		return err
	}

	return nil
}

func ValidateSignature(signedMessage string, message string, pubKey *crypto.PublicKey) bool {
	sigBytes, err := base64.StdEncoding.DecodeString(signedMessage)
	messageHash := common.DoubleHashB([]byte(message))

	fmt.Println([]byte(message))
	signature, err := crypto.ParseSignature(sigBytes, crypto.S256())
	if err != nil {
		Logger.Error("decode sign to hash error")
		return false
	}
	return signature.Verify(messageHash, pubKey)
}

/**
Filter duplicate elements through a double loop
It is suitable for cases with a small amount of data
*/
func RemoveRepeatByLoop(slc []string) []string {
	result := make([]string, 0) // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, slc[i])
		}
	}
	return result
}

/**
The unique feature of the map key value is deduplicated
Suitable for large data volume
*/
func RemoveDuplicateByMap(arr []string) []string {
	resArr := make([]string, 0)
	tmpMap := make(map[string]interface{})
	for _, val := range arr {
		if _, ok := tmpMap[val]; !ok {
			resArr = append(resArr, val)
			tmpMap[val] = nil
		}
	}
	return resArr
}

func GetEffectiveTime(height int64, effectiveHeight int64) int64 {
	return NowSecond() + effectiveHeight*MiningFrequency
}
