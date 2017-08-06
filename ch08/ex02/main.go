package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/textproto"
	"strings"
)

const (
	statusFileOK                  = 150
	statusCommandOK               = 200
	statusReady                   = 220
	statusLoggedOut               = 221
	statusCloseDataConnection     = 226
	statusLoggedIn                = 230
	statusRequestedActionOK       = 250
	statusPathCreated             = 257
	statusUserOK                  = 331
	statusCmdNotImplemented       = 502
	statusRequestedActionNotTaken = 550
)

type client struct {
	conn net.Conn
	r    *textproto.Reader
}

func main() {
	fmt.Println("Start FTP server!")

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

func (c *client) sendResponse(code int, message string) error {
	res := fmt.Sprintf("%d %s\n", code, message)
	_, err := io.WriteString(c.conn, res)
	if err != nil {
		return err
	}
	return nil
}

func parseCommand(line string) []string {
	cmds := strings.Split(line, " ")
	return cmds
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Printf("%s has connected.\n", addr)

	// new client
	var c client
	c.conn = conn
	c.r = textproto.NewReader(bufio.NewReader(conn))
	err := c.sendResponse(statusReady, "Ready for FTP service.")
	if err != nil {
		log.Println(err)
		return
	}

	var dataConn net.Conn

	for {
		line, err := c.r.ReadLine()
		if err != nil {
			if err == io.EOF {
				log.Printf("Disconnected.\n")
				return
			}
			log.Println(err)
			return
		}
		fmt.Printf("CLIENT: %s\n", line)
		cmds := parseCommand(line)

		switch cmds[0] {
		case "CWD":
			err = c.cmdCd(cmds)
		case "LIST":
			err = c.cmdList(dataConn, cmds)
		case "NLST":
			err = c.cmdList(dataConn, cmds)
		case "PASS":
			fmt.Printf("User logged in.\n")
			err = c.sendResponse(statusLoggedIn, "WELCOME!")
		case "PORT":
			dataConn, err = c.cmdPort(cmds)
		case "PWD":
			err = c.cmdPwd()
		case "QUIT":
			err = c.sendResponse(statusLoggedOut, "bye")
		case "RETR":
			fmt.Println("NOT IMPLEMENTED")
		case "STOR":
			fmt.Println("NOT IMPLEMENTED")
		case "TYPE":
			err = c.cmdType(cmds)
		case "USER":
			fmt.Printf("User: %s\n", cmds[1])
			err = c.sendResponse(statusUserOK, "User mame OK. Need password.")

		default:
			msg := fmt.Sprintf("command [%v] is not implemented.", cmds[0])
			fmt.Println(msg)
			err = c.sendResponse(statusCmdNotImplemented, msg)
		}
		if err != nil {
			log.Println(err)
		}
	}
}
