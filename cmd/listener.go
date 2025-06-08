package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"ping/config"
	"ping/contracts"
	"time"

	"slices"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
)

type PingPongLog struct {
	CreatedToMineDuration    time.Duration
	MineToBackendDuration    time.Duration
	CreatedToBackendDuration time.Duration
}

var (
	ConsumedPong = make([]PingPongLog, 0, 12000)
	pingABI, _   = contracts.PingMetaData.GetAbi()
)

var ListenerCmd = &cobra.Command{
	Use:  "listen",
	RunE: RunE,
}

func RunE(cmd *cobra.Command, _ []string) error {
	defer report()
	pingParser := config.GlobalConfig.PingContract
	eventChan := make(chan types.Log, 10000)
	go consumeEvent(eventChan)

	for log := range eventChan {
		pong, _ := pingParser.ParsePong(log)
		createdTime := time.UnixMilli(int64(pong.CreatedTimestamp.Uint64()))
		mineTime := time.Unix(int64(pong.BlockTimestamp.Uint64()), 0)
		now := time.Now()
		ConsumedPong = append(ConsumedPong, PingPongLog{
			CreatedToMineDuration:    mineTime.Sub(createdTime),
			MineToBackendDuration:    now.Sub(mineTime),
			CreatedToBackendDuration: now.Sub(createdTime),
		})
		if len(ConsumedPong) == 12000 {
			break
		}
	}
	return nil
}

func consumeEvent(eventChan chan types.Log) {
	ethClient := config.GlobalConfig.Client

	header, err := ethClient.HeaderByNumber(context.Background(), nil) // Get current block header
	if err != nil {
		log.Fatalf("Failed to get current block header: %v", err)
	}
	lastProcessedBlock := header.Number.Uint64()
	log.Printf("Starting to listen for Pong() events from block %d\n", lastProcessedBlock+1)

	totalConsumed := 0

	for {
		header, err = ethClient.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Fatalf("Failed to get current block: %v", err)
		}
		currentBlock := header.Number.Uint64()
		if currentBlock <= lastProcessedBlock {
			log.Printf("Current block %d is less than or equal to last processed block %d. Sleeping for 1 seconds.", currentBlock, lastProcessedBlock)
			time.Sleep(time.Second)
			continue
		}

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(lastProcessedBlock + 1)),
			ToBlock:   header.Number,
			Addresses: []common.Address{common.HexToAddress(config.GlobalConfig.PingAddress)},
			Topics:    [][]common.Hash{{pingABI.Events["Pong"].ID}},
		}

		logs, err := ethClient.FilterLogs(context.Background(), query)
		if err != nil {
			log.Printf("Failed to filter logs: %v. Retrying...", err)
			time.Sleep(1 * time.Second) // Wait a bit before retrying on RPC error
			continue
		}

		totalConsumed += len(logs)
		log.Printf("Got %d logs, total consumed: %d", len(logs), totalConsumed)

		for _, vLog := range logs {
			// Ensure we don't process logs from blocks we've already covered
			if vLog.BlockNumber <= lastProcessedBlock {
				continue
			}

			eventChan <- vLog
		}

		lastProcessedBlock = currentBlock
	}
}

func report() {
	if len(ConsumedPong) == 0 {
		log.Println("No Pong events consumed, nothing to report.")
		return
	}

	createdToMineDurations := make([]time.Duration, len(ConsumedPong))
	mineToBackendDurations := make([]time.Duration, len(ConsumedPong))
	createdToBackendDurations := make([]time.Duration, len(ConsumedPong))

	for i, log := range ConsumedPong {
		createdToMineDurations[i] = log.CreatedToMineDuration
		mineToBackendDurations[i] = log.MineToBackendDuration
		createdToBackendDurations[i] = log.CreatedToBackendDuration
	}

	printStats("CreatedToMineDuration", createdToMineDurations)
	printStats("MineToBackendDuration", mineToBackendDurations)
	printStats("CreatedToBackendDuration", createdToBackendDurations)
}

func printStats(name string, durations []time.Duration) {
	if len(durations) == 0 {
		return
	}

	d := make([]time.Duration, len(durations))
	copy(d, durations)
	slices.Sort(d)

	min := d[0]
	max := d[len(d)-1]

	p90Index := int(float64(len(d)) * 0.90)
	p95Index := int(float64(len(d)) * 0.95)

	if p90Index >= len(d) {
		p90Index = len(d) - 1
	}
	if p95Index >= len(d) {
		p95Index = len(d) - 1
	}

	p90 := d[p90Index]
	p95 := d[p95Index]

	var total time.Duration
	for _, dur := range d {
		total += dur
	}
	avg := total / time.Duration(len(d))

	fmt.Printf("--- %s Stats ---\n", name)
	fmt.Printf("Total count: %d\n", len(d))
	fmt.Printf("Max: %v\n", max)
	fmt.Printf("Min: %v\n", min)
	fmt.Printf("P90: %v\n", p90)
	fmt.Printf("P95: %v\n", p95)
	fmt.Printf("Average: %v\n", avg)
	fmt.Println()
}
