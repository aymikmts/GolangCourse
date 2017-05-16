package xkcd

import (
	"fmt"
	"io"
	"strings"
)

// ShowIndexは、ローカルにあるJSONからインデックスを作成し、表示します。
func SearchIndex(out io.Writer, indexList *IndexList, title string) error {
	for _, item := range indexList.Items {
		if strings.Contains(item.XKCDData.Title, title) {
			fmt.Fprintf(out, "-----------------------------\n")
			fmt.Fprintf(out, "URL: %s\nTranscript:\n%s\n", item.URL, item.XKCDData.Transcript)
		}
	}

	return nil

}
