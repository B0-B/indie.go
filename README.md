<h2 align=center><strong>indie</strong></h2>

---

<p align=center>This complete README is encoded in below png file</p>

<p align=center><img text-align="center" src="target.png"></p>


A fast implementation to hide information into images. Each pixel is scanned for security and validity reasons and subsequently RGB tweaked to encode a secret text which lies in the difference between an original and target file. indie.go was developed as an alternative steganographic method with high memory capacity, negligible quality loss and high security.

### The advantage over password managers
Indie involves no passwords one needs to memorize! If an original image is picked, indie will draw it's information and encode the provided plain text then output a target file which looks like a copy of the original image. To decode your secret again indie will just need the original and target file path. Without the original image file the algorithm at hand is considered cryptographically secure by the [theorem of perfect secrecy](https://en.wikipedia.org/wiki/One-time_pad).

---

<br>

## Setup

If you are using indie with [go](https://golang.org/) make sure to have it installed.
For Ubuntu/Debian use for instance
```bash
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
```
see [here](https://golang.org/dl/) for custom OS installers.

Next, [download indie](https://github.com/B0-B/indie.go/archive/refs/heads/master.zip) or clone repository
```bash
git clone https://github.com/B0-B/indie.go.git
cd /indie.go
```

<br>

## Getting Started
The typical <strong>encryption line</strong>
```bash
go run indie.go -c -e -o /path/to/original/file.png -t /optional/target/path.png -s "Confidential Hello World!" 
```
encodes and saves the provided message string in a slightly altered copy (of the original image) which is exported to the target path.
The opt. `-c` command will scan and print the available capacity for this picture in bytes. The return should yield

```bash
Capacity ( parrot.png ):  553352  bytes
Encrypt text into /optional/target/path.png using original parrot.png image.
```

To <strong>decrypt</strong> use the `-d` flag and provide original and target path 
```bash
go run indie.go -d -o /path/to/original/file.png -t /optional/target/path.png
```
The expected output should look like this
```bash
~$ go run indie.go -d -o /path/to/original/file.png -t /optional/target/path.png

# output
Decrypt text from ./out.png using original parrot.png image.

------------- secret --------------
    Confidential Hello World! 
-----------------------------------
```
note the encoded secret at the end. Thats it!

Please visit next section for usage parameters.

## Usage

