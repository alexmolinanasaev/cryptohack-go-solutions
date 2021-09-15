package xor

import (
	"encoding/hex"
)

func KnowYourXor() string {
	encrypted := "0e0b213f26041e480b26217f27342e175d0e070a3c5b103e2526217f27342e175d0e077e263451150104"
	k := "crypto{"
	encodedBytes, _ := hex.DecodeString(encrypted)

	unxoredKey, _ := xorStrings(encodedBytes[:7], []byte(k))
	unxoredKey = append(unxoredKey, 'y')

	x := len(encodedBytes)/len(unxoredKey) + 1
	repeatedUnxoredKey := []byte{}
	for i := 0; i < x; i++ {
		repeatedUnxoredKey = append(repeatedUnxoredKey, unxoredKey...)
	}

	decrypted, _ := xorStrings(repeatedUnxoredKey[:len([]byte(encodedBytes))], encodedBytes)

	return string(decrypted)
}
