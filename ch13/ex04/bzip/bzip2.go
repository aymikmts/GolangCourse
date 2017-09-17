// パッケージ bzip は、bzip2圧縮(bzip.org)を使うライターを提供します。
package bzip

import (
	"io"
	"os/exec"
	"sync"
)

type writer struct {
	sync.Mutex                // ex03 で追加
	w          io.WriteCloser // 基底の出力ストリーム
	wg         sync.WaitGroup
}

// NewWriter はbzip2の圧縮ストリーム用のライターを返します。
func NewWriter(out io.Writer) (io.WriteCloser, error) {
	var w writer
	cmd := exec.Command("/usr/bin/bzip2")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	w.w = stdin

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err = cmd.Start(); err != nil {
		return nil, err
	}
	w.wg.Add(1)
	go func() {
		io.Copy(out, stdout)
		w.wg.Done()
	}()

	return &w, nil
}

func (w *writer) Write(data []byte) (int, error) {
	// ex03 で追加
	w.Lock()
	defer w.Unlock()
	//

	var total int

	for len(data) > 0 {
		n, err := w.w.Write(data)
		if err != nil {
			return total + n, err
		}
		total += n
		data = data[total:]
	}
	return total, nil
}

// Close は圧縮データをすべてはき出して、ストリームを閉じます。
// 基底のio.Writerは閉じません。
func (w *writer) Close() error {
	// ex03 で追加
	w.Lock()
	defer w.Unlock()
	//

	w.w.Close()
	w.wg.Wait()
	return nil
}
