package encoding

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/alexmolinanasaev/cryptohack-go-solutions/utils"
)

// request url
const (
	hostName string = "socket.cryptohack.org"
	portNum  string = "13377"
)

// encoding types
const (
	base64Encoding = "base64"
	hexEncoding    = "hex"
	rot13Encoding  = "rot13"
	utf8Encoding   = "utf-8"
	bigintEncoding = "bigint"
)

type Encoded struct {
	T   string      `json:"type"`
	Msg interface{} `json:"encoded"`
}

type Decoded struct {
	Msg string `json:"decoded"`
}

func EncodingChallenge() string {
	cmd := ``
	conn, err := utils.ConnectTelnet(hostName, portNum)
	if err != nil {
		return fmt.Errorf("encoding challenge error: %s", err).Error()

	}
	defer conn.Close()

	// for i := 0; i < 10; i++ {
	fmt.Println("sent message to server:", cmd)

	msg, err := utils.DoDial(conn, cmd)
	if err != nil {
		fmt.Println("got empty message from server:", msg)
		// continue
	}

	fmt.Println("got message from server:", msg)

	if msg == "" {
		// fmt.Println("HERE!")
		// i -= 1
		// continue
	}
	// fmt.Println("cmd =", cmd)

	var encodedResponse Encoded
	err = json.Unmarshal([]byte(msg), &encodedResponse)
	if err != nil {
		return fmt.Errorf("cannot unmarshal server response: %s", err).Error()
	}

	decoded := `{"decoded":"`

	switch encodedResponse.T {
	case base64Encoding:
		decoded += base64Decode(encodedResponse.Msg.(string))
	case hexEncoding:
		decoded += hexDecode(encodedResponse.Msg.(string))
	case rot13Encoding:
		decoded += rot13Decode(encodedResponse.Msg.(string))
	case utf8Encoding:
		decoded += utf8Decode(encodedResponse.Msg.([]interface{}))
	case bigintEncoding:
		decoded += bigintDecode(encodedResponse.Msg.(string))
	}

	decoded += `"}`

	// fmt.Println("asd", string(response))
	// fmt.Println(decoded.Msg)
	// fmt.Println(string(response))
	// }

	fmt.Println("sent message to server:", decoded)

	msg, err = utils.DoDial(conn, decoded)
	if err != nil {
		fmt.Println("got empty message from server:", msg)
		// continue
	}

	fmt.Println("got message from server:", msg)

	return ""
}

type rot13Reader struct {
	r io.Reader
}

func decodeRune(x byte) byte {
	capital := x >= 'A' && x <= 'Z'
	if !capital && (x < 'a' || x > 'z') {
		return x // Not a letter
	}

	x += 13
	if capital && x > 'Z' || !capital && x > 'z' {
		x -= 26
	}
	return x
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i := 0; i <= n; i++ {
		b[i] = decodeRune(b[i])
	}
	return n, err
}

func base64Decode(msg string) string {
	b, _ := base64.RawStdEncoding.DecodeString(msg)
	return string(b)
}

func hexDecode(msg string) string {
	b, _ := hex.DecodeString(msg)
	return string(b)
}

func rot13Decode(msg string) string {
	buf := new(strings.Builder)
	s := strings.NewReader(msg)
	r := rot13Reader{s}
	io.Copy(buf, &r)

	return buf.String()
}

func utf8Decode(msg []interface{}) string {
	msgBytes := make([]byte, len(msg))
	for i := range msg {
		msgBytes[i] = byte(msg[i].(float64))
	}

	return string(msgBytes)
}

// ez hacked
func bigintDecode(msg string) string {
	return hexDecode(msg)
}
