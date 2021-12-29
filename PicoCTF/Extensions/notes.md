# Extensions
## Category: Forensics

#### Description:
> This is a really weird text file TXT? Can you find the flag?

#### Flag.txt
```Console
❯ file flag.txt
flag.txt: PNG image data, 1697 x 608, 8-bit/color RGB, non-interlaced

❯ binwalk flag.txt

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             PNG image, 1697 x 608, 8-bit/color RGB, non-interlaced
91            0x5B            Zlib compressed data, compressed
```
* Changed extension to `.png` and opened the file.
* The image has the flag:
  * picoCTF{now_you_know_about_extensions}
