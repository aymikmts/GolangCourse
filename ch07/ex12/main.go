// Ex12はデータベースのエントリを作成、読み出し、更新、削除が可能なサーバーを立ち上げます。
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const templ = `
<h1>Item List</h1>
<table border=1>
<tr style='text-align: left'>
	<th>Item</th>
	<th>Price</th>
</tr>

{{range $key, $val := .Items}}
<tr>
	<td>{{$key}}</td>
	<td>{{$val}}</td>
</tr>
{{end}}
</table>
`

func main() {
	db := database{Items: map[string]dollars{"shoes": 50, "socks": 5}}
	http.HandleFunc("/", db.root)
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database struct {
	Items map[string]dollars
	mu    sync.Mutex
}

// rootページ
func (db *database) root(w http.ResponseWriter, req *http.Request) {
	const templ = `
	<table border=1>
	<tr style='text-align: left'><th>内容</th><th>Path</th><th>Link</th></tr>
	<tr><td>リスト表示</td><td>/list</td><td>{{.Link}}</td></tr>
	<tr><td>作成</td><td>/create?item=[item]&price=[price]</td></tr>
	<tr><td>読み込み</td><td>/read?item=[item]</td></tr>
	<tr><td>更新</td><td>/update?item=[item]&price=[price]</td></tr>
	<tr><td>削除</td><td>/delete?item=[item]</td></tr>
	</table>
	`
	t := template.Must(template.New("root").Parse(templ))
	var data struct {
		Link template.HTML
	}
	data.Link = "<a href=/list>❏</a>"
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

// リストページ
func (db *database) list(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	t := template.Must(template.New("list").Parse(templ))
	if err := t.Execute(w, db); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	db.mu.Lock()
	defer db.mu.Unlock()

	price, ok := db.Items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db *database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	db.mu.Lock()
	defer db.mu.Unlock()

	_, ok := db.Items[item]
	if ok {
		fmt.Fprintf(w, "%q is already exist.\n", item)
		return
	}

	str := req.URL.Query().Get("price")
	price, err := strconv.Atoi(str)
	if err != nil || price < 0 {
		fmt.Fprintf(w, "price is invalid value: %v\nerr: %v\n", str, err)
		return
	}

	db.Items[item] = dollars(price)
	fmt.Fprintf(w, "create %s: %s\n", item, dollars(price))
}

func (db *database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	db.mu.Lock()
	defer db.mu.Unlock()

	price, ok := db.Items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	db.mu.Lock()
	defer db.mu.Unlock()

	_, ok := db.Items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	str := req.URL.Query().Get("price")
	price, err := strconv.Atoi(str)
	if err != nil || price < 0 {
		fmt.Fprintf(w, "price is invalid value: %v\nerr: %v\n", str, err)
	}

	db.Items[item] = dollars(price)
	fmt.Fprintf(w, "update %s: %s\n", item, dollars(price))
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	db.mu.Lock()
	defer db.mu.Unlock()

	_, ok := db.Items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	delete(db.Items, item)
	fmt.Fprintf(w, "delete %q\n", item)
}
