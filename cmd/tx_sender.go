package main

import (
	"crypto/ecdsa"
	"log"
	"math/big"
	"math/rand"
	"ping/config"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"
)

var PingCmd = &cobra.Command{
	Use:  "ping",
	RunE: RunPing,
}

func RunPing(cmd *cobra.Command, _ []string) error {
	_, privateKeys := loadWallets()

	var wg sync.WaitGroup
	for _, pk := range privateKeys {
		wg.Add(1)
		go func(pk *ecdsa.PrivateKey) {
			defer wg.Done()
			runPing(pk)
		}(pk)
	}

	wg.Wait()

	return nil
}

func runPing(pk *ecdsa.PrivateKey) {
	pingContract := config.GlobalConfig.PingContract
	txOpts, err := bind.NewKeyedTransactorWithChainID(pk, config.GlobalConfig.ChainID)
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()
	for range 60 {
		<-ticker.C
		sleepDuration := time.Duration(rand.Intn(13501)+500) * time.Millisecond
		time.Sleep(sleepDuration)

		now := time.Now()
		log.Printf("Sending ping at %s", now.Format(time.RFC3339))
		_, err := pingContract.Ping(txOpts, big.NewInt(now.UnixMicro()))
		if err != nil {
			log.Fatal(err)
		}
	}
}
