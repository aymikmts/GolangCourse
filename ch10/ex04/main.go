package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type List struct {
	Deps []string
}

func main() {
	// 	// 引数でパスを指定
	// 	// "go list -json "引数"" でjsonファイルを取得
	// 	// jsonファイルをParse
	// 	// jsonの"Imports"を[]stringに入れる
	// 	// []stringを再帰的にgo listする

	// 	// 参考： ch8/crawl
	// 	// json: ch04/ex13 (poster), ch4/movie

	if len(os.Args) == 1 {
		fmt.Printf("usage: ./ex04 [package path]\n")
		os.Exit(1)
	}

	worklist := make(chan []string)
	var n int
	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, path := range list {
			if !seen[path] {
				seen[path] = true
				n++
				go func(path string) {
					worklist <- crawl(path)
				}(path)
			}
		}
	}

	// out, err := exec.Command("go", "list", "-json", os.Args[1]).Output()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// //fmt.Println(string(out))

	// var list List
	// err = json.Unmarshal(out, &list)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//fmt.Println(list.Imports)
}

var tokens = make(chan struct{}, 20)

func crawl(path string) []string {
	fmt.Println(path)
	tokens <- struct{}{}
	out, err := exec.Command("go", "list", "-json", path).Output()
	if err != nil {
		log.Fatalln(err)
	}

	var list List
	err = json.Unmarshal(out, &list)
	if err != nil {
		log.Fatalln(err)
	}
	<-tokens
	return list.Deps
}

// func main(){
// 	// 引数でパスを指定
// 	// "go list -json "引数"" でjsonファイルを取得
// 	// jsonファイルをParse
// 	// jsonの"Imports"を[]stringに入れる
// 	// []stringを再帰的にgo listする

// 	// 参考： ch8/crawl
// 	// json: ch04/ex13 (poster), ch4/movie

// 	if len(os.Args) == 1 {
// 		log.Fatalf("need an augment.\n usage: ./ex04 [package path]\n")
// 	}

// 	command := "go list " + os.Args[1]
// 	out, err := exec.Command(command, "-json").Output()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(out)
// }
