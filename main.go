/*
Indie is a steganographic program which hides text into images.
*/

// == packages ==
package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// == parameters ==
var __SIZE__ = 1

// == functions ==
func loadImage(filePath string) (image.Image, image.Config, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("indie loadImage open ERROR:", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	//reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(f))
	f, err = os.Open(filePath)
	defer f.Close()
	conf, _, e := image.DecodeConfig(f)
	if e != nil {
		fmt.Println("indie decoding ERROR:", e)
	}
	if err != nil {
		fmt.Println("indie loadImage ERROR:", err)
	}
	return img, conf, err
}

func spanImage(imageObject image.Image, configObject image.Config) (matrix [][][]uint8) {
	h, w := configObject.Height, configObject.Width
	// determine dimensions and init matrix
	matrix = make([][][]uint8, h)
	for i := 0; i < h; i++ {
		matrix[i] = make([][]uint8, w)
		for j := 0; j < w; j++ {
			matrix[i][j] = make([]uint8, 3)
		}
	}

	// fill matrix
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Println(imageObject.At(i, j).RGBA())
		}
	}

	// print info
	capacity := __SIZE__ * h * w / 8
	fmt.Println(capacity, "Bytes")
	return
}

func main() {
	img, conf, err := loadImage("indigo.jpeg")
	matrix := spanImage(img, conf)
	if err != nil {
		fmt.Println("indie ERROR:", err)
	}
	fmt.Println(matrix)
}
