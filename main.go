/*
Indie is a steganographic program which hides text into images.
*/

// == packages ==
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"math"
	"os"
	"os/user"
	"strconv"
	"math/rand"
)

var __version__ = "1.0.0"

// == parameters ==
var SIZE = 1
var bits = 8
var mem = int(math.Pow(2, float64(bits))) - 1

// == types ==
type Changeable interface {
	Set(x, y int, c color.Color)
}
type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var myFlags arrayFlags

// == functions ==
func ascii(s string) string {
	res := ""
	for i := 0; i < len(s); i += 8 {
		sub := s[i : i+8]
		integer, err := strconv.ParseInt(sub, 2, 64)
		check(err)
		res += string([]byte{uint8(integer)})
	}
	return res
}
func binary(s string) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}

func bitsToVector(fourBitString string) (out [3]int) {
	if fourBitString == "0000" {
		out[0] = SIZE
		out[1] = -SIZE
		out[2] = -SIZE
	} else if fourBitString == "0001" {
		out[0] = 0
		out[1] = 0
		out[2] = -SIZE
	} else if fourBitString == "0010" {
		out[0] = 0
		out[1] = 0
		out[2] = SIZE
	} else if fourBitString == "0011" {
		out[0] = 0
		out[1] = -SIZE
		out[2] = 0
	} else if fourBitString == "0100" {
		out[0] = 0
		out[1] = -SIZE
		out[2] = -SIZE
	} else if fourBitString == "0101" {
		out[0] = 0
		out[1] = -SIZE
		out[2] = SIZE
	} else if fourBitString == "0110" {
		out[0] = 0
		out[1] = SIZE
		out[2] = 0
	} else if fourBitString == "0111" {
		out[0] = 0
		out[1] = SIZE
		out[2] = -SIZE
	} else if fourBitString == "1000" {
		out[0] = 0
		out[1] = SIZE
		out[2] = SIZE
	} else if fourBitString == "1001" {
		out[0] = -SIZE
		out[1] = 0
		out[2] = 0
	} else if fourBitString == "1010" {
		out[0] = -SIZE
		out[1] = 0
		out[2] = -SIZE
	} else if fourBitString == "1011" {
		out[0] = -SIZE
		out[1] = 0
		out[2] = SIZE
	} else if fourBitString == "1100" {
		out[0] = SIZE
		out[1] = -SIZE
		out[2] = 0
	} else if fourBitString == "1101" {
		out[0] = SIZE
		out[1] = SIZE
		out[2] = 0
	} else if fourBitString == "1110" {
		out[0] = SIZE
		out[1] = SIZE
		out[2] = -SIZE
	} else if fourBitString == "1111" {
		out[0] = SIZE
		out[1] = SIZE
		out[2] = SIZE
	}
	return
}

func capacity(matrix [][][]int) (out int) {
	// returns capacity in bytes
	h, w := len(matrix), len(matrix[0])

	// convert ascii string to binary
	pix := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {

			r := matrix[i][j][0]
			g := matrix[i][j][1]
			b := matrix[i][j][2]
			// determine if pixel is too dark or too bright
			if r > mem-SIZE || r < 2 {
				// do nothing
			} else if g > mem-SIZE || g < 2 {
				// do nothing
			} else if b > mem-SIZE || b < 2 {
				// do nothing
			} else {
				pix += 1
			}
		}
	}
	bitsPerPixel := 4
	bytes := int(bitsPerPixel * pix / 8)
	return bytes
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func decode(privatePath, publicPath string) string {

	img_priv, conf_priv, err := loadImage(privatePath)
	check(err)
	img_pub, conf_pub, err := loadImage(publicPath)
	check(err)
	m1 := spanImage(img_priv, conf_priv)
	m2 := spanImage(img_pub, conf_pub)

	if len(m1) != len(m2) || len(m1[0]) != len(m2[0]) {
		fmt.Println("indie compatibility ERROR")
	}
	out := ""
	breaker := false
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m1[0]); j++ {
			v := make([]int, 3)
			v[0] = m2[i][j][0] - m1[i][j][0]
			v[1] = m2[i][j][1] - m1[i][j][1]
			v[2] = m2[i][j][2] - m1[i][j][2]
			if v[0] == 0 && v[1] == 0 && v[2] == 0 {
				// do nothing
			} else {
				chunk := vectorToBits(v)
				if chunk == "" {
					breaker = true
					break
				} else {
					out += chunk
				}
			}
		}
		if breaker {
			break
		}
	}
	//fmt.Println("decoded:", out, ascii(out))

	return out
}

