# Easy Peasy
## Category: Cryptography

#### Description
> A one-time pad is unbreakable, but can you manage to recover the flag? (Wrap with picoCTF{})
 * Downloaded `opt.py`
 * Connected to `nc mercury.picoctf.net 20266`

```Console
❯ nc mercury.picoctf.net 20266
******************Welcome to our OTP implementation!******************
This is the encrypted flag!
5b1e564b6e415c0e394e0401384b08553a4e5c597b6d4a5c5a684d50013d6e4b

What data would you like to encrypt?
```
 * Found python script that solved it, saved as `solution.py`
```Console
❯ python3 -c "print('a'*49968);print('a'*32)" | nc mercury.picoctf.net 20266
******************Welcome to our OTP implementation!******************
This is the encrypted flag!
5b1e564b6e415c0e394e0401384b08553a4e5c597b6d4a5c5a684d50013d6e4b

<-----SNIP----->
What data would you like to encrypt? Here ya go!
0346071d3d1904593d1903573d1950033d1958592a3d1905593d1900573f3d19
```

```Console
❯ python3
Python 3.9.7 (default, Nov  3 2021, 13:00:13)
[Clang 13.0.0 (clang-1300.0.29.3)] on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>> ef=0x5b1e564b6e415c0e394e0401384b08553a4e5c597b6d4a5c5a684d50013d6e4b
>>> ea=0x0346071d3d1904593d1903573d1950033d1958592a3d1905593d1900573f3d19
>>> pa=0x6161616161616161616161616161616161616161616161616161616161616161
>>> '{:x}'.format(ef^ea^pa)
'3939303732393936653666376433393766366561303132386234353137633233'
```

* Converted to ascii using Cyberchef:
  * Flag picoCTF{99072996e6f7d397f6ea0128b4517c23}
