// Ex12はウェブコミックxkcdのオフラインインデックスを作成します。
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const XKCDURL = "https://xkcd.com/"
const JSONURLSuffix = "/info.0.json"

var num = flag.String("n", "", "comic number")

type XKCDResult struct {
	Num        int
	Transcript string
	Title      string
	Img        string
}

// downloadJSONは、指定された番号のJSONデータをxkcd.comから取得し、保存します。
func downloadJSON(num string) error {
	resp, err := http.Get(XKCDURL + num + JSONURLSuffix)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("failed to get json: %s", resp.Status)
	}

	// save as a json file
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalf("Failed to get JSON: %s", err)
	}

	fname := "./" + num + ".json"
	err = ioutil.WriteFile(fname, data, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to save json file: %s", err)
	}

	return nil
}

// showIndexは、ローカルにあるJSONからインデックスを作成し、表示します。
func showIndex() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			fmt.Println(file.Name())
		}
	}
}

func main() {

	fmt.Printf("!!! NOT IMPLEMENTED !!!\n")
	os.Exit(1)

	flag.Parse()

	if *num != "" {
		// json file is not exist, download from URL
		if _, err := os.Stat("./" + *num + ".json"); err != nil {
			err := downloadJSON(*num)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v", err)
			}
		} else {
			fmt.Printf("JSON file is already exist.\n")
		}
	}

	showIndex()

}
