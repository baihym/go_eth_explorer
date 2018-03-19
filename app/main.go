package main

import (
	"fmt"
	"log"
	"strconv"
	"os"

	"github.com/baihym/go_eth_explorer/app/config"
	"github.com/baihym/go_eth_explorer/app/db/mysql"
	"github.com/baihym/go_eth_explorer/app/handle"
	"github.com/baihym/go_eth_explorer/app/rpc"
)

func main() {

	if len(os.Args) == 2 {
		if i, err := strconv.ParseInt(os.Args[1], 10, 64); err == nil && i >= 0 {
			handle.WriteHandleBlock(config.ETHLastBlockNumberFile, i)
		}
	}

	// Get tokens which want to search txs
	if err := mysql.InitDBTokens(); err != nil {
		panic("InitDBTokens err: " + err.Error())
	}

	// Init rpc
	rpc.InitEthRPCClient(config.ETHRPCAddress)

	handleBlock := handle.ReadHandleBlock(config.ETHLastBlockNumberFile)
	mostRecentBlock := rpc.EthBlockNumber()

	log.Println(fmt.Sprintf("handleBlock=%d, mostRecentBlock=%d", handleBlock, mostRecentBlock))

	for handleBlock <= mostRecentBlock {

		log.Println(handleBlock)

		handle.WriteHandleBlock(config.ETHLastBlockNumberFile, handleBlock)

		handle.SearchAndSaveTransaction(int(handleBlock))

		handleBlock++
	}
}
