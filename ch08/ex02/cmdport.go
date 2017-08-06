package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func parseAddr(addr string) (string, error) {
	nums := strings.Split(addr, ",")
	if len(nums) != 6 {
		return "", fmt.Errorf("Invalid address: %s", addr)
	}

	var port int
	p1, err := strconv.Atoi(nums[4])
	if err != nil {
		return "", err
	}
	p2, err := strconv.Atoi(nums[5])
	if err != nil {
		return "", err
	}

	port = p1*256 + p2

	ret := fmt.Sprintf("%s.%s.%s.%s:%d", nums[0], nums[1], nums[2], nums[3], port)
	fmt.Printf("IP address: %s\n", ret)
	return ret, nil
}

func (c *client) cmdPort(cmds []string) (net.Conn, error) {
	if len(cmds) < 2 {
		return nil, fmt.Errorf("PORT command needs argument")
	}

	ipAddr, err := parseAddr(cmds[1])
	if err != nil {
		return nil, err
	}
	fmt.Printf("try to connect: %s\n", ipAddr)

	dataConn, err := net.Dial("tcp", ipAddr)
	if err != nil {
		return nil, err
	}

	err = c.sendResponse(statusCommandOK, "PORT command successful.")
	if err != nil {
		return nil, err
	}
	return dataConn, nil
}
