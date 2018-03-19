package handle

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadHandleBlock(filePath string) int64 {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := ioutil.WriteFile(filePath, []byte("0"), 0644); err != nil {
			panic("Save last handle block number file err: " + err.Error())
		}
	}
	blockNumberBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic("Read last handle block number err: " + err.Error())
	}

	blockNumber, err := strconv.ParseInt(strings.Replace(string(blockNumberBytes), "\n", "", -1), 10, 64)
	if err != nil {
		panic("Convert last handle block number err: " + err.Error())
	}

	return blockNumber
}

func WriteHandleBlock(filePath string, blockNumber int64) {
	err := ioutil.WriteFile(filePath, []byte(fmt.Sprintf("%d", blockNumber)), 0644)
	if err != nil {
		panic("Write last handle block number err: " + err.Error())
	}
}
