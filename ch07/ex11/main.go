// Ex11はデータベースのエントリを作成、読み出し、更新、削除が可能なサーバーを立ち上げます。
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{items: map[string]dollars{"shoes": 50, "socks": 5}}
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
	items map[string]dollars
	mu    sync.Mutex
}

func (db *database) root(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "/list\n")
	fmt.Fprintf(w, "/price?item=[item]\n")
	fmt.Fprintf(w, "/create?item=[item]&price=[price]\n")
	fmt.Fprintf(w, "/read?item=[item]\n")
	fmt.Fprintf(w, "/update?item=[item]&price=[price]\n")
	fmt.Fprintf(w, "/delete?item=[item]\n")
}

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	for item, price := range db.items {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	db.mu.Lock()
	defer db.mu.Unlock()

	price, ok := db.items[item]
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

	_, ok := db.items[item]
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

	db.items[item] = dollars(price)
	fmt.Fprintf(w, "create %s: %s\n", item, dollars(price))
}

func (db *database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	db.mu.Lock()
	defer db.mu.Unlock()

	price, ok := db.items[item]
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

	_, ok := db.items[item]
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

	db.items[item] = dollars(price)
	fmt.Fprintf(w, "update %s: %s\n", item, dollars(price))
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	db.mu.Lock()
	defer db.mu.Unlock()

	_, ok := db.items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	delete(db.items, item)
	fmt.Fprintf(w, "delete %q\n", item)
}
