package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// open "test.jpg"
	path := "./"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir("resize", os.ModePerm); err != nil {
		fmt.Println("resize"+"文件夹创建失败：", err.Error())
	} else {
		fmt.Println("resize" + "文件夹创建成功！")
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		format := filepath.Ext(f.Name())

		file, err := os.Open(f.Name())

		var img image.Image
		switch format {
		case ".jpeg":
			img, err = jpeg.Decode(file)
		case ".jpg":
			img, err = jpeg.Decode(file)
		case ".png":
			img, err = png.Decode(file)
		case ".gif":
			img, err = gif.Decode(file)
		default:
			continue
		}

		// decode jpeg into image.Image
		//img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio

		kuan := img.Bounds().Dx() //
		gao := img.Bounds().Dy()  //
		fmt.Println(kuan)
		fmt.Println(gao)

		m := resize.Resize(uint(kuan), 0, img, resize.Lanczos3)
		//m := resize.Resize(1000, 0, img, resize.Lanczos3)

		out, err := os.Create("resize/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
		fmt.Println("------------------")

	}
	return

}
