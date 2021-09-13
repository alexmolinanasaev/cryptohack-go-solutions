package encoding

import (
	"math/big"
)

func BytesBigInts() string {
	bingintDecodedAsString := "2829794709130443211351941280429255759507699000037731628396"
	bb, _ := big.NewInt(0).SetString(bingintDecodedAsString, 10)

	return string(bb.Bytes())
}
