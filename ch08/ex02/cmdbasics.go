package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

func (c *client) cmdCd(cmds []string) error {
	if len(cmds) < 2 {
		return fmt.Errorf("CWD command needs argument")
	}
	err := os.Chdir(cmds[1])
	if err != nil {
		msg := "\"" + cmds[1] + "\"" + " is not exist."
		log.Printf("[SERVER][ CWD]%s\n", msg)
		err = c.sendResponse(statusRequestedActionNotTaken, msg)
		if err != nil {
			return err
		}
		return err
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Printf("[SERVER][ CWD]current directory: %s\n", pwd)

	err = c.sendResponse(statusRequestedActionOK, "CWD command successful.")
	if err != nil {
		return err
	}
	return nil
}

func (c *client) cmdList(dataConn net.Conn, cmds []string) error {
	if dataConn == nil {
		return fmt.Errorf("data connection is not established")
	}
	defer dataConn.Close()

	err := c.sendResponse(statusFileOK, "file status OK")
	if err != nil {
		return err
	}

	var args []string
	args = append(args, "-l")

	if len(cmds) > 2 {
		args = append(args, cmds[1])
	}

	out, err := exec.Command("ls", args...).Output()
	if err != nil {
		return err
	}

	dataConn.Write(out)

	log.Printf("[SERVER][LIST]:\n%s\n", string(out))

	err = c.sendResponse(statusCloseDataConnection, "complete data connection")
	if err != nil {
		return err
	}

	return nil
}

func (c *client) cmdPwd() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	msg := "\"" + dir + "\" is current directory."
	log.Printf("[SERVER][ PWD]%s", msg)
	err = c.sendResponse(statusPathCreated, msg)
	if err != nil {
		return err
	}
	return nil
}
