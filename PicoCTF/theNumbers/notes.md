# The Numbers
## Category: Cryptography

##### Description
> The numbers... what do they mean?
 * Downloaded `the_numbers.png`

```Console
❯ binwalk the_numbers.png

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             PNG image, 774 x 433, 8-bit/color RGB, non-interlaced
91            0x5B            Zlib compressed data, compressed
```
 * The file shows an image with the following text:

```
16 9 3 15 3 20 6 { 20 8 5 14 21 13 2 5 18 19 13 1 19 15 14}
```
 * Decoded using [dcode.fr - letter/number cypher](https://www.dcode.fr/letter-number-cipher)
   * 16 9 3 15 3 20 6 = PICOCTF
   * 20 8 5 14 21 13 2 5 18 19  = THENUMBERS
   * 13 1 19 15 14 = MASON

__Flag:__ `PICOCTF{THENUMBERSMASON}`
