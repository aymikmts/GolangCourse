// Ex11は、コマンドラインからGitHubのイシューを作成、読み出し、更新、クローズします。
// GithubへのPOST部分は"github/ex11.go"に実装します。
package main

import (
	"GolangCourse/ch04/github"
	"flag"
	"fmt"
	"log"
	"os"
)

var create = flag.Bool("create", false, "create new issue.")
var edit = flag.Int("edit", 0, "number of issue for edit")

var title = flag.String("t", "new title", "title for the issue")
var body = flag.String("b", "empty", "body for the issue")
var user = flag.String("u", "aymikmts", "user name")

func main() {

	fmt.Printf("!!! NOT IMPLEMENTED !!!\n")
	os.Exit(1)

	flag.Parse()

	// issueの作成
	if *create {
		var entry github.IssueEntry
		entry.Title = *title
		entry.Body = *body
		entry.User = *user

		status, location, err := github.PostIssueEntry(&entry)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("status: %s\nlocation:%s", status, location)
	}
	if *edit != 0 {
		github.IssueNum = *edit

		var entry github.IssueEntry
		entry.Title = *title
		entry.Body = *body
		entry.User = *user

		status, err := github.EditIssueEntry(&entry)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("status: %s\n", status)
	}
}