func encode(filePath, targetPath, plainText string) error {

	// convert to matrix object
	img, conf, err := loadImage(filePath)
	r, g, b, _ := img.At(0, 0).RGBA()
	r, g, b = r>>8, g>>8, b>>8

	check(err)
	matrix := spanImage(img, conf)
	h, w := len(matrix), len(matrix[0])

	// check if there is enough capacity
	cap := capacity(matrix)
	size := len(plainText)
	if size > cap {
		fmt.Println("Not enough capacity (" + string(cap) + " Bytes) for this image (" + string(size) + " Bytes).")
	}

	// convert ascii string to binary
	bin := binary(plainText)
	//breaker := false
	encodedString := ""
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
				if i == h-1 {
					if j == w-5 {
						break
					}
				}
				r := matrix[i][j][0]
				g := matrix[i][j][1]
				b := matrix[i][j][2]
				// determine if pixel is too dark or too bright
				if r > mem-SIZE-1 || r < 2 {
					// do nothing
				} else if g > mem-SIZE-1 || g < 2 {
					// do nothing
				} else if b > mem-SIZE-1 || b < 2 {
					// do nothing
				} else {
					//fmt.Println(len(bin), len(encodedString), (i*w+j)*4)
					if (i*w+j+1)*4 <= len(bin) {
						//fmt.Println("entered")
						// determine vector
						chunk := bin[(i*w+j)*4 : (i*w+j+1)*4] // 4 bit chunk
						encodedString += chunk
						v := bitsToVector(chunk)

						// crop pixel
						matrix[i][j][0] = r + v[0]
						matrix[i][j][1] = g + v[1]
						matrix[i][j][2] = b + v[2]
				
					} else {
						/* add obscuring combination e.g. -SIZE,SIZE,-SIZE which have no effect */
						// sample a random vector which has not effect
						u := rand.Float64()
						// crop pixel
						if u < 0.2 {
							matrix[i][j][0] = r - SIZE
							matrix[i][j][1] = g + SIZE
							matrix[i][j][2] = b - SIZE
						} else if u < 0.4 {
							matrix[i][j][0] = r - SIZE
							matrix[i][j][1] = g - SIZE
							matrix[i][j][2] = b - SIZE
						} else if u < 0.6 {
							matrix[i][j][0] = r - SIZE
							matrix[i][j][2] = b - SIZE
						} else if u < 0.6 {
							matrix[i][j][0] = r - SIZE
							matrix[i][j][2] = b - SIZE
						} else if u < 0.8 {
							matrix[i][j][0] = r + SIZE
							matrix[i][j][2] = b - SIZE
						} else {
							matrix[i][j][0] = r + SIZE
							matrix[i][j][2] = b + SIZE
						}
						
					}
				}
			
			// if encodedString == bin {
			// 	breaker = true
			// 	break
			// }
		}
		// if breaker {
		// 	break
		// }
	}
	err = saveImage(targetPath, matrix)
	return err
}

func loadImage(filePath string) (image.Image, image.Config, error) {
	f, err := os.Open(filePath)
	check(err)
	defer f.Close()
	img, _, err := image.Decode(f)
	f, err = os.Open(filePath)
	defer f.Close()
	conf, _, e := image.DecodeConfig(f)
	check(e)
	return img, conf, err
}

func saveImage(filePath string, matrix [][][]int) error {

	// create a writable img type
	cimg := image.NewRGBA(image.Rect(0, 0, len(matrix), len(matrix[0])))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if bits == 8 {
				cimg.Set(i, j, color.RGBA{uint8(matrix[i][j][0]), uint8(matrix[i][j][1]), uint8(matrix[i][j][2]), 255})
			}

		}
	}
	x, y, z, _ := cimg.At(0, 0).RGBA()
	x, y, z = x>>8, y>>8, z>>8
	savePath, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0600)
	check(err)
	defer savePath.Close()
	err = png.Encode(savePath, cimg)
	check(err)

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
			r, g, b = r>>8, g>>8, b>>8
			// if i == 0 && j == 0 {
			// 	fmt.Println("raw RGBA:", r, g, b)
			// }

			//fmt.Printf("[X : %d Y : %v] R : %v, G : %v, B : %v, A : %v  \n", i, j, r, g, b, a)
			matrix[i][j][0] = int(r)
			matrix[i][j][1] = int(g)
			matrix[i][j][2] = int(b)
		}
	}
	return
}

