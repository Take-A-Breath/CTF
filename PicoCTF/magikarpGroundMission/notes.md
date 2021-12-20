# Magikarp Ground Mission
## Tags: General Skills

#### Description:
> Do you know how to move between directories and read files in the shell? Start the container, `ssh` to it, and then `ls` once connected
> to begin. Login via `ssh` as `ctf-player` with the password, `abcba9f7`

```Console
❯ ssh ctf-player@venus.picoctf.net -p 55079
ctf-player@venus.picoctf.net's password:
Welcome to Ubuntu 18.04.5 LTS (GNU/Linux 5.4.0-1041-aws x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage
This system has been minimized by removing packages and content that are
not required on a system that users do not log into.

To restore this content, you can run the 'unminimize' command.
Last login: Mon Dec 20 21:04:16 2021 from 127.0.0.1
ctf-player@pico-chall$
```
 * Presented with two files in current directory. One with the first part of the flag, the other with instructions to get the next part:

```Console
ctf-player@pico-chall$ ls
1of3.flag.txt  instructions-to-2of3.txt
ctf-player@pico-chall$ cat 1of3.flag.txt
picoCTF{xxsh_
```
 * Instructions to part 2 of 3 of the flag:
```Console
ctf-player@pico-chall$ cat instructions-to-2of3.txt
Next, go to the root of all things, more succinctly `/
```
 * Change directory to root:
```console
ctf-player@pico-chall$ ls
2of3.flag.txt  dev   instructions-to-3of3.txt  media  proc  sbin  tmp
bin	       etc   lib		       mnt    root  srv   usr
boot	       home  lib64		       opt    run   sys   var
ctf-player@pico-chall$ cat 2of3.flag.txt
0ut_0f_\/\/4t3r_
```
 * Instructions for the 3rd part of the flag:
```Console
ctf-player@pico-chall$ cat instructions-to-3of3.txt
Lastly, ctf-player, go home... more succinctly `~`
```
 * changed directory to `~/`
```Console
ctf-player@pico-chall$ cd ~/
ctf-player@pico-chall$ ls
3of3.flag.txt  drop-in
ctf-player@pico-chall$ cat 3of3.flag.txt
21cac893}
```

