package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// open "test.jpg"
	path := "./"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	pic := []string{".jpg", ".png"}
	if os.Mkdir("resize", os.ModePerm); err != nil {
		fmt.Println("resize"+"文件夹创建失败：", err.Error())
	} else {
		fmt.Println("resize" + "文件夹创建成功！")
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		for _, v := range pic {
			if !strings.HasSuffix(f.Name(), v) {
				continue
			}
			//todo 逻辑
			fmt.Println(f.Name())

			file, err := os.Open(f.Name())
			if err != nil {
				log.Fatal(err)
			}

			// decode jpeg into image.Image
			img, err := jpeg.Decode(file)
			if err != nil {
				log.Fatal(err)
			}
			file.Close()

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

	}
	return

}
