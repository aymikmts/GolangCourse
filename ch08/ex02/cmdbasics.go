package main

import "os"
import "fmt"
import "net"
import "os/exec"

func (c *client) cmdCd(cmds []string) error {
	if len(cmds) < 2 {
		return fmt.Errorf("CWD command needs argument")
	}
	err := os.Chdir(cmds[1])
	if err != nil {
		msg := "\"" + cmds[1] + "\"" + " is not exist."
		err = c.sendResponse(statusRequestedActionNotTaken, msg)
		if err != nil {
			return err
		}
		return err
	}
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

	err := c.sendResponse(statusFileOK, "data connection for ***")
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

	fmt.Printf("ls command: %s\n", string(out))

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
	err = c.sendResponse(statusPathCreated, msg)
	if err != nil {
		return err
	}
	return nil
}
