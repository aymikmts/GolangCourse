package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type locationSet struct {
	location string
	url      string
}

type worldClocks struct {
	index int
	clock string
}

func main() {
	if len(os.Args) == 1 {
		os.Exit(0)
	}

	out := make(chan worldClocks)

	args := os.Args[1:]
	locationSets, err := parseArg(args) // locationとurlを分離
	if err != nil {
		log.Fatal(err)
	}
	for i, ls := range locationSets {
		go getWorldClock(out, ls, i)
	}

	for {
		showWorldClocks(out, len(args))
	}
}

// parseArgはargからlocationとurlを分離します。
func parseArg(args []string) ([]*locationSet, error) {
	clocks := []*locationSet{}

	for _, timeZone := range args {
		var location locationSet
		v := strings.Split(timeZone, "=")
		if len(v) != 2 {
			return nil, fmt.Errorf("argument should be \"[TZ]=[localhost:port]\".")
		}
		location.location = v[0]
		location.url = v[1]
		clocks = append(clocks, &location)
	}

	return clocks, nil
}

func getWorldClock(out chan<- worldClocks, ls *locationSet, index int) {
	conn, err := net.Dial("tcp", ls.url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		var clocks worldClocks
		reader := bufio.NewReader(conn)
		bytes, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		time := fmt.Sprintf("%s %s", ls.location, string(bytes))
		clocks.clock = time
		clocks.index = index
		out <- clocks
	}
}

func showWorldClocks(clocks <-chan worldClocks, num int) {
	time := make([]string, num)

	for i := 0; i < num; i++ {
		clock := <-clocks
		time[clock.index] = clock.clock
	}
	fmt.Println(strings.Join(time, "  "))
}
