package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func (c *client) cmdRetr(dataConn net.Conn, cmds []string) error {
	if dataConn == nil {
		return fmt.Errorf("data connection is not established")
	}
	defer dataConn.Close()

	msg := "opening " + string(DataType) + " mode data connection"
	err := c.sendResponse(statusFileOK, msg)
	if err != nil {
		return err
	}

	// サーバーのファイルを開く
	f, err := os.Open(cmds[1])
	if err != nil {
		c.sendResponse(statusRequestedActionNotTaken, err.Error())
		return err
	}
	defer f.Close()
	log.Printf("[SERVER][RETR]File opened: %s\n", cmds[1])

	// クライアントにデータを送信する
	size, err := io.Copy(dataConn, f)
	if err != nil {
		c.sendResponse(statusRequestedActionNotTaken, err.Error())
		return err
	}

	msg = fmt.Sprintf("complete transfer. size: %d", size)
	err = c.sendResponse(statusCloseDataConnection, msg)
	if err != nil {
		return err
	}
	log.Printf("[SERVER][RETR]%s\n", msg)

	return nil
}
