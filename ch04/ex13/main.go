// Ex13は、映画名を入力すると、Open Movie Databeseからポスターをダウンロードします。
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"unicode"
)

const OMDBURL = "http://www.omdbapi.com/"

type OMDBResult struct {
	Title    string
	Year     string
	Poster   string
	Response string // 映画情報の有無
}

// SearchMovieはOMDBに映画情報を問い合わせます。
func SearchMovie(terms []string) (*OMDBResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(OMDBURL + "?t=" + q)
	if err != nil {
		resp.Body.Close()
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

	return &result, nil
}

// SavePosterはポスターを保存します。
func SavePoster(movieData *OMDBResult) error {

	// 映画情報がなかったときはerror
	if movieData.Response == "False" {
		return fmt.Errorf("Movie is not found.")
	}

	// JSONデータをGet
	resp, err := http.Get(movieData.Poster)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("save the poster failed: %s", resp.Status)
	}

	// 画像データの保存
	file, err := os.Create(imageName(movieData.Title))
	if err != nil {
		return err
	}
	defer file.Close()

	io.Copy(file, resp.Body)

	return nil
}

// imageNameは、タイトルから画像のファイル名を設定します。
func imageName(title string) string {
	var ret []rune

	// アルファベットと数値を抽出
	s := []rune(title)
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			ret = append(ret, r)
		}
	}

	return "./" + string(ret) + ".jpg"
}

func main() {
	movieData, err := SearchMovie(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	err = SavePoster(movieData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
