// Ex13は、映画名を入力すると、Open Movie Databeseからポスターをダウンロードし、保存します。
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const OMDBURL = "http://www.omdbapi.com/"

type OMDBResult struct {
	Title  string
	Year   string
	Poster string
}

// SearchIssuesはGitHubのイシュートラッカーに問い合わせます。
func SearchMovie(terms []string) (*OMDBResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Printf("q:%v\nurl:%v\n", q, OMDBURL+"?t="+q)
	resp, err := http.Get(OMDBURL + "?t=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search movie failed: %s", resp.Status)
	}

	var result OMDBResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	fmt.Printf("%v\n", result)

	return &result, nil
}

func SavePoster(data *OMDBResult) error {
	resp, err := http.Get(data.Poster)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("save the poster failed: %s", resp.Status)
	}

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }

	//title := strings.TrimSpace(data.Title)
	//fname := "./" + title + ".jpg"
	fname := "./" + "tmpname" + ".jpg"
	fmt.Println(fname)
	file, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	//file.Write(body)
	io.Copy(file, resp.Body)

	return nil
}

func main() {
	data, err := SearchMovie(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = SavePoster(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
