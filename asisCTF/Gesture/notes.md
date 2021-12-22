# Gesture
## Category: Forensics
### ASIS CTF

#### Description:
A **[gesture](https://asisctf.com/tasks/gesture.img_36fb98da9a7832a88dc6fdd82dba65d94f0294c5.txz)** is a file that is intended to indicate or emphasize something!

**Note:** `flag` is lowercase string, put the flag in `ASIS{flag}` too.

Downloaded the file and decompressed.

Ran `binwalk` on the provided file.

```Console
❯ binwalk gesture.img

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             YAFFS filesystem, little endian
```

From [https://yaffs.net](Yaffs.net)

>Yaffs (Yet Another Flash File System) is an open-source file system specifically designed to be fast, robust and suitable for embedded use with NAND and NOR Flash. It is widely used with Linux, RTOSs, or no OS at all, in consumer devices, avionics, and critical infrastructure. It is available under GNU Public License, GPL, or on commercial terms from Aleph One.

Found tool to "unyaff" the file on [github](https://github.com/justsoso8/yaffs2utils)
```Console
total 60
drwxr-xr-x 15 root        root        4096 Dec 21 12:21 .
drwxr-xr-x  3 takeabreath takeabreath 4096 Dec 21 12:21 ..
drwxr-xr-x  2 takeabreath takeabreath 4096 Oct 20 22:00 ADM
drwxr-xr-x  2 takeabreath takeabreath 4096 May  7  2021 Alarms
drwxr-xr-x  3 takeabreath takeabreath 4096 May  7  2021 Android
drwxr-xr-x  3 takeabreath takeabreath 4096 Oct 20 20:39 DCIM
drwxr-xr-x  2 takeabreath takeabreath 4096 Oct 20 18:54 Documents
drwxr-xr-x  2 takeabreath takeabreath 4096 Oct 20 22:03 Download
drwxr-xr-x  2 takeabreath takeabreath 4096 May  7  2021 Movies
drwxr-xr-x  2 takeabreath takeabreath 4096 May  7  2021 Music
drwxr-xr-x  2 takeabreath takeabreath 4096 May  7  2021 Notifications
drwxr-xr-x  2 takeabreath takeabreath 4096 Oct 20 20:46 Pictures
drwxr-xr-x  2 takeabreath takeabreath 4096 May  7  2021 Podcasts
drwxr-xr-x  2 takeabreath takeabreath 4096 May  7  2021 Ringtones
drwxr-xr-x  3 takeabreath takeabreath 4096 Oct 20 19:26 Signal
```

Found a file `latentorism` in the directory `Android/data/com.android.gallery3d/cache/.nomedia`

Changed file extension to `.jpg` and opened it. 
![[Pasted image 20211221183228.png]]

```
70900 49992 98182
87494 78427 17955
```
 Nothing immediately seemed to decrypt this. Back to searching.
 
Found a backup file in the `Signal` directory. 
>  Backup decryption steps
> 
> In the subsections below we'll describe the low-level process of decrypting a backup file. The process is described using snippets of Python 3 code which have the following dependencies:
> 
>     protobuf: The Google Protocol Buffers Python support library.
>     cryptography: A library containing the various cryptographic functions we'll need.
> 
> You can install these using pip as follows:
> 
> $ pip install protobuf cryptography
> 
> In the examples below, we'll assume that the backup file has been opened in binary mode like so:
> 
> >>> backup_file = open("signal-0000-00-00-00-00-00.backup", "rb")
> 
> And the passphrase has been stored in a string like so (spaces are optional):
> 
> >>> passphrase = "00000 00000 00000 00000 00000 00000"

It would appear the numbers above may be the passphrase for the signal backup file. This can be decrypted using [signal-backup-decode](https://github.com/pajowu/signal-backup-decode)

```Console
└─$ signal-backup-decode --password '709004999298182874947842717955' signal-2021-10-21-08-01-13.backup 
18:53:35 [INFO] Output path: signal-2021-10-21-08-01-13
18:53:35 [INFO] Input file: signal-2021-10-21-08-01-13.backup
18:53:35 [INFO] Database Version: 119
             Bytes read: [00:00:00] [##################################################] 3.62MB/3.62MB
Read vs. written frames: [00:00:00] [##################################################]   417/417  
```

Has the following contents:
```Console
└─$ ls -la                       
total 448
drwxr-xr-x 5 takeabreath takeabreath   4096 Dec 21 12:53 .
drwxr-xr-x 3 takeabreath takeabreath   4096 Dec 21 12:53 ..
drwxr-xr-x 2 takeabreath takeabreath   4096 Dec 21 12:53 attachment
drwxr-xr-x 2 takeabreath takeabreath   4096 Dec 21 12:53 preference
-rw-r--r-- 1 takeabreath takeabreath 438272 Dec 21 12:53 signal_backup.db
drwxr-xr-x 2 takeabreath takeabreath   4096 Dec 21 12:53 sticker
```

Opened `signal_backup.db` with `sqlitebrowser`
```SQL
SELECT * FROM SMS
```

Found a cell with what looks like morse code

```
-. ...-- ...- . .-. ..--.- --... .-. ..- ..... - ..--.- - .---- -. . ..--.- ... . ...-- ..--.- .. ..--.- ....- ..--.- -. ----- - .... .---- -. -.... ..--.- .. ..... ..--.- .-.. ----- ----- ..--.- .--. ...-- .-. -.-. . -. --... ..--.- ... ...-- -.-. ..- .-. .
```

Decoded with [this website](https://morsecode.world/international/translator.html)

Found the flag:
N3VER#7RU5T#T1NE#SE3#I#4#N0TH1N6#I5#L00#P3RCEN7#S3CURE

The website is expecting underscores and in lowercase:
Flag: 
`ASIS{n3ver_7ru5t_t1ne_se3_i_4_n0th1n6_i5_l00_p3rcen7_s3cure}`