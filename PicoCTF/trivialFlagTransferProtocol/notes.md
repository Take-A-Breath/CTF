# Trivial Flag Transfer Protocol
## Category: Forensics

#### Description:
> Description
> 
> Figure out how they moved the flag.
 * Provided with a `.pcapng` file
   * opened in Wireshark

#### WiresharkWiki on TFTP
 * The full post can be found [here](https://wiki.wireshark.org/TFTP)

> Trivial File Transfer Protocol (TFTP)
> 
> TFTP is used to transfer files in a very simple way.
> 
> Compared to other file transfer protcols (like: FTP or HTTP), TFTP is much simpler (and much smaller in code size) and therefore easier to
> implement. Because of this, it's often used in embedded devices (e.g. thin clients) to get files from a server at bootup time (typically
> inconjunction with BOOTP).
> 
> Sometimes TFTP is also used to upload firmware files from the user to an embedded device, but as these devices become more and more advanced,
> HTTP is more often used for this purpose today.
> 
> Protocol dependencies
> 
>     UDP: Typically, TFTP uses UDP as its transport protocol. The well known UDP port for TFTP traffic is 69.

#### In Wireshark:
* File > Export Objects > TFTP
* Save all the files:
  * instructions.txt
  * picture1.bmp
  * picture2.bmp
  * picture3.bmp
  * plan
  * program.deb

```Console
❯ cat instructions.txt
GSGCQBRFAGRAPELCGBHEGENSSVPFBJRZHFGQVFTHVFRBHESYNTGENAFSRE.SVTHERBHGNJNLGBUVQRGURSYNTNAQVJVYYPURPXONPXSBEGURCYNA

❯ cat plan
VHFRQGURCEBTENZNAQUVQVGJVGU-QHRQVYVTRAPR.PURPXBHGGURCUBGBF
```

