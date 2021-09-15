package xor

import (
	"encoding/hex"

	"github.com/alexmolinanasaev/cryptohack-go-solutions/utils"
)

func FavByte() string {
	encoded := "73626960647f6b206821204f21254f7d694f7624662065622127234f726927756d"
	encodedBytes, _ := hex.DecodeString(encoded)

	var secretByte uint8 = 0
	var maxSecretIterations int = 20

	for i := 0; i < maxSecretIterations; i++ {
		decoded := xorString(encodedBytes, secretByte)
		flag, found := utils.ParseFlag(decoded)
		if found {
			return flag
		}

		secretByte += 1
	}

	return ""
}

func xorString(encoded []byte, secretByte byte) string {
	decoded := ""
	for _, c := range encoded {
		decoded += string(rune(c ^ secretByte))
	}

	return decoded
}
