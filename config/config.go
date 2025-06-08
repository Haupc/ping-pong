package config

import (
	"context"
	"log"
	"math/big"
	"ping/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

var GlobalConfig Config

type Config struct {
	RpcURL         string `mapstructure:"rpc_url"`
	PingAddress    string `mapstructure:"ping_address"`
	RootPrivateKey string `mapstructure:"root_private_key"`
	EthSendAmount  int64  `mapstructure:"eth_send_amount"`

	Client       *ethclient.Client
	PingContract *contracts.Ping
	ChainID      *big.Int
}

func LoadConfig() {
	viper.SetConfigFile("config.yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("failed to load config")
	}

	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		log.Fatalln("failed to load config")
	}

	GlobalConfig.Client, err = ethclient.Dial(GlobalConfig.RpcURL)
	if err != nil {
		log.Fatalln("failed to connect to rpc")
	}

	GlobalConfig.ChainID, err = GlobalConfig.Client.ChainID(context.Background())
	if err != nil {
		log.Fatalln("failed to get chain id")
	}

	GlobalConfig.PingContract, err = contracts.NewPing(common.HexToAddress(GlobalConfig.PingAddress), GlobalConfig.Client)
	if err != nil {
		log.Fatalln("failed to connect to ping contract")
	}
}
