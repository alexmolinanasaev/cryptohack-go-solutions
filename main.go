package main

import (
	"fmt"

	"github.com/alexmolinanasaev/cryptohack-go-solutions/general/encoding"
	"github.com/alexmolinanasaev/cryptohack-go-solutions/general/xor"
	"github.com/alexmolinanasaev/cryptohack-go-solutions/introduction"
)

func main() {
	fmt.Printf("[1.2]great snakes: %s\n", introduction.GreatSnakes())
	fmt.Printf("[1.3]network attacks: %s\n", introduction.NetworkAttacks())

	fmt.Printf("[2.1]ASCII: %s\n", encoding.ASCII())
	fmt.Printf("[2.2]Hex: %s\n", encoding.Hex())
	fmt.Printf("[2.3]Base64: %s\n", encoding.Base64())
	fmt.Printf("[2.4]Bytes and Big Integers: %s\n", encoding.BytesBigInts())

	fmt.Printf("[3.1]XOR Starter: %s\n", xor.XorStarter())
	fmt.Printf("[3.2]XOR Properties: %s\n", xor.XorProperties())
	fmt.Printf("[3.3]Favourite byte: %s\n", xor.FavByte())
	fmt.Printf("[3.4]You either know, XOR you don't: %s\n", xor.KnowYourXor())
	fmt.Printf("[3.5]Lemur XOR: %s\n", xor.LemurXor())
}
