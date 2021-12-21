# Speeds and Feeds
## Category: Reverse Engineering

##### Description:
> There is something on my shop network running at nc mercury.picoctf.net 59953, but I can't tell what it is. Can you?
 * Googled the first line of output: `G17 G21 G40 G90 G64 P0.003 F50`
 * Linux CNC G-Code

* From [LinuxCNC.org](http://linuxcnc.org/)
> LinuxCNC controls CNC machines. It can drive milling machines, lathes, 3D printers, laser cutters, plasma cutters, robot arms, hexapods,
> and more.
> 
>     Runs under Linux (optionally with realtime extensions).
> 
>     Simple installation on Debian and Ubuntu, or via our Live/Install DVD/USB images.
> 
>     Accepts G-code input, drives CNC machines in response.
> 
>     Active user community.
> 
>     Several different GUIs available.
> 
>     Compatible with many popular machine control hardware interfaces.
> 
>     Supports rigid tapping, cutter compensation, and many other advanced control features.
> 
>     Full source code available under the terms of the GNU GPLv2 (General Public License version 2).

 * Found a g-code simulator at [NCviewer.com](https://ncviewer.com/)
   * Outputs a 3D rendering of the flag

__Flag:__ `picoCTF{num3r1cal_c0ntr0l_f3fea95b}`
