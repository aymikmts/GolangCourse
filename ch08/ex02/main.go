package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/textproto"
	"os"
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
	statusSyntaxError             = 500
	statusCmdNotImplemented       = 502
	statusRequestedActionNotTaken = 550
)

type client struct {
	conn net.Conn
	r    *textproto.Reader
}

func main() {
	fmt.Println("Start FTP server!")

	listener, err := net.Listen("tcp", ":21")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
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

	if len(cmds) > 1 {
		args := strings.Join(cmds[1:], " ")
		log.Printf("[CLIENT][%s]%s\n", cmds[0], args)
	} else {
		log.Printf("[CLIENT][%s]\n", cmds[0])
	}
	return cmds
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("root directory: %s\n", rootDir)

	addr := conn.RemoteAddr().String()
	log.Printf("[SERVER][    ]\"%s\" has connected.\n", addr)

	// new client
	var c client
	c.conn = conn
	c.r = textproto.NewReader(bufio.NewReader(conn))
	err = c.sendResponse(statusReady, "Ready for FTP service.")
	if err != nil {
		log.Println(err)
		return
	}

	var dataConn net.Conn

	for {
		line, err := c.r.ReadLine()
		if err != nil {
			if err == io.EOF {
				// TODO: クライアントが切断されたとき、カレントディレクトリに戻る
				err := os.Chdir(rootDir)
				if err != nil {
					msg := "\"" + rootDir + "\"" + " is not exist."
					log.Printf("[SERVER][ CWD]%s\n", msg)
					return
				}
				log.Printf("[SERVER][    ]\"%s\" has disconnected.\n", addr)
				return
			}
			log.Println(err)
			return
		}
		cmds := parseCommand(line)

		switch cmds[0] {
		case "CDUP":
			cd := []string{"", ".."}
			err = c.cmdCd(cd)
		case "CWD":
			err = c.cmdCd(cmds)
		case "EPRT":
			dataConn, err = c.cmdEprt(cmds)
		case "LIST":
			err = c.cmdList(dataConn, cmds, true)
		case "NLST":
			err = c.cmdList(dataConn, cmds, false)
		case "PASS":
			log.Printf("[SERVER][PASS]User logged in.\n")
			err = c.sendResponse(statusLoggedIn, "WELCOME!")
		case "PORT":
			dataConn, err = c.cmdPort(cmds)
		case "PWD":
			err = c.cmdPwd()
		case "QUIT":
			err = c.sendResponse(statusLoggedOut, "BYE!")
		case "RETR":
			c.cmdRetr(dataConn, cmds)
		case "STOR":
			c.cmdStor(dataConn, cmds)
		case "TYPE":
			err = c.cmdType(cmds)
		case "USER":
			log.Printf("[SERVER][USER]User: %s\n", cmds[1])
			err = c.sendResponse(statusUserOK, "User mame OK. Need password.")

		default:
			msg := fmt.Sprintf("command [%v] is not implemented.", cmds[0])
			log.Printf("[SERVER][    ]%s\n", msg)
			err = c.sendResponse(statusCmdNotImplemented, msg)
		}
		if err != nil {
			log.Println(err)
		}
	}
}
