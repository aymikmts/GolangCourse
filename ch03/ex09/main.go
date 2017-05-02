// Ex09はフラクタルをレンダリングして画像データをクライアントへ書き出します。
/* クライアントは、以下のパラメータを指定できます。
   x		: x座標
   y		: y座標
   scale	: 拡大値(例えば、"2"と指定すると2倍の大きさになる)
   color	: 色付きか否か(true/false)
   fractal	: フラクタルの種類(mandelbrot/newton))
*/
package main

import (
	//"GolangCource/ch03/mandelbrot"

	"io"
	"log"
	"net/http"
	"os"

	"../mandelbrot"
)

var out io.Writer = os.Stdout

func main() {
	http.HandleFunc("/", mandelbrot.Handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
