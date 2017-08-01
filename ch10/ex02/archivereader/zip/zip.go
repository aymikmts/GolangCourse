package zip

import (
	"GolangCourse/ch10/ex02/archivereader"
	"archive/zip"
	"fmt"
)

type zipReadCloser struct {
	r *zip.ReadCloser
}

func init() {
	archivereader.RegisterFormat("zip", "PK\003\004", 0, decode)
}

func (zrc *zipReadCloser) ShowFileList() error {
	for _, f := range zrc.r.File {
		fmt.Println(f.Name)
	}

	return nil
}

func decode(fname string) (archivereader.Reader, error) {
	r, err := zip.OpenReader(fname)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return &zipReadCloser{r: r}, nil
}
