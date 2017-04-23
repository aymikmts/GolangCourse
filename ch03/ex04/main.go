// Ex04は面を計算してSVGデータをクライアントに書き出します。
/* クライアントは、以下のパラメータを指定できます。
modelType	:モデルタイプ(Default/EggCase/Moguls/Saddle)
color		:グラデーション(gradient)
width		:キャンバス幅
height		:キャンバス高さ
cells		:ます目の数
xyrange		:軸の範囲
zscale		:z単位あたりの画素数
*/
package main

import (
	"log"
	"net/http"

	"../surface" // 相対パス指定
	//"GolangCourse/ch03/surface"	// GOPATH以下指定
)

func main() {
	http.HandleFunc("/", surface.Handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
