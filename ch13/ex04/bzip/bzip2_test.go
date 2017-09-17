package bzip_test

import (
	"bytes"
	"compress/bzip2" // reader
	"io"
	"sync"
	"testing"

	"GolangCourse/ch13/ex04/bzip" // writer
)

// ex03 で追加
func TestBzip2InParallel(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := bzip.NewWriter(&compressed)

	// 1 million回同じメッセージを書き込む
	tee := io.MultiWriter(&uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "hello")
	}

	// 1 million回同じメッセージを並列的に書き込む
	tee = io.MultiWriter(w)
	var wg sync.WaitGroup
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			io.WriteString(tee, "hello")
			wg.Done()
		}()
	}
	wg.Wait()

	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// 圧縮されたストリームのサイズをチェック
	if got, want := compressed.Len(), 255; got != want {
		t.Errorf("1 million hellos compressed to %d bytes, want %d", got, want)
	}

	// オリジナルと圧縮されていないものとを比較する
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Error("decompression yielded a different message")
	}
}

// gopl.io/ch13/bzip/bzip2_test.go
func TestBzip2(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := bzip.NewWriter(&compressed)

	// 1 million回同じメッセージを書き込む。
	tee := io.MultiWriter(w, &uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "hello")
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// 圧縮されたストリームのサイズをチェック
	if got, want := compressed.Len(), 255; got != want {
		t.Errorf("1 million hellos compressed to %d bytes, want %d", got, want)
	}

	// オリジナルと圧縮されていないものとを比較する
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Errorf("decompression yielded a different message")
	}
}
