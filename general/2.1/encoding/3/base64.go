package encoding

import (
	"encoding/base64"
	"encoding/hex"
)

func Base64() string {
	hexEncodedString := "72bca9b68fc16ac7beeb8f849dca1d8a783e8acf9679bf9269f7bf"
	base64Encoded, _ := hex.DecodeString(hexEncodedString)

	return base64.RawStdEncoding.EncodeToString(base64Encoded)
}
