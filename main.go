package main

import (
	"fmt"

	ascii "github.com/alexmolinanasaev/cryptohack-go-solutions/general/2.1/encoding/1"
	hex "github.com/alexmolinanasaev/cryptohack-go-solutions/general/2.1/encoding/2"
	base64 "github.com/alexmolinanasaev/cryptohack-go-solutions/general/2.1/encoding/3"
	bytesbigints "github.com/alexmolinanasaev/cryptohack-go-solutions/general/2.1/encoding/4"
	encodingchallenge "github.com/alexmolinanasaev/cryptohack-go-solutions/general/2.1/encoding/5"

	greatsnakes "github.com/alexmolinanasaev/cryptohack-go-solutions/introduction/1.2"
	networkattacks "github.com/alexmolinanasaev/cryptohack-go-solutions/introduction/1.3"
)

func main() {
	fmt.Printf("[1.2]great snakes: %s\n", greatsnakes.GreatSnakes())
	fmt.Printf("[1.3]network attacks: %s\n", networkattacks.NetworkAttacks())

	fmt.Printf("[2.1]ASCII: %s\n", ascii.ASCII())
	fmt.Printf("[2.2]Hex: %s\n", hex.Hex())
	fmt.Printf("[2.3]Base64: %s\n", base64.Base64())
	fmt.Printf("[2.4]Bytes and Big Integers: %s\n", bytesbigints.BytesBigInts())
	// fmt.Printf("[2.5]Encoding challenge: %s\n", encodingchallenge.EncodingChallenge())
	fmt.Println(encodingchallenge.EncodingChallenge())
}
