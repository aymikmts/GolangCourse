package prettyprint

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestModifyLink(t *testing.T) {

	file, err := os.Open("test.html")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	hostName = "http://gopl.io"
	buf, err := ModifyLink(file, nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(buf)
}
