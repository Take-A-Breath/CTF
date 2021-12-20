> Using tabcomplete in the Terminal will add years to your life, esp. when dealing with long rambling directory structures and filenames:
 * File Addadshashanammu.zip

```Console
❯ unzip Addadshashanammu.zip
Archive:  Addadshashanammu.zip
   creating: Addadshashanammu/
   creating: Addadshashanammu/Almurbalarammi/
   creating: Addadshashanammu/Almurbalarammi/Ashalmimilkala/
   creating: Addadshashanammu/Almurbalarammi/Ashalmimilkala/Assurnabitashpi/
   creating: Addadshashanammu/Almurbalarammi/Ashalmimilkala/Assurnabitashpi/Maelkashishi/
   creating: Addadshashanammu/Almurbalarammi/Ashalmimilkala/Assurnabitashpi/Maelkashishi/Onnissiralis/
   creating: Addadshashanammu/Almurbalarammi/Ashalmimilkala/Assurnabitashpi/Maelkashishi/Onnissiralis/Ularradallaku/
  inflating: Addadshashanammu/Almurbalarammi/Ashalmimilkala/Assurnabitashpi/Maelkashishi/Onnissiralis/Ularradallaku/fang-of-haynekhtnamet
```

* cd through the directories until reaching file `fang-of-haynekhtnamet`
  * run strings on the file to find the flag:

```Console
<-----SNIP----->
AWAVI
AUATL
[]A\A]A^A_
*ZAP!* picoCTF{l3v3l_up!_t4k3_4_r35t!_d32e018c}
;*3$"
GCC: (Ubuntu 7.5.0-3ubuntu1~18.04) 7.5.0
crtstuff.c
<-----SNIP----->
```
