# Glory of the Garden
## Category: Forensics

__Description:__
> This garden contains more than it seems.

* Links to `garden.jpg`

```Console
❯ binwalk garden.jpg

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             JPEG image data, JFIF standard 1.01
382           0x17E           Copyright string: "Copyright (c) 1998 Hewlett-Packard Company"
```
 * Used `hexedit` to view the hexcode.
 * The flag is located in the ASCII section:

```Console
46 7B 6D 6F  72 65 5F 74  68 61 6E 5F
6D 33 33 74  73 5F 74 68  65 5F 33 79
33 36 35 37  42 61 42 32  43 7D 22 0A
```
__Flag:__ `picoCTF{more_than_m33ts_the_3y3657BaB2C}`
