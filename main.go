/*
Indie is a steganographic program which hides cypher into images.
*/

// == packages ==
package main
import (
    "os"
	"fmt"	
	"image" 	
    _"image/jpeg"
    _"image/png" 
)

// == functions ==
func loadImage (filePath string) (image.Image, error) {
    f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    img, _, err := image.Decode(f)
    return img, err
}

func main () {
    img,_ := loadImage("indigo.jpeg")
    fmt.Println(img)
}