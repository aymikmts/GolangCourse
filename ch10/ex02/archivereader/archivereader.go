// memo: Image packageを参考に作成
// http://golang-jp.org/src/image/format.go

package archivereader

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type Reader interface {
	ShowFileList() error
}

// ErrFormat indicates that decoding encountered an unknown format.
var ErrFormat = errors.New("archivereader: unknown format")

// A format holds an image format's name, magic header, offset to magic header and how to decode it.
type format struct {
	name, magic string                       // ファイルフォーマット名、magicフィールド
	offset      int                          // magicフィールドまでのオフセット値
	decode      func(string) (Reader, error) // decode func
}

var formats []format

// RegisterFormatは、Decodeで使うアーカイブファイルフォーマットを登録します。
// nameは、"tar"や"zip"のようなフォーマット名を指します。
// magicは、マジックナンバー(フォーマット識別子)を指します。
// offsetは、マジックナンバーまでのオフセット値を指します。
// decodeは、アーカイブファイルを開くための関数です。
func RegisterFormat(name, magic string, offset int, decode func(string) (Reader, error)) {
	formats = append(formats, format{name, magic, offset, decode})
}

// A reader is an io.Reader that can also peek ahead.
type reader interface {
	io.Reader
	Peek(int) ([]byte, error)
}

// asReader converts an io.Reader to a reader.
func asReader(r io.Reader) reader {
	if rr, ok := r.(reader); ok {
		return rr
	}
	return bufio.NewReader(r)
}

// Match reports whether magic matches b. Magic may contain "?" wildcards.
func match(magic string, b []byte) bool {
	//fmt.Printf("magic:%v b:%v\n", magic, string(b))
	if len(magic) != len(b) {
		return false
	}
	for i, c := range b {
		if magic[i] != c && magic[i] != '?' {
			return false
		}
	}
	return true
}

// Sniff determines the format of r's data.
func sniff(r reader) format {
	for _, f := range formats {
		b, err := r.Peek(len(f.magic) + f.offset)
		//fmt.Printf("b:%s, %d, %v,  f.magic:%v, %v, %d\n", string(b), len(b), b[f.offset:], f.magic, []byte(f.magic), len(f.magic))
		if err == nil && match(f.magic, b[f.offset:]) {
			return f
		}
	}
	return format{}
}

func ReadArchive(fname string) (Reader, string, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	rr := asReader(file)
	f := sniff(rr)
	if f.decode == nil {
		return nil, "", ErrFormat
	}
	m, err := f.decode(fname)
	return m, f.name, err
}
