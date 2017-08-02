// [memo]
// すべてのパッケージ(ImportPathごとに)Depsを見て、引数で渡されたパッケージがあったら、
// そのときのImportPathを記録していく。
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type PackageInfo struct {
	ImportPath string
	Deps       []string
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: ./ex04 [package path]\n")
		os.Exit(1)
	}

	targetPkg, err := getPackageInfos(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(targetPkg[0].ImportPath)

	allPkgs, err := getPackageInfos("...")
	if err != nil {
		log.Fatalln(err)
	}
	// for i, p := range allPkgs {
	// 	fmt.Println(i, p.ImportPath)
	// }

	depPkgs := getDependentPackages(targetPkg[0].ImportPath, allPkgs)

	showResult(targetPkg[0].ImportPath, depPkgs)
}

func getPackageInfos(path ...string) ([]*PackageInfo, error) {

	args := []string{"list", "-e", "-json"}
	args = append(args, path...)
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		return nil, err
	}

	var pkgs []*PackageInfo
	dec := json.NewDecoder(bytes.NewReader(out))

	for {
		var pkg PackageInfo
		err = dec.Decode(&pkg)

		if err != nil {
			if err == io.EOF {
				return pkgs, nil
			}
			return nil, err
		}
		pkgs = append(pkgs, &pkg)
	}
}

func getDependentPackages(targetPkg string, allPkgs []*PackageInfo) []string {
	var depPkgs []string
	for _, pkg := range allPkgs {
		for _, dep := range pkg.Deps {
			if targetPkg == dep {
				depPkgs = append(depPkgs, pkg.ImportPath)
			}
		}
	}
	return depPkgs
}

func showResult(targetPkg string, depPkgs []string) {
	fmt.Printf("Target package path:\n\t%s\n", targetPkg)
	fmt.Printf("Dependent packages:\n")
	for _, p := range depPkgs {
		fmt.Printf(" \t%s\n", p)
	}
}
