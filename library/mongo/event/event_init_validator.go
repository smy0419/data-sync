package event

import (
	"github.com/AsimovNetwork/asimov/chaincfg"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/data-sync/library/common"
)

type InitValidatorEvent struct{}

// Execute when the block height is 0
func (initValidatorEvent InitValidatorEvent) handle(b blockInfo, args map[string]interface{}) error {
	if b.height != 0 {
		return nil
	}

	validators := make([]asimovCommon.Address, 0)
	if common.Cfg.Env == common.ENV_DEVELOP_NET {
		validators = chaincfg.DevelopNetParams.GenesisCandidates
	} else if common.Cfg.Env == common.ENV_TEST_NET {
		validators = chaincfg.TestNetParams.GenesisCandidates
	} else if common.Cfg.Env == common.ENV_MAIN_NET {
		validators = chaincfg.MainNetParams.GenesisCandidates
	}
	for _, validator := range validators {
		err := validatorService.Insert(b.height, b.time, validator.Hex())
		if err != nil {
			return err
		}
	}

	return nil
}
