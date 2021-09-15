package utils

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"regexp"
)

var flagRegexp = regexp.MustCompile(`crypto{([\s\S]+?)}`)

func ConnectTelnet(hostName, port string) (net.Conn, error) {
	conn, err := net.Dial("tcp", hostName+":"+port)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to %s: %s", hostName, err)
	}

	return conn, nil
}

func DoDial(conn net.Conn, cmd string) (string, error) {
	fmt.Fprintf(conn, cmd+"\n")

	res := ""
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				res = res + message
				break
			} else {
				return "", fmt.Errorf("cannot read response: %s", err)
			}
		}
		res = res + message
	}

	return res, nil
}

func ParseFlag(msg string) (string, bool) {
	matches := flagRegexp.FindAllString(msg, -1)
	if matches == nil || len(matches) <= 0 {
		return "", false
	}

	return matches[0], true
}
