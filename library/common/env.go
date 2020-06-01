package common

import (
	"fmt"
	asimovCommon "github.com/AsimovNetwork/asimov/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const (
	ENV_DEVELOP_NET = "developnet"
	ENV_TEST_NET    = "testnet"
	ENV_MAIN_NET    = "mainnet"
)

type config struct {
	Env              string
	GinMode          string   `yaml:"gin_mode"`
	LogDir           string   `yaml:"log_dir"`
	Mysql            string   `yaml:"mysql"`
	ShowSql          bool     `yaml:"show_sql"`
	BlockChainRpc    string   `yaml:"block_chain_rpc"`
	FaucetPrivateKey string   `yaml:"faucet_private_key"`
	ProjectName      string   `yaml:"project_name"`
	Port             string   `yaml:"port"`
	MongodbUrl       string   `yaml:"mongodb_url"`
	MongodbDB        string   `yaml:"mongodb_db"`
	MongodbUser      string   `yaml:"mongodb_user"`
	MongodbPassword  string   `yaml:"mongodb_password"`
	DingTalkUrl      string   `yaml:"ding_talk_url"`
	DingTalkSecret   string   `yaml:"ding_talk_secret"`
	AtMobiles        []string `yaml:"at_mobiles"`
	Redis            string   `yaml:"redis"`
}

var Cfg config

func init() {
	argNum := len(os.Args)
	if argNum < 2 {
		fmt.Println("Missing environment parameters")
		os.Exit(1)
	}

	envTag := os.Args[1]
	var yamlFile []byte

	yamlFile, err := ioutil.ReadFile("config_" + envTag + ".yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Cfg)
	if err != nil {
		panic(err)
	}

	Cfg.Env = envTag
}

func (cfg config) GetAsimovNet() string {
	if cfg.Env == ENV_DEVELOP_NET {
		return asimovCommon.DevelopNet.String()
	} else if cfg.Env == ENV_TEST_NET {
		return asimovCommon.TestNet.String()
	} else if cfg.Env == ENV_MAIN_NET {
		return asimovCommon.MainNet.String()
	}
	return ""
}
