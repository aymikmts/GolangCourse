// magicの参考: http://www.redout.net/data/tar.html#magic

package tar

import (
	"GolangCourse/ch10/ex02/archivereader"
	"archive/tar"
	"fmt"
	"io"
	"os"
)

type tarReader struct {
	r *tar.Reader
}

func init() {
	archivereader.RegisterFormat("tar", "ustar", 257, decode)
}

func (tr *tarReader) ShowFileList() error {
	for {
		header, err := tr.r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(header.Name)
	}
	return nil
}

func decode(fname string) (archivereader.Reader, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}

	r := tar.NewReader(f)
	return &tarReader{r: r}, nil
}
