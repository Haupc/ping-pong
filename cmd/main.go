package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"log"
	"os"
	"ping/config"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func main() {
	cobra.OnInitialize(config.LoadConfig)
	rootCmd.AddCommand(ListenerCmd)
	rootCmd.AddCommand(GenerateWalletsCmd)
	rootCmd.AddCommand(FundWalletsCmd)
	rootCmd.AddCommand(PingCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func loadWallets() ([]common.Address, []*ecdsa.PrivateKey) {
	f, err := os.Open("wallets.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	wallets := []Wallet{}
	json.NewDecoder(f).Decode(&wallets)

	var walletAddresses []common.Address
	var privateKeys []*ecdsa.PrivateKey

	for _, wallet := range wallets {
		walletAddresses = append(walletAddresses, common.HexToAddress(wallet.Address))
		privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(wallet.PrivateKey, "0x"))
		if err != nil {
			log.Fatal(err)
		}
		privateKeys = append(privateKeys, privateKey)
	}

	return walletAddresses, privateKeys
}
