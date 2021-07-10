# Benchmarks

## Essentials
The following image (an original) has a square shape and a side length of <strong>1052px/side</strong>. This results in a <strong>1106704px ~ 1.1 Megapixel</strong> in total. The intrinsic size 453.947 Bytes.

<p align=center><img src='./originals/test_1052.png' height=400></p>

Generally the storage capacity of indie.go yields <strong>4 bit/pixel</strong> or equivalently <strong>0.5 Bytes/pixel</strong> which apriori gives an upper bound (meaning all pixels are used) of <strong>0.5 Bytes/pixel * 1106704 pixel = 553352 Bytes ~ 0.55 MegaByte</strong>. Therefore the allowed storage yield (4bit/px) per pixel is independent of the image memory size, instead the total pixel size gives the upper capacity bound. However this "compression"-like effect is limited by the fact of RGB mixture as too primitive pixels (like in illustrations) will be increasingly excluded leading to less storage space after all. In practice this refers to passing a "black" image (which takes negl. space due to minimal entropy) and large pixel size to ensure high storing capacity. But a black image would create a target file which yields the whole information (low security due to easy guessing of a totally black original). Indie will skip such pixels which burns available storage size in total. 

The encoding will not have significant memory impact, if any. On balance, the encoding procedure can be dessembled in taking a <strong>0.45 MB</strong> image file and encoding (up to) <strong>0.55 MB</strong> of information embedded in a target file of same size (~0.45MB).
A low energy demand is reached by harvesting natural entropy from an original image. Natural distortion (iso) but also post processing backgrounds of the image can improve security significantly.

<br>

## Speed to Size Dependence
To demonstrate a typical speed test the image above contains an appropriate amount of entropy and intrinsic storage capacity <strong>553352 Bytes</strong> (100%). The metrics below were executed on a virtual machine with 4 CPU threads.

|  Size [px] | 50  |  250 |  500 |  750 | 1052 | 2000 | 5000 |
|---|---|---|---|---|---|---|---|
| Encryption Time  [s]| .208  | .271 | .413  | .602  | .886 | 3.54  | 9.52  |
|  Storage Cap. [KB] |  1.25 | 31.3  | 125  | 281  |  553 |  2000 | 12500  | 
|  Original Size [KB] | 4.6  | 87.8  |  412 | 795  | 1600  | 4700  | 22700  |