#### Picture files:
```Console
❯ binwalk picture1.bmp

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             PC bitmap, Windows 3.x format,, 605 x 454 x 24


❯ binwalk picture2.bmp

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             PC bitmap, Windows 3.x format,, 4032 x 3024 x 24
2815484       0x2AF5FC        Broadcom header, number of sections: 793596227,
5539633       0x548731        rzip compressed data - version 87.76 (1415270489 bytes)
6120249       0x5D6339        LANCOM OEM file
8201345       0x7D2481        LANCOM firmware header, model: "QXKRYLQXKQXKQWKOUJNTIKQFIODIODJPELRGMSHMSHMSHLRGJPEHNCIODNTIRXMRXMZbWgqejuinznkwkiuiqxmlmcOPFCD:@@6?>4@?5A>5B?6A>5?<3>;2>;2>;2>;", firmware version: "JPWJ", RC74, build 87 ("OVIPWJQX")
8249741       0x7DE18D        LANCOM firmware header, model: "OVJPXMPXMPXMPXMPXMOWLOWLPXMOWLOWLOWLPXMOWLOWLQYNT\Q[eY]j^arefwjbsf`maWaU=D9/3(8:0:;1AB8=>4>>4=<2<;1<90>;2=:1=:1=:1>;2?<3?<3?<3?<", firmware version: "KPWJ", RC73, build 88 ("NUHOVIQX")
8273945       0x7E4019        LANCOM firmware header, model: "X`UU]RT\QV^SW_TU]RS[PT\QV^SV^SV^S[eYal`eqeduhdxkfzmi}pj|om{odoc`h]T[PAG<:?4:>39:0;:0=<2=<2=<2<90;8/<90?<3A>5@=4?<3?<3>;2=:1>;2?<", firmware version: "TYaV", RC77, build 95 ("PWJRYMT\")
10291544      0x9D0958        Broadcom header, number of sections: 324294729,
12727226      0xC233BA        StuffIt Deluxe Segment (data): fVefVefVefVdeUcdT`aQ_`P``Ra`R`_Q`_QbaScbTebVfbWb^Sa]R_[P[VMTOFQLCTNDYSHWQFWQFWQEWQDWQD[UH_YL`ZM_YL]WJ]WJ\VI]WJ]WJ^XK_YLc]PlfYnh[
13247747      0xCA2503        StuffIt Deluxe Segment (data): fVdeUbcS`aQ_`P_`P``PaaQ``P``P__O__O^^N^^N^^N^^N\\L[[KYYI\ZK]ZK\YJ^[L\YJZWHZWHZWHZVG[VG]XI\WHZUFWRCUPAUPAVQBWRCYTEYTEYTEXSDXSDXSD
13389886      0xCC503E        rzip compressed data - version 89.67 (1263815251 bytes)
13514042      0xCE353A        StuffIt Deluxe Segment (data): fVcdTbdT`cS^aQ\_OSWGPVEJP?KQ@V\KW]LX^M`fUjn^lo_XZJBC3JK;QQAQO@TQBTPAUPASN?RM>UPATO@TO@UPATO@TO@TO@UPAVQBUPAUPAUPAVQBUPATPARO@SPA
13654843      0xD05B3B        HPACK archive data
13840991      0xD3325F        StuffIt Deluxe Segment (data): fVgiYfiYcfVbeUadT_bR\_O\_O_bRadT`cS^aQ\_OZ]M]_O`aQ_`P_`P^_O^^N^^N^^N__O``P`^OebSb_Pc`Qb`Q__O^_O`aQbcScdTcdT^_O[\LUVFTUEWWGXYIWZJ
14459717      0xDCA345        StuffIt Deluxe Segment (data): fV`aQYZJTUEWXHYZJUUESSCWWGYWHZWH\YJa^OeaRa\MUO@[TE]TF[RDXOAaXJ[RDRI;SJ<UL>UL>UL>VM?XOAXOAWN@TK>SJ=UL?WNAUL?RI<QH;TK>VM@WNATK>RI<
14532293      0xDDBEC5        StuffIt Deluxe Segment (data): fV_`PhiYacS[^NUYIW]Lem\[eTckZw}lyzjjgXRM>LE6NE7UL>UL>VM?YPBWN@VM?VM?WN@WN@VM?RI;PG9QH:SJ<TK=TK=UL>WN@UL>TK=UL>VM?VM?UL>TK=SJ<TK=
14908154      0xE37AFA        StuffIt Deluxe Segment (data): fVop`vwguvfpqamn^kl\ttdjgX_ZKZTE^WHb[Lb[L`YJ^WH`YJb[Ld]Nc\Mb[Lb[L`YJ_XIaZKc\Mf_PjcTe^Od]Nf_PaZKc]N^YJXSDYTE\WH\WHSO@TQBb_PspaecT
15451851      0xEBC6CB        rzip compressed data - version 90.73 (1432243550 bytes)
15844847      0xF1C5EF        VMware4 disk image
24952928      0x17CC060       StuffIt Deluxe Segment (data): fdbcaZ[YOPNPQO]^\`a_WXV[\Zefddec^_]OPNMNLXYW`a_mnlmnl\][YZXRSQCDB?@>DEC@A?BCADECCDB:;9897BCACDBEFDFGE675'(&./-<=;;:9;98<:9A?>B@?


❯ binwalk picture3.bmp

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             PC bitmap, Windows 3.x format,, 807 x 605 x 24
```

#### Picture2.bmp
* Extracts multiple files with the `.sit` extension

* It appears `instructions.txt` and `plan` are both incoded in `rot13`
* Wrote a [script](https://github.com/Take-A-Breath/ROT13) in `golang` to convert them:

```Console
❯ go run main.go
Input: GSGCQBRFAGRAPELCGBHEGENSSVPFBJRZHFGQVFTHVFRBHESYNTGENAFSRE.SVTHERBHGNJNLGBUVQRGURSYNTNAQVJVYYPURPXONPXSBEGURCYNA
Output: TFTPDOESNTENCRYPTOURTRAFFICSOWEMUSTDISGUISEOURFLAGTRANSFER.FIGUREOUTAWAYTOHIDETHEFLAGANDIWILLCHECKBACKFORTHEPLAN

Input: VHFRQGURCEBTENZNAQUVQVGJVGU-QHRQVYVTRAPR.PURPXBHGGURCUBGBF
Output: IUSEDTHEPROGRAMANDHIDITWITH-DUEDILIGENCE.CHECKOUTTHEPHOTOS
```
 * There's a lot of information in `picture3.bmp`
 * Tried the following:

```Console
❯ steghide extract -sf picture3.bmp
Enter passphrase:
```
 * `-sf, --stegofile select stego file`
 * Tried password: `DUEDILIGENCE`

```Console
❯ steghide extract -sf picture3.bmp
Enter passphrase:
wrote extracted data to "flag.txt".
```

