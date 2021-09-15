package xor

import (
	"encoding/hex"
	"fmt"
)

func XorProperties() string {
	key1 := "a6c8b6733c9b22de7bc0253266a3867df55acde8635e19c73313"

	op1Res := "37dcb292030faa90d07eec17e3b1c6d8daf94c35d4c9191a5e1e"
	op2Res := "c1545756687e7573db23aa1c3452a098b71a7fbf0fddddde5fc1"
	op3Res := "04ee9855208a2cd59091d04767ae47963170d1660df7f56f5faf"

	op1Bytes, _ := hex.DecodeString(op1Res)
	op2Bytes, _ := hex.DecodeString(op2Res)
	op3Bytes, _ := hex.DecodeString(op3Res)

	key1Bytes, _ := hex.DecodeString(key1)
	key2Bytes, _ := xorStrings(key1Bytes, op1Bytes)
	key3Bytes, _ := xorStrings(key2Bytes, op2Bytes)

	flag, _ := xorStrings(key1Bytes, key2Bytes)
	flag, _ = xorStrings(flag, key3Bytes)
	flag, _ = xorStrings(flag, op3Bytes)

	return string(flag)
}

func xorStrings(key1, key2 []byte) ([]byte, error) {
	if len(key1) != len(key2) {
		return nil, fmt.Errorf("keys have different len")
	}

	res := make([]byte, len(key1))

	for i := range key1 {
		res[i] = key1[i] ^ key2[i]
	}

	return res, nil
}
