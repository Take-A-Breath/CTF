# ARMssembly 0
## Category: Reverse Engineering

#### Description
> What integer does this program print with arguments 182476535 and 3742084308? File: chall.S Flag format:
> 	picoCTF{XXXXXXXX} -> (hex, lowercase, no 0x, and 32 bits. ex. 5614267 would be picoCTF{0055aabb})

 * Assembly code is looking for return of the greater of the two integers in hex.

```Console
❯ python3
Python 3.9.7 (default, Nov  3 2021, 13:00:13)
[Clang 13.0.0 (clang-1300.0.29.3)] on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>> '{:x}'.format(3742084308)
'df0bacd4'
```

* Flag: picoCTF{df0bacd4}
