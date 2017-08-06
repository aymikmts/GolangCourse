package main

func (c *client) cmdType(cmds []string) error {
	var mode string
	switch cmds[1] {
	case "A":
		mode = "ASCII"
	case "I":
		mode = "BINARY"
	default:
		mode = "BINARY"
	}

	msg := "type is \"" + mode + "\""
	err := c.sendResponse(statusCommandOK, msg)
	if err != nil {
		return err
	}
	return nil
}
