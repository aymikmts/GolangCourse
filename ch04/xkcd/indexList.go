// xkcdはウェブコミックxkcdのオフラインインデックスを作成します。
package xkcd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const (
	XKCDURL       = "https://xkcd.com/"
	JSONURLSuffix = "/info.0.json"
)

type IndexList struct {
	Items []*IndexData
}

type IndexData struct {
	Num      int
	URL      string
	XKCDData *XKCDData
}

type XKCDData struct {
	Num        int
	Transcript string
	Title      string
	Img        string
}

// ParseIndexListは、インデックスデータのjsonファイルから構造体へデコードします。
func ParseIndexList(fname string) (*IndexList, error) {
	var result IndexList

	// インデックスデータjsonがなければ作成する
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		fp, err := os.Create(fname)
		defer fp.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to create file:%s :%v", fname, err)
		}
		return &result, nil
	}

	// インデックスデータjsonを開く
	fp, err := os.Open(fname)
	defer fp.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to open file:%s: %v", fname, err)
	}

	// jsonからアンマーシャルする
	reader := bufio.NewReader(fp)
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to open file:%s: %v", fname, err)
	}
	if len(buf) != 0 {
		err = json.Unmarshal(buf, &result)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

// ShowIndexは、ローカルにあるJSONからインデックスを作成し、表示します。
func ShowIndex(out io.Writer, indexList *IndexList) error {
	for _, item := range indexList.Items {
		fmt.Fprintf(out, "#%-5d %9.9s %.25s %.55s\n", item.Num, item.XKCDData.Title, item.URL, item.XKCDData.Transcript)
	}

	return nil
}
