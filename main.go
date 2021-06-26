/*
Indie is a steganographic program which hides text into images.
*/

// == packages ==
package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"strings"
)

// == parameters ==
var __SIZE__ = 1

type Changeable interface {
	Set(x, y int, c color.RGBA64)
}

type error interface {
	Error() string
}

// == functions ==
func binary(s string) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}

func bitsToVector(fourBitString string) (out [3]int) {
	if len(fourBitString) != 4 {
		fmt.Println("ERROR bitsToDiffVector: input must be 4 bit string")
	} else if fourBitString == "0000" {
		out[0] = 1
		out[1] = -1
		out[2] = -1
	} else if fourBitString == "0001" {
		out[0] = 0
		out[1] = 0
		out[2] = -1
	} else if fourBitString == "0010" {
		out[0] = 0
		out[1] = 0
		out[2] = 1
	} else if fourBitString == "0011" {
		out[0] = 0
		out[1] = -1
		out[2] = 0
	} else if fourBitString == "0100" {
		out[0] = 0
		out[1] = -1
		out[2] = -1
	} else if fourBitString == "0101" {
		out[0] = 0
		out[1] = -1
		out[2] = 1
	} else if fourBitString == "0110" {
		out[0] = 0
		out[1] = 1
		out[2] = 0
	} else if fourBitString == "0111" {
		out[0] = 0
		out[1] = 1
		out[2] = -1
	} else if fourBitString == "1000" {
		out[0] = 0
		out[1] = 1
		out[2] = 1
	} else if fourBitString == "1001" {
		out[0] = -1
		out[1] = 0
		out[2] = 0
	} else if fourBitString == "1010" {
		out[0] = -1
		out[1] = 0
		out[2] = -1
	} else if fourBitString == "1011" {
		out[0] = -1
		out[1] = 0
		out[2] = 1
	} else if fourBitString == "1100" {
		out[0] = 1
		out[1] = -1
		out[2] = 0
	} else if fourBitString == "1101" {
		out[0] = 1
		out[1] = 1
		out[2] = 0
	} else if fourBitString == "1110" {
		out[0] = 1
		out[1] = 1
		out[2] = -1
	} else if fourBitString == "1111" {
		out[0] = 1
		out[1] = 1
		out[2] = 1
	}
	return
}

func capacity(matrix [][][]int) (out int) {
	// returns capacity in bytes
	h, w := len(matrix), len(matrix[0])

	// convert ascii string to binary
	c := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {

			r := matrix[i][j][0]
			g := matrix[i][j][1]
			b := matrix[i][j][2]
			// determine if pixel is too dark or too bright
			if r > 65533 || r < 2 {
				// do nothing
			} else if g > 65533 || g < 2 {
				// do nothing
			} else if b > 65533 || b < 2 {
				// do nothing
			} else {
				// determine vector
				c += 1
			}
		}
	}
	c /= 2
	return c
}

func decode(privatePath, publicPath string) string {

	img_priv, conf_priv, err := loadImage(privatePath)
	if err != nil {
		fmt.Println("indie loadImage ERROR:", err)
	}
	img_pub, conf_pub, err := loadImage(publicPath)
	if err != nil {
		fmt.Println("indie loadImage ERROR:", err)
	}
	m1 := spanImage(img_priv, conf_priv)
	m2 := spanImage(img_pub, conf_pub)

	if len(m1) != len(m2) || len(m1[0]) != len(m2[0]) {
		fmt.Println("indie compatibility ERROR")
	}
	out := ""
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m1[0]); j++ {
			v := make([]int, 3)
			v[0] = m2[i][j][0] - m1[i][j][0]
			v[1] = m2[i][j][1] - m1[i][j][1]
			v[2] = m2[i][j][2] - m1[i][j][2]
			if v[0] == 0 && v[1] == 0 && v[2] == 0 {
				// do nothing
				fmt.Println(out, v, m2[i][j], m1[i][j])
			} else {
				out += vectorToBits(v)
				fmt.Println(out, v, m2[i][j], m1[i][j])
			}

		}
	}
	fmt.Println("decoded:", out)

	return out
}

func encode(filePath, plainText string) error {

	// convert to matrix object
	img, conf, err := loadImage(filePath)
	if err != nil {
		return err
	}
	matrix := spanImage(img, conf)
	h, w := len(matrix), len(matrix[0])

	// check if there is enough capacity
	cap := capacity(matrix)
	size := len(plainText)
	if size > cap {
		x := "Not enough capacity (" + string(cap) + " Bytes) for this image (" + string(size) + " Bytes)."
		err = errors.New(x)
		return err
	}

	// convert ascii string to binary
	bin := binary(plainText)
	fmt.Println(bin)
	pixels := len(bin) / 4
	breaker := false
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i*w+j > pixels-4 {
				breaker = true
				break
			} else {
				r := matrix[i][j][0]
				g := matrix[i][j][1]
				b := matrix[i][j][2]
				// determine if pixel is too dark or too bright
				if r > 65533 || r < 2 {
					// do nothing
				} else if g > 65533 || g < 2 {
					// do nothing
				} else if b > 65533 || b < 2 {
					// do nothing
				} else {
					// determine vector
					v := bitsToVector(bin[(i*w+j)*4 : (i*w+j+1)*4])
					// crop pixel
					matrix[i][j][0] = r + v[0]
					matrix[i][j][1] = g + v[1]
					matrix[i][j][2] = b + v[2]
				}
			}
		}
		if breaker {
			break
		}
	}

	// save image
	err = saveImage(filePath, matrix)
	return err
}

