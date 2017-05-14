package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// イシュー作成先。
// 今回は"aymikmts/hello-world"リポジトリに固定
const PostIssueURL = "https://api.github.com/repos/aymikmts/hello-world/issues"

var IssueNum int

type IssueEntry struct {
	Title     string
	Body      string
	User      string `json:"assignee"`
	State     string `json:"state,omitempty"`
	Milestone int
	Labels    []string
	Assignees []string
}

// PostIssueEntryはGitHubにイシューを作成します。
func PostIssueEntry(entry *IssueEntry) (string, string, error) {
	data, err := json.MarshalIndent(*entry, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("[JSON]\n%s\n", data)

	req, err := http.NewRequest(
		"POST",
		PostIssueURL,
		bytes.NewBuffer([]byte(data)),
	)
	if err != nil {
		fmt.Printf("new request:%s\n", err)
		return "", "", fmt.Errorf("failed to create new issue: %s", err)
	}

	// Content-Type設定
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		resp.Body.Close()
		fmt.Printf("crient Do:%s\n", err)
		return "", "", fmt.Errorf("failed to create new issue: %s", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("body:", string(body))
	defer resp.Body.Close()

	return resp.Status, "local", nil
}

// EditIssueEntryはGitHubにイシューを作成します。
func EditIssueEntry(entry *IssueEntry) (string, error) {
	data, err := json.MarshalIndent(*entry, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("[JSON]\n%s\n", data)

	patchUrl := PostIssueURL + "/" + strconv.Itoa(IssueNum)
	fmt.Printf("url:%s\n", patchUrl)

	req, err := http.NewRequest(
		"PATCH",
		patchUrl,
		bytes.NewBuffer([]byte(data)),
	)
	if err != nil {
		fmt.Printf("new request:%s\n", err)
		return "", fmt.Errorf("failed to edit an issue: %s", err)
	}

	// Content-Type設定
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		resp.Body.Close()
		fmt.Printf("crient Do:%s\n", err)
		return "", fmt.Errorf("failed to edit an issue: %s", err)
	}

	fmt.Println("header:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("body:", string(body))
	defer resp.Body.Close()
	return resp.Status, nil
}
