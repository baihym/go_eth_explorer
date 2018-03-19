package rpc

import (
	"github.com/onrik/ethrpc"
)

var client *ethrpc.EthRPC

func InitEthRPCClient(url string) {
	client = ethrpc.NewEthRPC(url)
}

func EthBlockNumber() int64 {
	blockNumber, err := client.EthBlockNumber()
	if err != nil {
		panic(err)
	}
	return int64(blockNumber)
}

func EthGetBlockByNumber(blockNumber int) *ethrpc.Block {
	block, err := client.EthGetBlockByNumber(blockNumber, true)
	if err != nil {
		panic(err)
	}
	return block
}

func EthGetTransactionReceipt(hash string) *ethrpc.TransactionReceipt {
	transactionReceipt, err := client.EthGetTransactionReceipt(hash)
	if err != nil {
		panic(err)
	}
	return transactionReceipt
}

func EthCall(params ethrpc.T) string {
	result, err := client.EthCall(params, "latest")
	if err != nil {
		panic(err)
	}
	return result
}
