# Benchmarks

## Essentials
The following image (an original) has a square shape and a side length of <strong>1052px/side</strong>. This results in a <strong>1106704px ~ 1.1 Megapixel</strong> in total. The intrinsic size 453.947 bytes.

<p align=center><img src='./originals/test_1052.png' height=400></p>

Generally the storage capacity of indie.go yields <strong>4 bit/pixel</strong> or equivalently <strong>0.5 Bytes/pixel</strong> which apriori gives an upper bound (meaning all pixels are used) of <strong>0.5 Bytes/pixel * 1106704pixel = 553352Bytes ~ 0.55 MegaByte</strong>. Therefore the allowed storage yield (4bit/px) per pixel is independent of the image memory size, instead the total pixel size gives the upper capacity bound. However this "compression" is limited by the fact of RGB mixture as too primitive pixels (like in illustrations) will be increasingly excluded leading to less storage space after all. In practice this refers to passing a "black" image (which takes negl. space due to minimal entropy) and large pixel size to ensure high storing capacity. But a black image would create a target file which yields the whole information (low security due to easy guessing of a totally black original). Indie will skip such instances/pixels (low security) which burns available storage size in total. 

The encoding will not have significant memory impact, if any. On balance, the encoding procedure can be dessembled in taking a <strong>0.45 MB</strong> image file and encoding (up to) <strong>0.55 MB</strong> of information embedded in a target file of same size (~0.45MB).
The compression as well as the algorithmic low energy demand are reached by harvesting natural entropy from an original image. Natural distortion (iso) but also post processing backgrounds of the image can improve security significantly. 