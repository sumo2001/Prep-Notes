# Books on Bug Hunt

## Scope Discovery [Helps to find Domains]
+ WHOIS and Reverse WHOIS
  + they need to supply identifying information, such as their mailing address, phone number, and email address, to a domain registrar.
  + Reverse WHOIS is extremely useful for finding obscure or internal domains not otherwise disclosed to the public.(https://viewdns.info/reversewhois/)
  + nslookup from ip addresses will also helps us to find some owner info
+ Certificate Parsing
  + An SSL certificate’s Subject Alternative Name field lets certificate owners specify additional hostnames that use the same certificate
  + crt.sh, Censys, and Cert Spotter to find certificates for a domain.
 ## Sub-Domain Enumeration [Helps to find Sub-Domains]
  + Tools like Sublist3r, SubBrute, Amass, and Gobuster can enumerate subdomains automatically with a variety of wordlists and strategies.
     + Sublist3r works by querying search engines and online subdomain databases
     + SubBrute is a brute-forcing tool that guesses possible sub-domains until it finds real ones. 
     + Amass uses a combination of DNS zone transfers, certificate parsing, search engines, and subdomain databases to find subdomains
     + Wordlists
         + https://github.com/danielmiessler/SecLists/
         + https://github.com/assetnote/commonspeak2/   [Wordlist generator]
   + Gobuster is a tool for brute-forcing to discover subdomains, directories, and files on target web servers, Its DNS mode is used for subdomain brute-forcing. In this mode, you can use the flag -d to specify the domain you want to brute-force and -w to specify the wordlist
   + Altdns (https://github.com/infosec-au/altdns/), which discovers subdomains with names that are permutations of other subdomain names
 ## Service Enumeration
 + Enumerate the services hosted on the machines
 + To find services on a machine without actively scanning it, you can use Shodan, a search engine that lets the user find machines connected to the internet.
 + Alternatives to Shodan include Censys and Project Sonar.
 ## Directory Brute-Forcing
 + You can use Dirsearch or Gobuster for directory brute-forcing. These tools use wordlists to construct URLs, and then request these URLs from a web server.
    + 200 - directory or file exists
    + 404 - directory or file doesn’t exist
    + 403 - it exists but is protected

