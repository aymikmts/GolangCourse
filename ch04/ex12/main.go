// Ex12はウェブコミックxkcdのオフラインインデックスを作成します。
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	XKCDURL       = "https://xkcd.com/"
	JSONURLSuffix = "/info.0.json"
)

var dNum = flag.Int("download", -1, "comic number for downloading")
var show = flag.Bool("show", false, "show index")

var indexFName = "./index.json" // インデックスデータファイル

type IndexList struct {
	Items []*IndexData
}

type IndexData struct {
	Num      int
	URL      string
	XLCDData *XKCDData
}

type XKCDData struct {
	Num        int
	Transcript string
	Title      string
	Img        string
}

// parseIndexListは、インデックスデータのjsonファイルから構造体へデコードします。
func parseIndexList(fname string) (*IndexList, error) {
	fp, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("failed to open file:%s: %v", fname, err)
	}
	defer fp.Close()

	var result IndexList
	reader := bufio.NewReader(fp)
	if err = json.NewDecoder(reader).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// 指定されたナンバーのコミックデータが構造体にあるかどうかを返します。
func isExistIndex(indexList *IndexList, num int) bool {
	for _, item := range indexList.Items {
		if item.Num == num {
			return true
		}
	}
	return false
}

// downloadJSONは、指定された番号のJSONデータをxkcd.comから取得し、構造体に変換します。
func downloadIndexData(num int) (*IndexData, error) {
	url := XKCDURL + strconv.Itoa(num) + JSONURLSuffix
	fmt.Printf("json url: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get json: %s", resp.Status)
	}

	var data XKCDData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	var indexData IndexData
	indexData.Num = num
	indexData.URL = url
	indexData.XLCDData = &data

	return &indexData, nil
	// save as a json file
	// data, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	log.Fatalf("Failed to get JSON: %s", err)
	// }

	// fname := "./" + num + ".json"
	// err = ioutil.WriteFile(fname, data, os.ModePerm)
	// if err != nil {
	// 	log.Fatalf("Failed to save json file: %s", err)
	// }

}

// インデックスデータを更新する
func updateIndex(data *IndexData, indexList *IndexList) error {
	indexList.Items = append(indexList.Items, data)
	return nil
}

// showIndexは、ローカルにあるJSONからインデックスを作成し、表示します。
func showIndex() {
	// var jsonFiles []string

	// // jsonファイルのリストを作る
	// files, err := ioutil.ReadDir(".")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, file := range files {
	// 	if strings.HasSuffix(file.Name(), ".json") {
	// 		jsonFiles = append(jsonFiles, file.Name())
	// 	}
	// }
	// fmt.Println(jsonFiles)

	// // jsonファイル読み込み
	// for _, file := range jsonFiles {
	// 	fp, err := os.OpenFile(file, os.O_RDONLY, 0644)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	reader := bufio.NewReader(fp)

	// 	var result XKCDData
	// 	json.NewDecoder(reader).Decode(&result)

	// 	fmt.Printf("#%d\t%s\n", result.Num, result.Title)
	// }

}

func main() {

	fmt.Printf("!!! NOT IMPLEMENTED !!!\n")
	os.Exit(1)

	// インデックス情報をjsonファイルから構造体にデコードする
	indexList, err := parseIndexList(indexFName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	flag.Parse()
	if *dNum != -1 {
		// json file is not exist, download from URL
		if !isExistIndex(indexList, *dNum) {
			indexData, err := downloadIndexData(*dNum)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v", err)
				os.Exit(1)
			}
			err = updateIndex(indexData, indexList)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v", err)
				os.Exit(1)
			}
		} else {
			fmt.Printf("JSON file is already exist.\n")
		}
	}

	if *show {
		showIndex()
	}
}
