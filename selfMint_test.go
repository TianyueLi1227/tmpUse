package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/RiemaLabs/modular-indexer-committee/ord/stateless"
)

func Test_SelfMint(t *testing.T) {
	// WARNING! This test is recommended to have .cache up to height at least 837000
	// Otherwise it will take too long to catch up to the latest height
	var catchupHeight uint = stateless.SelfMintEnableHeight
	records, err := stateless.LoadOPIRecords("./data/837000-ord_transfers.csv") // TODO: update to the self-mint data
	if err != nil {
		log.Fatalf(fmt.Sprintf("error happened: %v", err))
	}
	ordGetterTest, arguments := loadMain()
	queue, err := CatchupStage(ordGetterTest, &arguments, stateless.BRC20StartHeight-1, catchupHeight)
	if err != nil {
		log.Fatalf(fmt.Sprintf("error happened: %v", err))
	}
	if queue.Header.Height != catchupHeight {
		log.Println("Queue header not updated correctly")
	}
	ordGetterTest.LatestBlockHeight = catchupHeight
	go ServiceStage(ordGetterTest, &arguments, queue, 10*time.Millisecond)
	for {
		if ordGetterTest.LatestBlockHeight == queue.LatestHeight() {
			queue.Header.VerifyState(&records)
			log.Printf("Block: %d is verified!\n", ordGetterTest.LatestBlockHeight)
			ordGetterTest.LatestBlockHeight++
		}
		if ordGetterTest.LatestBlockHeight >= 837590 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
}
