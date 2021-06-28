<h2 align=center><strong>indie</strong></h2>

A fast implementation to hide information into images. Each pixel is scanned for security and validity reasons and subsequently RGB tweaked to encode a secret text which lies in the difference between an original and target file. indie.go was developed as an alternative steganographic method with high memory capacity, negligible quality loss and high security.

### The advantage compared to password managers
Indie involves no passwords one needs to memorize! If an original image is picked, indie will draw it's information and encode the provided plain text then output a target file which looks like a copy of the original image. To decode your secret again indie will just need the original and target file path. As long as no one has the original image file, the algorithm is considered cryptographically secure by the (theorem of perfect secrecy)[https://en.wikipedia.org/wiki/One-time_pad].