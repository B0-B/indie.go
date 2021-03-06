<h1 align=center><strong>indie CLI</strong></h1>


<br>

<p align=center>This complete README is hidden <br> in the target.png file below</p>
<p align=center><img text-align="center" src="target.png"></p>
<p align=center>without the original readme.png <br>this could have also been a lie <br>as none could ever tell. </p>
<p align=center><br>Convince yourself & decrypt it<br> <strong>go run indie.go -d -o readme.png -t target.png</strong></p>
<br>

[![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Hide%20your%20secrets%20in%20images!&url=https://github.com/B0-B/indie.go&hashtags=cli,golang,encoding,images,steganography,encryption-decryption)
---

A new implementation to hide information into images with high speed. Each pixel is scanned for security and validity reasons and subsequently RGB tweaked to encode a secret text which lies in the difference between an original and target file. indie.go was developed as an alternative steganographic method with high memory capacity, negligible quality loss and high security.

### The advantage over password managers
Indie involves no passwords one needs to memorize! If an original image is picked, indie will draw it's information and encode the provided plain text then output a target file which looks like a copy of the original image. To decode your secret again indie will just need the original and target file path. Without the original image file the algorithm at hand is considered cryptographically secure by the [theorem of perfect secrecy](https://en.wikipedia.org/wiki/One-time_pad).

---

<br>

# Branches
- `master` Release Branch
- `benchmark` Performance Benchmark Branch

<br>

# Performance Benchmarks

| Capacity [bit/pixel] | Speed  |   |   |   |
|---|---|---|---|---|
| 4 |  - |   |   |   |

<br>

# Setup

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

# Getting Started
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
note the encoded secret at the end. Thats it! Please visit next section for usage parameters.

<br>

# Usage
See the options for indie below. Note that the order is not necessary but is just a recommendation.
The options list is obtained by `-h` flag.

```bash
~/indie.go$ go run indie.go -h

 ____  _  _  ____  ____  ____ 
(_  _)( \( )(  _ \(_  _)( ___)
 _)(_  )  (  )(_) )_)(_  )__) 
(____)(_)\_)(____/(____)(____)
           
help options:

        go run indie.go [-h/-c/-vb/-e/-v] [-o original file] [-t target file] [-f/-w opt. output path]

options within brackets can be used simulataneously

  -c    Prints available capacity in bytes.
  -d    Decrypt option.
  -e    Encrypt option.
  -f string
        Draw text from file path.
  -h    Help - prints all options.
  -o string
        Original image path to use for encryption.
  -s string
        Draw text from CL string input.
  -t string
        Specify target path (optional).
  -v    Outputs the current version.
  -vb
        Verbose information. Helpful for debugging.
  -w string
        Write the output to a file instead of terminal.
```

<br>

# Upcoming
* New format encoders for jpg and gif compatibility