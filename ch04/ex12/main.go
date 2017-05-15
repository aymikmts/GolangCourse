// Ex12はウェブコミックxkcdのオフラインインデックスを作成します。
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
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

func DownloadJSON(num string) error {
	resp, err := http.Get(XKCDURL + num + JSONURLSuffix)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("failed to get json: %s", resp.Status)
	}

	// var result XKCDResult
	// if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
	// 	resp.Body.Close()
	// 	return err
	// }
	// resp.Body.Close()

	// save as a json file
	data, err := json.MarshalIndent(resp.Body, "", "    ")
	return nil
}

func main() {

	fmt.Printf("!!! NOT IMPLEMENTED !!!\n")
	os.Exit(1)

	flag.Parse()

	if *num != "" {
		url := XKCDURL + *num + JSONURLSuffix
		fmt.Println(url)
	}
}