func vectorToBits(v []int) (out string) {
	if v[0] == SIZE && v[1] == -SIZE && v[2] == -SIZE {
		out = "0000"
	} else if v[0] == 0 && v[1] == 0 && v[2] == -SIZE {
		out = "0001"
	} else if v[0] == 0 && v[1] == 0 && v[2] == SIZE {
		out = "0010"
	} else if v[0] == 0 && v[1] == -SIZE && v[2] == 0 {
		out = "0011"
	} else if v[0] == 0 && v[1] == -SIZE && v[2] == -SIZE {
		out = "0100"
	} else if v[0] == 0 && v[1] == -SIZE && v[2] == SIZE {
		out = "0101"
	} else if v[0] == 0 && v[1] == SIZE && v[2] == 0 {
		out = "0110"
	} else if v[0] == 0 && v[1] == SIZE && v[2] == -SIZE {
		out = "0111"
	} else if v[0] == 0 && v[1] == SIZE && v[2] == SIZE {
		out = "1000"
	} else if v[0] == -SIZE && v[1] == 0 && v[2] == 0 {
		out = "1001"
	} else if v[0] == -SIZE && v[1] == 0 && v[2] == -SIZE {
		out = "1010"
	} else if v[0] == -SIZE && v[1] == 0 && v[2] == SIZE {
		out = "1011"
	} else if v[0] == SIZE && v[1] == -SIZE && v[2] == 0 {
		out = "1100"
	} else if v[0] == SIZE && v[1] == SIZE && v[2] == 0 {
		out = "1101"
	} else if v[0] == SIZE && v[1] == SIZE && v[2] == -SIZE {
		out = "1110"
	} else if v[0] == SIZE && v[1] == SIZE && v[2] == SIZE {
		out = "1111"
	} else {
		out = ""
	}

	
	return
}

// == CL flags ==
var (
	c *bool
	o *string
	t *string
	e *bool
	d *bool
	f *string
	s *string
	w *string
	h *bool
	v *bool
)

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

// == main ==
func init() {
	c = flag.Bool("c", false, "Prints available capacity in bytes.")
	o = flag.String("o", "", "Original image path to use for encryption.")
	t = flag.String("t", "", "Specify target path (optional).")
	f = flag.String("f", "", "Draw text from file path.")
	e = flag.Bool("e", false, "Encrypt option.")
	d = flag.Bool("d", false, "Decrypt option.")
	s = flag.String("s", "", "Draw text from CL string input.")
	w = flag.String("w", "", "Write the output to a file instead of terminal.")
	h = flag.Bool("h", false, "Help - prints all options.")
	v = flag.Bool("v", false, "Outputs the current version.")
}

func main() {

	flag.Parse()
	fmt.Println("-------- DEV OUTPUT --------")
	fmt.Println("original path:", *o)
	fmt.Println("capacity:", *c)
	fmt.Println("target:", *t)
	fmt.Println("encrypt:", *e)
	fmt.Println("decrypt:", *d)
	fmt.Println("write result:", *w)
	fmt.Println("----------------------------")

	// try to load original
	if *o != "" {
		img, conf, err := loadImage(*o)
		check(err)
		matrix := spanImage(img, conf)

		if *h {
			fmt.Println("indie help options:")
			Usage()
		} else {
			if *c {
				fmt.Println("Capacity (", *o, "): ", capacity(matrix), " bytes")
			}
	
			if *d && *e {
				fmt.Println("Please choose either decrypt '-d' or encrypt '-e' flag option.")
			} else {
	
				// need to get target
				home, err := user.Current()
				check(err)
				target := string(home.HomeDir) + "/indie.png"
				if *t != "" {
					target = *t
				}
	
				if *e {
	
					fmt.Println("Encrypt text into", *t, "using original", *o, "image.")
	
					// and plain text
					plainText := ""
					if *s != "" {
						plainText = *s
					} else if *f != "" {
						content, err := ioutil.ReadFile(*f)
						check(err)
						plainText = string(content)
					}
	
					err := encode(*o, target, plainText)
					check(err)
				} else if *d {
					fmt.Println("Decrypt text from", *t, "using original", *o, "image.")
					secret := decode(*o, target)
					if *w != "" {
						ioutil.WriteFile(*w, []byte(ascii(secret)), 0644)
	
					} else {
						fmt.Println("\n------------- secret --------------\n\t", ascii(secret), "\n-----------------------------------")
					}
				}
			}
		}		
	}
}

// how to run
// go run main.go -c -o=parrot.png
