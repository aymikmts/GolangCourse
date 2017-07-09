package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type attribute struct {
	name  string
	value string
}

type element struct {
	name string
	attr []attribute
}

func main() {
	elms := parseArgs(os.Args[1:])
	if elms == nil {
		fmt.Fprintf(os.Stderr, "argument is nil.\n")
		os.Exit(1)
	}

	dec := xml.NewDecoder(os.Stdin)
	var stack []*xml.StartElement // 要素のスタック

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, &tok) // プッシュ
		case xml.EndElement:
			stack = stack[:len(stack)-1] // ポップ
		case xml.CharData:
			if containsAll(stack, elms) {
				for _, s := range stack {
					fmt.Printf("%v ", s.Name.Local)
				}
				fmt.Printf(": %s\n", tok)
			}
		}
	}
}

func parseArgs(input []string) []*element {
	if len(input) < 1 {
		return nil
	}

	var elms []*element

	var elm *element
	for _, s := range input {
		if strings.Contains(s, "=") {
			a := strings.Split(s, "=")
			if len(a) != 2 {
				fmt.Fprintf(os.Stderr, "invalid argment: %v\n", a)
				continue
			}
			attr := attribute{a[0], a[1]}
			elm.attr = append(elm.attr, attr)

			continue
		}

		if elm != nil {
			elms = append(elms, elm)
		}
		elm = &element{name: s}
	}

	elms = append(elms, elm)
	return elms
}

func containsAttr(x []xml.Attr, y []attribute) bool {
	if len(x) < len(y) {
		return false
	}

	ret := true
	for _, xAttr := range x {
		for _, yAttr := range y {
			ret = false
			if xAttr.Name.Local == yAttr.name {
				if xAttr.Value != yAttr.value {
					return false
				} else {
					ret = true
				}
			}
		}
	}
	if !ret {
		return false
	}
	return true
}

func containsAll(stack []*xml.StartElement, elms []*element) bool {
	for len(elms) <= len(stack) {
		if len(elms) == 0 {
			return true
		}
		if stack[0].Name.Local == elms[0].name {
			if containsAttr(stack[0].Attr, elms[0].attr) {
				elms = elms[1:]
			}
		}
		stack = stack[1:]
	}
	return false
}
