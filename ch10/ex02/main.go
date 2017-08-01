package main

import (
	"GolangCourse/ch10/ex02/archivereader"
	_ "GolangCourse/ch10/ex02/archivereader/tar"
	_ "GolangCourse/ch10/ex02/archivereader/zip"
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("need an argument.\n")
		os.Exit(1)
	}

	r, kind, err := archivereader.ReadArchive(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	if r == nil {
		log.Fatalf("failed to read. reader is nil.")
	}

	fmt.Printf("-- Input format --\n%s\n", kind)

	fmt.Printf("\n-- File List-- \n")
	err = r.ShowFileList()
	if err != nil {
		log.Fatalln(err)
	}
}
