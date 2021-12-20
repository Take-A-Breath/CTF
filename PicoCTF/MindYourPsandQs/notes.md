```Console
❯ cat values
Decrypt my super sick RSA:
c: 240986837130071017759137533082982207147971245672412893755780400885108149004760496
n: 831416828080417866340504968188990032810316193533653516022175784399720141076262857
e: 65537
```

* Found tool to possibly solve this on [Github](https://github.com/Ganapati/RsaCtfTool)

```Console
❯ python3 RsaCtfTool.py --createpub -n 831416828080417866340504968188990032810316193533653516022175784399720141076262857 -e 65537
private argument is not set, the private key will not be displayed, even if recovered.
-----BEGIN PUBLIC KEY-----
MD0wDQYJKoZIhvcNAQEBBQADLAAwKQIiHAxBgH1RH59SN9ErcR9hpjwmLb4HDGp/
04uYDgJkQgtvyQIDAQAB
-----END PUBLIC KEY-----
```
[From StackOverflow](https://stackoverflow.com/questions/49878381/rsa-decryption-using-only-n-e-and-c)

`python RsaCtfTool -n 58900433780152059829684181006276669633073820320761216330291745734792546625247 -e 65537 --uncipher 56191946659070299323432594589209132754159316947267240359739328886944131258862`
 * Cannot run because requirements can't be installed for the python package

 * Tried to decode on [decode.fr](https://www.dcode.fr/rsa-cipher)
   * Decrypted as: picoCTF{sma11_N_n0_g0od_23540368}
