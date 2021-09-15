package introduction

func GreatSnakes() string {
	ords := []byte{81, 64, 75, 66, 70, 93, 73, 72, 1, 92, 109, 2, 84, 109, 66, 75, 70, 90, 2, 92, 79}

	res := ""

	for _, o := range ords {
		res += string(o ^ 0x32)
	}

	return res
}
