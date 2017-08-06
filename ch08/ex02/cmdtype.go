package main

type mode string

// ASCII or BINARY
const (
	ModeASCII  mode = "ASCII"
	ModeBINARY mode = "BINARY"
)

// DataType is ASCII or BINARY mode.
var DataType mode

func (c *client) cmdType(cmds []string) error {
	switch cmds[1] {
	case "A":
		DataType = ModeASCII
	case "I":
		DataType = ModeBINARY
	default:
		DataType = ModeBINARY
	}

	msg := "type is \"" + string(DataType) + "\""
	err := c.sendResponse(statusCommandOK, msg)
	if err != nil {
		return err
	}
	return nil
}
