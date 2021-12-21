# MacroHard WeakEdge
## Category: Forensics

#### Description:
I've hidden a flag in this file. Can you find it? Forensics is fun.pptm

* This is a hunt. Unpackaged file with binwalk and found an odd string in a file called `hidden`

```Console
❯ cat hidden
Z m x h Z z o g c G l j b 0 N U R n t E M W R f d V 9 r b j B 3 X 3 B w d H N f c l 9 6 M X A 1 f Q
```

* Stripped out the spaces in Go:
```Golang
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.ReplaceAll("Z m x h Z z o g c G l j b 0 N U R n t E M W R f d V 9 r b j B 3 X 3 B w d H N f c l 9 6 M X A 1 f Q", " ", "")
	fmt.Printf("%q", s)
}
```
* Output: ZmxhZzogcGljb0NURntEMWRfdV9rbjB3X3BwdHNfcl96MXA1fQ

```Console
❯ echo 'ZmxhZzogcGljb0NURntEMWRfdV9rbjB3X3BwdHNfcl96MXA1fQ' | base64 -d

flag: `picoCTF{D1d_u_kn0w_ppts_r_z1p5}`
