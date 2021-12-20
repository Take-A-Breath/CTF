# tunn3l v1s10n
## Category: Forensics

#### Description
> We found this file. Recover the flag.
 * Downloaded `tunn3l_v1s10n.dms`

__What is a `.dms` file?__ [FileInfo.com](https://fileinfo.com/extension/dms)
> A DMS file is a compressed Amiga disk image created using Amiga's Disk Masher System (DMS). It contains data intended for use with an
> Amiga computer. DMS files were originally used to transfer data among Amiga users via physical floppy disks, but are now most often used
> to share Amiga games, programs, and other files online, for use with Amiga emulators such as Amiga Forever.

__How to Open__ (ibid)
> Many modern decompression programs do not support the DMS format, so you might have to download a program such as The Unarchiver (Mac)
> or WinFellow (Windows) if you want to view the files contained within a DMS archive.

 * Found a [Stack Overflow thread](https://stackoverflow.com/questions/67403845/i-extracted-a-file-using-binwalk-i-discovered-it-has-a-troc-file-how-do-i-read) that explains how to extract the files contained within:

```Console
❯ binwalk --dd=".*" tunn3l_v1s10n.dms

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
33211         0x81BB          TROC filesystem, 1263425345 file entries
948694        0xE79D6         StuffIt Deluxe Segment (data): f:IK
```

* Checked it with `hexedit`
  * First characters are 'BM', which means this is a `.bmp` file.

* Had to edit the hex header files. See [this video](https://www.youtube.com/watch?v=X4kJiQdDn7M&ab_channel=MartinCarlisle)
* Also see [This wikipedia article](https://en.wikipedia.org/wiki/BMP_file_format#Bitmap_file_header)
* Flag: `picoCTF{qu1t3_a_v13w_2020}`
