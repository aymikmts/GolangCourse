// jpegコマンドは標準入力からPNG画像を読み込んで、
// 標準出力へJPEG画像を書き出します。
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

var format = flag.String("f", "jpeg", "output image format. You can choose \"jpeg\", \"png\" and \"gif\".")

func main() {
	flag.Parse()

	img, kind, err := readImg(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	// 出力フォーマットの決定
	var f string
	if *format == "" { // フラグなしだったら、入力のフォーマットにする
		f = kind
	} else {
		f = *format
	}

	err = outImg(f, img, os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}

func readImg(in io.Reader) (image.Image, string, error) {
	img, kind, err := image.Decode(in)
	if err != nil {
		return nil, "", err
	}
	fmt.Fprintf(os.Stderr, "Input format = %v\n", kind)
	return img, kind, nil
}

func outImg(f string, img image.Image, out io.Writer) error {
	var err error
	switch f {
	case "jpeg":
		err = toJPEG(img, out)
	case "png":
		err = toPNG(img, out)
	case "gif":
		err = toGIF(img, out)
	default:
		return fmt.Errorf("Invalid output format: %s", f)
	}
	if err != nil {
		return fmt.Errorf("Failed to output image: %v", err)
	}
	return nil
}

func toJPEG(img image.Image, out io.Writer) error {
	err := jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "Output format = jpeg\n")
	return nil
}

func toPNG(img image.Image, out io.Writer) error {
	err := png.Encode(out, img)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "Output format = png\n")
	return nil
}

func toGIF(img image.Image, out io.Writer) error {
	err := gif.Encode(out, img, &gif.Options{256, nil, nil})
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "Output format = gif\n")
	return nil
}
