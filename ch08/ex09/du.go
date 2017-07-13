package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")
var showdir = flag.Bool("d", false, "show each directory size")

type dirData struct {
	name      string
	fileSizes int64
}

func main() {
	// 最初のディレクトリを決める。
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// ファイルツリーのそれぞれのルートを並列に操作する。
	fileSizes := make(chan int64)
	//dirs := make(chan dirData)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
		//go walkDir(root, &n, dirs)
	}
	go func() {
		n.Wait()
		close(fileSizes)
		//close(dirs)
	}()

	// 定期的に結果を表示する
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
	//var dirname string
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			//case dir, ok := <-dirs:
			if !ok {
				break loop // fileSizesが閉じられた
			}
			nfiles++
			nbytes += size
			//nbytes += dir.fileSizes
			//dirname = dir.name
		case <-tick:
			printDiskUsage(nfiles, nbytes)
			//printDiskUsage(dirname, nfiles, nbytes)
			//		case <-dir:
			//			printDirUsage(dir)
		}
	}
	//printDiskUsage("total", nfiles, nbytes)
	printDiskUsage(nfiles, nbytes)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	//func walkDir(dir string, n *sync.WaitGroup, data chan<- dirData) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
			//go walkDir(subdir, n, data)
		} else {
			//var elm dirData
			//elm.name = entry.Name()
			//elm.fileSizes = entry.Size()
			fileSizes <- entry.Size()
			//data <- elm
		}
	}
}

// semaは、direntsでの平行性を制限するための計数セマフォです。
var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // tokenを獲得
	defer func() { <-sema }() // tokenを開放

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.3f GB\n", nfiles, float64(nbytes)/1e9)
}

//func printDiskUsage(name string, nfiles, nbytes int64) {
//	fmt.Printf("%s: %d files  %.3f GB\n", name, nfiles, float64(nbytes)/1e9)
//}