func loadImage(filePath string) (image.Image, image.Config, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("indie loadImage open ERROR:", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
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

func saveImage(filePath string, matrix [][][]int) error {
	imgFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("indie saveImage ERROR:", err)
		return err
	}
	defer imgFile.Close()

	// choose right encoder
	img, err := png.Decode(imgFile)
	// img, err := jpeg.Decode(imgFile)
	// if err != nil {
	// 	img, err = png.Decode(imgFile)
	// 	if err != nil {
	// 		img, err = gif.Decode(imgFile)
	// 		if err != nil {
	// 			fmt.Println("indie Decode ERROR:", err)
	// 			return err
	// 		}
	// 	}
	// }
	if cimg, ok := img.(Changeable); ok {
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				c := matrix[i][j]
				cimg.Set(j, i, color.RGBA64{uint16(c[0]), uint16(c[1]), uint16(c[2]), 65535})
			}
		}
	}

	// save with new name
	var opt jpeg.Options
	opt.Quality = 100
	savePathArr := strings.Split(filePath, ".")
	pathArr := strings.Split(filePath, "/")
	p := ""
	for i := 0; i < len(pathArr)-1; i++ {
		p += pathArr[i] + "/"
	}
	savePath, _ := os.Create(p + "indie." + savePathArr[len(savePathArr)-1])
	err = png.Encode(savePath, img)
	if err != nil {
		fmt.Println("indie Encode ERROR:", err)
	}

	return err
}

func spanImage(imageObject image.Image, configObject image.Config) (matrix [][][]int) {
	h, w := configObject.Height, configObject.Width
	// determine dimensions and init matrix
	matrix = make([][][]int, h)
	for i := 0; i < h; i++ {
		matrix[i] = make([][]int, w)
		for j := 0; j < w; j++ {
			matrix[i][j] = make([]int, 3)
		}
	}

	// fill matrix
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			r, g, b, _ := imageObject.At(i, j).RGBA() // returns rgba each in 16 bit alpha color
			//fmt.Printf("[X : %d Y : %v] R : %v, G : %v, B : %v, A : %v  \n", i, j, r, g, b, a)
			matrix[i][j][0] = int(r)
			matrix[i][j][1] = int(g)
			matrix[i][j][2] = int(b)
		}
	}
	return
}

func vectorToBits(v []int) (out string) {
	if v[0] == 1 && v[1] == -1 && v[2] == -1 {
		out = "0000"
	} else if v[0] == 0 && v[1] == 0 && v[2] == -1 {
		out = "0001"
	} else if v[0] == 0 && v[1] == 0 && v[2] == 1 {
		out = "0010"
	} else if v[0] == 0 && v[1] == -1 && v[2] == 0 {
		out = "0011"
	} else if v[0] == 0 && v[1] == -1 && v[2] == -1 {
		out = "0100"
	} else if v[0] == 0 && v[1] == -1 && v[2] == 1 {
		out = "0101"
	} else if v[0] == 0 && v[1] == 1 && v[2] == 0 {
		out = "0110"
	} else if v[0] == 0 && v[1] == 1 && v[2] == -1 {
		out = "0111"
	} else if v[0] == 0 && v[1] == 1 && v[2] == 1 {
		out = "1000"
	} else if v[0] == -1 && v[1] == 0 && v[2] == 0 {
		out = "1001"
	} else if v[0] == -1 && v[1] == 0 && v[2] == -1 {
		out = "1010"
	} else if v[0] == -1 && v[1] == 0 && v[2] == 1 {
		out = "1011"
	} else if v[0] == 1 && v[1] == -1 && v[2] == 0 {
		out = "1100"
	} else if v[0] == 1 && v[1] == 1 && v[2] == 0 {
		out = "1101"
	} else if v[0] == 1 && v[1] == 1 && v[2] == -1 {
		out = "1110"
	} else if v[0] == 1 && v[1] == 1 && v[2] == 1 {
		out = "1111"
	}
	return
}

func main() {
	file := "parrot.png"
	err := encode(file, "Hello World!")
	if err != nil {
		fmt.Println("Error in encoding:", err)
	} else {
		fmt.Println(file, "encoded.")
	}
	decode(file, "indie.png")
	if err != nil {
		fmt.Println("Error in decoding:", err)
	}
}
