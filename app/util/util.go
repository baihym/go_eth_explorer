package util

import (
	"fmt"
	"math/big"
	"strings"
)

func GetActualHex(h string) string {
	h = strings.TrimLeft(h, "0")

	var hex string
	if strings.Index(h, "0x") == 0 {
		hex = string(h[2:])
	} else {
		hex = h
	}

	if len(h)%2 != 0 {
		hex = "0" + hex
	}

	return "0x" + hex
}

func HexToBig(h string) *big.Int {
	i := big.NewInt(0)

	h = strings.Replace(h, "0x", "", -1)
	if h == "" {
		return i
	}

	if _, ok := i.SetString(h, 16); !ok {
		panic(fmt.Sprintf("Couldn't convert hex %#v\n", h))
	}
	return i
}
