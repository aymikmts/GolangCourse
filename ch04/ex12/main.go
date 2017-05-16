// Ex12はウェブコミックxkcdのオフラインインデックスを作成します。
// 実装は"xkcd"下
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"../xkcd"
)

var dNum = flag.Int("download", -1, "comic number for downloading")
var show = flag.Bool("show", false, "show index")
var search = flag.String("search", "", "search title")

var indexFName = "./index.json" // インデックスデータファイル

func main() {
	// インデックス情報をjsonファイルから構造体にデコードする
	indexList, err := xkcd.ParseIndexList(indexFName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	flag.Parse()
	if *dNum != -1 {
		// json file is not exist, download from URL
		if !xkcd.IsExistIndex(indexList, *dNum) {
			indexData, err := xkcd.DownloadIndexData(*dNum)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v", err)
				os.Exit(1)
			}
			err = xkcd.UpdateIndex(indexData, indexList, indexFName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v", err)
				os.Exit(1)
			}
		} else {
			fmt.Printf("JSON file is already downloaded.\n")
		}
	}

	if *show {
		xkcd.ShowIndex(os.Stdout, indexList)
	}

	if *search != "" {
		xkcd.SearchIndex(os.Stdout, indexList, *search)
	}
}
