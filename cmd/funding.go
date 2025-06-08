package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"ping/config"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var GenerateWalletsCmd = &cobra.Command{
	Use:  "generate-wallets",
	RunE: GenerateWallets,
}

var FundWalletsCmd = &cobra.Command{
	Use:  "fund-wallets",
	RunE: FundWallets,
}

type Wallet struct {
	PrivateKey string `json:"private_key"`
	Address    string `json:"address"`
}

func GenerateWallets(cmd *cobra.Command, _ []string) error {
	var wallets []Wallet

	f, err := os.OpenFile("wallets.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for range 200 {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}

		privateKeyBytes := crypto.FromECDSA(privateKey)
		privateKeyHex := hexutil.Encode(privateKeyBytes)[2:]

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public key to ECDSA")
		}

		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

		wallets = append(wallets, Wallet{
			PrivateKey: privateKeyHex,
			Address:    address,
		})

	}

	json.NewEncoder(f).Encode(wallets)

	fmt.Println("200 wallets saved to wallets.json")
	return nil
}

func FundWallets(cmd *cobra.Command, _ []string) error {
	walletAddresses, _ := loadWallets()
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(config.GlobalConfig.RootPrivateKey, "0x"))
	if err != nil {
		log.Fatal(err)
	}

	pingContract := config.GlobalConfig.PingContract

	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, config.GlobalConfig.ChainID)
	if err != nil {
		log.Fatal(err)
	}

	ethSendAmount := big.NewInt(config.GlobalConfig.EthSendAmount)
	value := new(big.Int).Mul(ethSendAmount, big.NewInt(int64(len(walletAddresses))))
	txOpts.Value = value

	tx, err := pingContract.MultiSend(txOpts, walletAddresses, ethSendAmount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction sent:", tx.Hash().Hex())

	return nil
}
