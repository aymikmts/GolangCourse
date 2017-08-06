package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func (c *client) cmdEprt(cmds []string) (net.Conn, error) {
	param := strings.Split(cmds[1], "|")
	var msg string
	if len(param) < 4 {
		msg = "Invalid parameter: " + cmds[1]
		c.sendResponse(statusSyntaxError, msg)
		return nil, fmt.Errorf(msg)
	}
	ipAddr := "[" + param[2] + "]:" + param[3]
	log.Printf("[SERVER][EPRT]try to connect: %s\n", ipAddr)

	dataConn, err := net.Dial("tcp", ipAddr)
	if err != nil {
		return nil, err
	}

	err = c.sendResponse(statusCommandOK, "EPRT command successful.")
	if err != nil {
		return nil, err
	}
	return dataConn, nil
}
