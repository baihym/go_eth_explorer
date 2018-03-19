package handle

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/onrik/ethrpc"

	"github.com/baihym/go_eth_explorer/app/db/mysql"
	"github.com/baihym/go_eth_explorer/app/enums"
	"github.com/baihym/go_eth_explorer/app/rpc"
	"github.com/baihym/go_eth_explorer/app/util"
)

func SearchAndSaveTransaction(blockNumber int) {

	block := rpc.EthGetBlockByNumber(blockNumber)
	if len(block.Transactions) == 0 {
		return
	}

	for _, tx := range block.Transactions {

		txReceipt := rpc.EthGetTransactionReceipt(tx.Hash)

		if strings.Contains(tx.Input, "18160ddd") &&
			strings.Contains(tx.Input, "70a08231") &&
			strings.Contains(tx.Input, "dd62ed3e") &&
			strings.Contains(tx.Input, "a9059cbb") &&
			strings.Contains(tx.Input, "095ea7b3") &&
			strings.Contains(tx.Input, "23b872dd") {
			name := rpc.EthCall(ethrpc.T{From: tx.From, To: txReceipt.ContractAddress, Data: "0x06fdde03"})
			symbol := rpc.EthCall(ethrpc.T{From: tx.From, To: txReceipt.ContractAddress, Data: "0x95d89b41"})
			if name != "0x" && name != "0x0000000000000000000000000000000000000000000000000000000000000001" &&
				name != "0x00000000000000000000000000000000000000000000000000000000000000"+
					"200000000000000000000000000000000000000000000000000000000000000000" &&
				symbol != "0x" && symbol != "0x0000000000000000000000000000000000000000000000000000000000000001" &&
				symbol != "0x00000000000000000000000000000000000000000000000000000000000000"+
					"200000000000000000000000000000000000000000000000000000000000000000" {
				continue
			}
		}

		// token contract
		if mysql.DBTokens[tx.To] > 0 {

			// transfer
			if strings.Index(tx.Input, "0xa9059cbb") == 0 {

				if len(tx.Input) == 138 {
					_, address, amountHex := tx.Input[0:10], tx.Input[10:10+64], tx.Input[10+64:]
					address = util.GetActualHex(address)
					if address == "" || tx.From == "" {
						panic("address or tx.From is invalid: " + tx.Input)
					}
					amountHex = util.GetActualHex(amountHex)

					mysqlTx := mysql.Transaction{
						HashId:         mysql.GetHashIdByHashForceWithPanic(tx.Hash),
						HashIndex:      0,
						FromAddressId:  mysql.GetAddressIdByAddressForceWithPanic(tx.From),
						ToAddressId:    mysql.GetAddressIdByAddressForceWithPanic(address),
						BlockNumber:    int64(blockNumber),
						Amount:         util.HexToBig(amountHex).String(),
						TokenId:        mysql.DBTokens[tx.To],
						GasUsed:        fmt.Sprintf("%d", txReceipt.GasUsed),
						GasPrice:       (&tx.GasPrice).String(),
						Value:          (&tx.Value).String(),
						InoutType:      enums.TransactionTypeInOut,
						BlockTimestamp: int64(block.Timestamp),
					}

					mysql.InsertTransactionIfNotExistWithPanic(&mysqlTx)
					continue
				}
			}

			// get internal txs
			internalTxs, err := util.GetInternalTransactionByHash(tx.Hash)
			if err != nil {
				panic(fmt.Sprintf("Get internal txs err: %s", err.Error()))
			}

			// not transfer
			if len(internalTxs) == 0 {
				if tx.From == "" || tx.To == "" {
					continue
				}
				txReceipt := rpc.EthGetTransactionReceipt(tx.Hash)
				mysqlTx := mysql.Transaction{
					HashId:         mysql.GetHashIdByHashForceWithPanic(tx.Hash),
					HashIndex:      0,
					FromAddressId:  mysql.GetAddressIdByAddressForceWithPanic(tx.From),
					ToAddressId:    mysql.GetAddressIdByAddressForceWithPanic(tx.To),
					BlockNumber:    int64(blockNumber),
					Amount:         "0",
					TokenId:        mysql.DBTokens[tx.To],
					GasUsed:        fmt.Sprintf("%d", txReceipt.GasUsed),
					GasPrice:       (&tx.GasPrice).String(),
					Value:          (&tx.Value).String(),
					InoutType:      enums.TransactionTypeOther,
					BlockTimestamp: int64(block.Timestamp),
				}

				mysql.InsertTransactionIfNotExistWithPanic(&mysqlTx)
				continue
			}

			// handle internal txs
			var i int64 = 0
			for _, internalTx := range internalTxs {
				if internalTx.IsError != "0" || internalTx.ErrCode != "" {
					panic("get internal txs err, hash: " + tx.Hash)
				}
				if internalTx.From == "" || internalTx.To == "" {
					continue
				}

				i++

				if mysql.DBTokens[internalTx.To] > 0 {

					if strings.Index(internalTx.Input, "0xa9059cbb") == 0 && len(internalTx.Input) == 138 {
						_, address, amountHex := tx.Input[0:10], tx.Input[10:10+64], tx.Input[10+64:]
						address = util.GetActualHex(address)
						if address == "" || tx.From == "" {
							panic("address or tx.From is invalid: " + tx.Input)
						}
						amountHex = util.GetActualHex(amountHex)

						mysqlTx := mysql.Transaction{
							HashId:         mysql.GetHashIdByHashForceWithPanic(tx.Hash),
							HashIndex:      i,
							FromAddressId:  mysql.GetAddressIdByAddressForceWithPanic(internalTx.From),
							ToAddressId:    mysql.GetAddressIdByAddressForceWithPanic(address),
							BlockNumber:    internalTx.BlockNumber,
							Amount:         util.HexToBig(amountHex).String(),
							TokenId:        mysql.DBTokens[internalTx.To],
							GasUsed:        fmt.Sprintf("%d", internalTx.GasUsed),
							GasPrice:       fmt.Sprintf("%d", internalTx.Gas),
							Value:          fmt.Sprintf("%s", internalTx.Value),
							InoutType:      enums.TransactionTypeInOut,
							BlockTimestamp: internalTx.TimeStamp,
						}

						mysql.InsertTransactionIfNotExistWithPanic(&mysqlTx)
						continue
					}

					mysqlTx := mysql.Transaction{
						HashId:         mysql.GetHashIdByHashForceWithPanic(tx.Hash),
						HashIndex:      i,
						FromAddressId:  mysql.GetAddressIdByAddressForceWithPanic(internalTx.From),
						ToAddressId:    mysql.GetAddressIdByAddressForceWithPanic(internalTx.To),
						BlockNumber:    internalTx.BlockNumber,
						Amount:         fmt.Sprintf("%s", internalTx.Value),
						TokenId:        mysql.DBTokens[internalTx.To],
						GasUsed:        fmt.Sprintf("%d", internalTx.GasUsed),
						GasPrice:       fmt.Sprintf("%d", internalTx.Gas),
						Value:          fmt.Sprintf("%s", internalTx.Value),
						InoutType:      enums.TransactionTypeOther,
						BlockTimestamp: internalTx.TimeStamp,
					}

					mysql.InsertTransactionIfNotExistWithPanic(&mysqlTx)
					continue
				}

				mysqlTx := mysql.Transaction{
					HashId:         mysql.GetHashIdByHashForceWithPanic(tx.Hash),
					HashIndex:      i,
					FromAddressId:  mysql.GetAddressIdByAddressForceWithPanic(internalTx.From),
					ToAddressId:    mysql.GetAddressIdByAddressForceWithPanic(internalTx.To),
					BlockNumber:    internalTx.BlockNumber,
					Amount:         fmt.Sprintf("%s", internalTx.Value),
					TokenId:        1,
					GasUsed:        fmt.Sprintf("%d", internalTx.GasUsed),
					GasPrice:       fmt.Sprintf("%d", internalTx.Gas),
					Value:          fmt.Sprintf("%s", internalTx.Value),
					InoutType:      enums.TransactionTypeInOut,
					BlockTimestamp: internalTx.TimeStamp,
				}

				mysql.InsertTransactionIfNotExistWithPanic(&mysqlTx)
				continue
			}
		}

		// eth
		if tx.From == "" || tx.To == "" {
			continue
		}
		mysqlTx := mysql.Transaction{
			HashId:         mysql.GetHashIdByHashForceWithPanic(tx.Hash),
			HashIndex:      0,
			FromAddressId:  mysql.GetAddressIdByAddressForceWithPanic(tx.From),
			ToAddressId:    mysql.GetAddressIdByAddressForceWithPanic(tx.To),
			BlockNumber:    int64(blockNumber),
			Amount:         (&tx.Value).String(),
			TokenId:        1,
			GasUsed:        fmt.Sprintf("%d", txReceipt.GasUsed),
			GasPrice:       (&tx.GasPrice).String(),
			Value:          (&tx.Value).String(),
			InoutType:      enums.TransactionTypeInOut,
			BlockTimestamp: int64(block.Timestamp),
		}

		mysql.InsertTransactionIfNotExistWithPanic(&mysqlTx)
		continue
	}
}
