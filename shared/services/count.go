package services

import (
	"fmt"
	"math/big"
)

func Count(prevBS, currBS []byte) ([]byte, []byte) {
	prev := convertToBig(prevBS)
	curr := convertToBig(currBS)
	newCurr := count(prev, curr)
	fmt.Println(newCurr.String())
	return currBS, newCurr.Bytes()
}

func convertToBig(bs []byte) *big.Int {
	num := new(big.Int)
	return num.SetBytes(bs)
}

func count(prev, curr *big.Int) *big.Int {
	result := big.NewInt(0).Add(prev, curr)
	return result
}
