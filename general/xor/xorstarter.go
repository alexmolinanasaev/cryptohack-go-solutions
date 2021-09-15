package xor

func XorStarter() string {
	encoded := "label"
	var xorValue uint8 = 13
	decoded := make([]byte, 0)
	for _, c := range encoded {
		decoded = append(decoded, uint8(c)^xorValue)
	}

	return string(decoded)
}
