package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/baihym/go_eth_explorer/app/config"
)

type InternalTransactionItem struct {
	BlockNumber     int64  `json:"blockNumber"`
	TimeStamp       int64  `json:"timeStamp"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	ContractAddress string `json:"contractAddress"`
	Input           string `json:"input"`
	Gas             int64  `json:"gas"`
	GasUsed         int64  `json:"gasUsed"`
	IsError         string `json:"isError"`
	ErrCode         string `json:"errCode"`
}

type InternalTransactionItemJson struct {
	BlockNumber     string `json:"blockNumber"`
	TimeStamp       string `json:"timeStamp"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	ContractAddress string `json:"contractAddress"`
	Input           string `json:"input"`
	Gas             string `json:"gas"`
	GasUsed         string `json:"gasUsed"`
	IsError         string `json:"isError"`
	ErrCode         string `json:"errCode"`
}

func GetInternalTransactionByHash(hash string) ([]InternalTransactionItem, error) {
	httpResp, err := HttpClient.Get(config.ETHERSCANHost + "/api?module=account&action=txlistinternal&txhash=" + hash + "&apikey=" + config.ETHERSCANApiKeyToken)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request eth scan api err code : %d", httpResp.StatusCode))
	}

	var rst struct {
		Status  string                        `json:"status"`
		Message string                        `json:"message"`
		Result  []InternalTransactionItemJson `json:"result"`
	}
	if err = json.NewDecoder(httpResp.Body).Decode(&rst); err != nil {
		return nil, err
	}

	if rst.Status != "1" || rst.Message != "OK" {
		fmt.Println("get internal txs err: status=" + rst.Status + "; msg=" + rst.Message)
		return nil, nil
	}

	var list []InternalTransactionItem

	for _, item := range rst.Result {
		i := InternalTransactionItem{
			From:            item.From,
			To:              item.To,
			Value:           item.Value,
			ContractAddress: item.ContractAddress,
			Input:           item.Input,
			IsError:         item.IsError,
			ErrCode:         item.ErrCode,
		}
		i.BlockNumber, _ = strconv.ParseInt(item.BlockNumber, 10, 64)
		i.TimeStamp, _ = strconv.ParseInt(item.TimeStamp, 10, 64)
		i.Gas, _ = strconv.ParseInt(item.Gas, 10, 64)
		i.GasUsed, _ = strconv.ParseInt(item.GasUsed, 10, 64)
		list = append(list, i)
	}

	return list, nil
}
