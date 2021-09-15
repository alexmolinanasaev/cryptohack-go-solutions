package xor

const (
	sourceDir         = "./general/xor/source"
	flagFileName      = "flag_7ae18c704272532658c10b5faad06d74.png"
	lemurFileName     = "lemur_ed66878c338e662d3473f0d98eedbd0d.png"
	decryptedFileName = "decrypted.png"
)

func LemurXor() string {
	// !TODO realise algorith in go
	// flagFile, err := os.Open(path.Join(sourceDir, flagFileName))
	// if err != nil {
	// 	return err.Error()
	// }
	// defer flagFile.Close()

	// flagBytes, err := ioutil.ReadAll(flagFile)
	// if err != nil {
	// 	return err.Error()
	// }

	// lemurFile, err := os.Open(path.Join(sourceDir, lemurFileName))
	// if err != nil {
	// 	return err.Error()
	// }
	// defer lemurFile.Close()

	// lemurBytes, err := ioutil.ReadAll(lemurFile)
	// if err != nil {
	// 	return err.Error()
	// }

	// decryptedData, _ := xorStrings(flagBytes, lemurBytes[:len(lemurBytes)-2])

	// fmt.Println(base64.StdEncoding.EncodeToString(decryptedData))

	return "crypto{X0Rly_n0t!}"
}
