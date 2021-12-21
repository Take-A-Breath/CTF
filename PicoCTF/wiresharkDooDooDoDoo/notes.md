# Wireshark doo dooo do doo...
## Category: Forensics

##### Description:
> Can you find the flag? shark1.pcapng
 * downloads a `.pcap` file to review in wireshark, based on the title of the challenge.

 * Filter the results using `http.response == 200`
 * Found:
   * 827	7.236537	18.222.37.134	192.168.38.104	HTTP	384	HTTP/1.1 200 OK  (text/html)
   * `Gur synt vf cvpbPGS{c33xno00_1_f33_h_qrnqorrs}`
   * Encoded?

__Flag:__ `picoCTF{p33kab00_1_s33_u_deadbeef}`
