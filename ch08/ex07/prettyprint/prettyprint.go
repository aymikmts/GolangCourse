package prettyprint

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"os"

	"golang.org/x/net/html"
)

var buf *bytes.Buffer
var hostName string

func ModifyLink(body io.Reader, targetHost *url.URL) (*bytes.Buffer, error) {
	if targetHost != nil {
		hostName = targetHost.String()
		//fmt.Printf("ModifyLink: %s\n", hostName)
	} else {
		fmt.Fprintf(os.Stderr, "targetHost is nil.\n")
	}

	buf = new(bytes.Buffer)
	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}

	forEachNode(doc, startNode, endNode)

	return buf, nil
}

func forEachNode(n *html.Node,
	pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startNode(n *html.Node) {
	switch n.Type {
	case html.TextNode:
		startTextNode(n)
		return
	case html.ElementNode:
		startElementNode(n)
		return
	default:
		return
	}
}

func startTextNode(n *html.Node) {
	fmt.Fprintf(buf, "%s", n.Data)
}

func startElementNode(n *html.Node) {
	depth++
	if n.FirstChild == nil {
		return
	}

	attrs := attributes(n.Attr)
	if attrs == "" {
		fmt.Fprintf(buf, "\n%*s<%s>", depth, "", n.Data)
	} else {
		fmt.Fprintf(buf, "\n%*s<%s %s>", depth, "", n.Data, attrs)
	}
}

func endNode(n *html.Node) {
	switch n.Type {
	case html.TextNode:
		return
	case html.ElementNode:
		endElementNode(n)
		return
	default:
		return
	}
}

func endElementNode(n *html.Node) {
	if n.FirstChild == nil {
		attrs := attributes(n.Attr)
		if attrs == "" {
			switch n.Data {
			case "br":
				fmt.Fprintf(buf, "<%s/>\n", n.Data)
			default:
				fmt.Fprintf(buf, "\n%*s<%s />", depth, "", n.Data)
			}
		} else {
			fmt.Fprintf(buf, "\n%*s<%s %s />", depth, "", n.Data, attrs)
		}
	} else {
		switch n.Data {
		case "a", "code", "title", "tt", "h1":
			fmt.Fprintf(buf, "</%s>", n.Data)
		default:
			fmt.Fprintf(buf, "\n%*s</%s>", depth, "", n.Data)
		}
	}
	depth--
}

func attributes(attr []html.Attribute) string {
	var b bytes.Buffer

	for i, a := range attr {
		if i != 0 {
			b.WriteString(" ")
		}

		val := a.Val
		// if a.Key == "href" {
		// 	//fmt.Printf("attributes: %s %s\n", a.Key, a.Val)
		// 	if strings.Contains(val, hostName) {
		// 		//fmt.Printf("targetHost[%s] contains: %s\n", hostName, val)
		// 		val = strings.Replace(val, hostName+"/", "", 1)
		// 		//fmt.Printf("%v\n", val)
		// 		if path.Ext(val) == "" {
		// 			val += "index.html"
		// 		}
		// 		fmt.Printf("URL: %s -> %s\n", a.Val, val)
		// 	}
		// }
		if a.Namespace == "" {
			b.WriteString(a.Key)
			b.WriteString(`="`)
			b.WriteString(val)
			b.WriteString(`"`)
		} else {
			b.WriteString(a.Namespace)
			b.WriteString(":")
			b.WriteString(a.Key)
			b.WriteString(`="`)
			b.WriteString(val)
			b.WriteString(`"`)
		}
	}
	return b.String()
}
