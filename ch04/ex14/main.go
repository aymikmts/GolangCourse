// Ex14はバグレポート、マイルストーン、ユーザの一覧を閲覧可能にするウェブサーバーを作成します。
// Githubへの問い合わせ部分は"github/github.go"
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	//"GolangCourse/ch04/github"
	"../github"
)

// HTMLテンプレート
const templ = `
<html>
<body>
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>Milestone</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td>{{.Milestone | milestoneVal}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
</body>
</html>
`

// Milestone Objectのnilチェックをし、nilでなければパラメータを返す。
func milestoneVal(obj *github.Milestone) template.HTML {
	if obj == nil {
		return "null"
	}

	return template.HTML(fmt.Sprintf("<a href='%s'>%s</a>", obj.HTMLURL, obj.Title))
}

var report = template.Must(template.New("issuelist").Funcs(template.FuncMap{"milestoneVal": milestoneVal}).Parse(templ))

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
