package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// IsExistIndexは、指定されたナンバーのコミックデータが構造体にあるかどうかを返します。
func IsExistIndex(indexList *IndexList, num int) bool {
	for _, item := range indexList.Items {
		if item.Num == num {
			return true
		}
	}
	return false
}

// DownloadJSONは、指定された番号のJSONデータをxkcd.comから取得し、構造体に変換します。
func DownloadIndexData(num int) (*IndexData, error) {
	url := XKCDURL + strconv.Itoa(num) + JSONURLSuffix
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
	indexData.URL = XKCDURL + strconv.Itoa(num)
	indexData.XKCDData = &data

	return &indexData, nil
}

// UodateIndexは、インデックスデータを更新します。
func UpdateIndex(data *IndexData, indexList *IndexList, fname string) error {
	indexList.Items = append(indexList.Items, data)

	fp, err := os.Create(fname)
	defer fp.Close()
	if err != nil {
		return fmt.Errorf("failed to open file:%s: %v", fname, err)
	}

	buf, err := json.MarshalIndent(*indexList, "", "    ")
	if err != nil {
		return err
	}

	_, err = fp.Write(buf)
	if err != nil {
		return err
	}

	return nil
}
