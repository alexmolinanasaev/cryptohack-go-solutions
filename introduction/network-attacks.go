package introduction

import (
	"fmt"

	"github.com/alexmolinanasaev/cryptohack-go-solutions/utils"
)

const (
	hostName string = "socket.cryptohack.org"
	portNum  string = "11112"
)

func NetworkAttacks() string {
	cmd := `{"buy": "flag"}`
	conn, err := utils.ConnectTelnet(hostName, portNum)
	if err != nil {
		return fmt.Errorf("network attacks error: %s", err).Error()

	}
	defer conn.Close()
	msg, err := utils.DoDial(conn, cmd)
	if err != nil {
		return fmt.Errorf("network attacks error: %s", err).Error()
	}

	res, found := utils.ParseFlag(msg)
	if !found {
		return "network attacks flag not found"
	}

	return res
}
