// Githubは、GitHubのイシュートラッカーに対するGoのAPIを提供します。
// https://developer.github.com/v3/search/#search-issuesを参照のこと。
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	Milestone *Milestone // ex14用に追加
	CreatedAt time.Time  `json:"created_at"`
	Body      string     // マークダウン(Markdown)形式
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// ex14用に追加
type Milestone struct {
	Title   string
	HTMLURL string `json:"html_url"`
}

// SearchIssuesはGitHubのイシュートラッカーに問い合わせます。
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// すべての実行パスでresp.Bodyを閉じなければなりません。
	// (この処理を簡単にする'defer'を第5章で説明しています。))
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
