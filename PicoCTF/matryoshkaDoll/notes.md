# Matryoshka doll
### Category: Forensics

#### Description:
> Matryoshka dolls are a set of wooden dolls of decreasing size placed one inside another. What's the final one?
 * provides image: `dolls.jpg`

```Console
❯ binwalk -e dolls.jpg

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             PNG image, 594 x 1104, 8-bit/color RGBA, non-interlaced
3226          0xC9A           TIFF image data, big-endian, offset of first image directory: 8
272492        0x4286C         Zip archive data, at least v2.0 to extract, compressed size: 378956, uncompressed size: 383938, name: base_images/2_c.jpg
651614        0x9F15E         End of Zip archive, footer length: 22
```

* used `binwalk -e` to extract the files from the original `.jpg` image, then every image file into every directory until I found the flag.

* Directory: `/extractedDolls/base_images/_2_c.jpg.extracted/base_images/_3_c.jpg.extracted/base_images/_4_c.jpg.extracted`

```Console
❯ cat flag.txt
picoCTF{336cf6d51c9d9774fd37196c1d7320ff}%
```
