package service

import (
	"fmt"
	"github.com/AsimovNetwork/asimov/chaincfg"
	"github.com/AsimovNetwork/data-sync/library/common"
	"github.com/AsimovNetwork/data-sync/library/mysql"
	models "github.com/AsimovNetwork/data-sync/library/mysql/model"
	"strings"
)

type ChainNodeService struct{}

func (chainNodeService ChainNodeService) Insert(height int64) {
	// 每500个块高执行一次
	if height%500 != 0 {
		return
	}

	// 获取Seed节点
	var seeds []chaincfg.DNSSeed
	if common.Cfg.Env == common.ENV_DEVELOP_NET {
		seeds = chaincfg.DevelopNetParams.DNSSeeds
	} else if common.Cfg.Env == common.ENV_TEST_NET {
		seeds = chaincfg.TestNetParams.DNSSeeds
	} else if common.Cfg.Env == common.ENV_MAIN_NET {
		seeds = chaincfg.MainNetParams.DNSSeeds
	} else {
		return
	}

	// 获取各个Seed的peers
	for _, seed := range seeds {
		getPeersFromSeed(seed.Host)
	}
}

func getPeersFromSeed(seed string) {
	url := fmt.Sprintf("http://%s", seed)
	param := common.NewChainRequest("currentPeers", []interface{}{})
	result, ok := common.Post(url, param)
	if !ok {
		common.Logger.Error("get current peers failed")
		return
	}
	peers := (result).([]interface{})
	for _, peer := range peers {
		err := insert(peer.(string))
		if err != nil {
			common.Logger.Error(err)
			continue
		}
	}
}

func insert(peer string) error {
	ip := strings.Split(peer, ":")[0]
	city, subdivision, country, longitude, latitude, err := common.GeoInfo(ip)
	if err != nil {
		// if err, print log
		common.Logger.Errorf("get geo info error, ip=%s, err=%v", ip, err)
	}

	ipExist, err := ipExist(ip)
	if err != nil {
		return err
	}

	now := common.NowSecond()
	if ipExist {
		sql := "update t_chain_node set update_time = ? where ip = ?"
		_, err := mysql.Engine.Exec(sql, now, ip)
		return err
	} else {
		chainNode := models.TChainNode{
			Id:          mysql.GlobalIdService.NextId(),
			Ip:          ip,
			City:        city,
			Subdivision: subdivision,
			Country:     country,
			Longitude:   longitude,
			Latitude:    latitude,
			CreateTime:  now,
			UpdateTime:  now,
		}

		_, err := mysql.Engine.InsertOne(chainNode)
		return err
	}
}

func ipExist(ip string) (bool, error) {
	p := new(models.TChainNode)
	total, err := mysql.Engine.Where("ip = ?", ip).Count(p)
	return total > 0, err
}
