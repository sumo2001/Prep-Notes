# Tools on Recon




## Scope Discovery
 + WHOIS looks for the owner of a domain or IP.
    + ViewDNS.info reverse WHOIS (https://viewdns.info/reversewhois/) is a tool that searches for reverse WHOIS data by using a keyword. 
    + nslookup queries internet name servers for IP information about a host.
    + ViewDNS reverse IP (https://viewdns.info/reverseip/) looks for domains hosted on the same server, given an IP or domain.
    + crt.sh (https://crt.sh/), Censys (https://censys.io/), and Cert Spotter (https://sslmate.com/certspotter/) are platforms you can use to find certificate information about a domain.
    + Sublist3r (https://github.com/aboul3la/Sublist3r/), SubBrute (https://github.com/TheRook/subbrute/), Amass (https://github.com/OWASP/Amass/), and Gobuster (https://github.com/OJ/gobuster/) enumerate subdomains.
    + Daniel Miessler’s SecLists (https://github.com/danielmiessler/SecLists/) is a list of keywords that can be used during various phases of recon and hacking. For example, it contains lists that can be used to brute-force subdomains and filepaths.
    + Commonspeak2 (https://github.com/assetnote/commonspeak2/) generates lists that can be used to brute-force subdomains and filepaths using publicly available data. 
    + Altdns (https://github.com/infosec-au/altdns) brute-forces subdomains by using permutations of common subdomain names.
    + Nmap (https://nmap.org/) and Masscan (https://github.com/robertdavidgraham/masscan/) scan the target for open ports.
    + Shodan (https://www.shodan.io/), Censys (https://censys.io/), and Project Sonar (https://www.rapid7.com/research/project-sonar/) can be used to find services on targets without actively scanning them.
    + Dirsearch (https://github.com/maurosoria/dirsearch/) and Gobuster (https://github.com/OJ/gobuster) are directory brute-forcers used to find hidden filepaths.
    + EyeWitness (https://github.com/FortyNorthSecurity/EyeWitness/) and Snapper (https://github.com/dxa4481/Snapper/) grab screenshots of a list of URLs. They can be used to quickly scan for interesting pages among a list of enumerated paths.
## Web Hacking Reconnaissance
  + WASP ZAP (https://owasp.org/www-project-zap/) is a security tool that includes a scanner, proxy, and much more. Its web spider can be used to discover content on a web server.
  + GrayhatWarfare (https://buckets.grayhatwarfare.com/) is an online search engine you can use to find public Amazon S3 buckets.
  + Lazy s3 (https://github.com/nahamsec/lazys3/) and Bucket Stream (https://github.com/eth0izzle/bucket-stream/) brute-force buckets by using keywords.
## OSINT
  + The Google Hacking Database (https://www.exploit-db.com/google-hacking-database/) contains useful Google search terms that frequently reveal vulnerabilities or sensitive files. 
  + KeyHacks (https://github.com/streaak/keyhacks/) helps you determine whether a set of credentials is valid and learn how to use them to access the target’s services. 
  + Gitrob (https://github.com/michenriksen/gitrob/) finds potentially sensitive files that are pushed to public repositories on GitHub.
  + TruffleHog (https://github.com/trufflesecurity/truffleHog/) specializes in finding secrets in public GitHub repositories by searching for string patterns and high-entropy strings.
  + PasteHunter (https://github.com/kevthehermit/PasteHunter/) scans online paste sites for sensitive information.
  + Wayback Machine (https://archive.org/web/) is a digital archive of internet content. You can use it to find old versions of sites and their files.
  + Waybackurls (https://github.com/tomnomnom/waybackurls/) fetches URLs from the Wayback Machine.
## Tech Stack Fingerprinting
  + The CVE database (https://cve.mitre.org/cve/search_cve_list.html) contains publicly disclosed vulnerabilities. You can use its website to search for vulnerabilities that might affect your target.
  + Wappalyzer (https://www.wappalyzer.com/) identifies content management systems, frameworks, and programming languages used on a site.
  + BuiltWith (https://builtwith.com/) is a website that shows you which web technologies a website is built with.
  + StackShare (https://stackshare.io/) is an online platform that allows developers to share the tech they use. You can use it to collect information about your target.
  + Retire.js (https://retirejs.github.io/retire.js/) detects outdated JavaScript libraries and Node.js packages.
