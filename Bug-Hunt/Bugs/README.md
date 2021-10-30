## Injections [Severity 1]
+ SQL Injection: This occurs when user controlled input is passed to SQL queries. As a result, an attacker can pass in SQL queries to manipulate the outcome of such queries. 
+ Command Injection: This occurs when user input is passed to system commands. As a result, an attacker is able to execute arbitrary system commands on application servers.

### Active Command Injection
+ Blind command injection occurs when the system command made to the server does not return the response to the user in the HTML document.  Active command injection will return the response to the user.  It can be made visible through several HTML elements.
### Defense:
+ Using an allow list: when input is sent to the server, this input is compared to a list of safe input or characters. If the input is marked as safe, then it is processed. Otherwise, it is rejected and the application throws an error.
+ Stripping input: If the input contains dangerous characters, these characters are removed before they are processed.
## Broken Authentication [Severity 2] 
+ A user would enter these credentials, the server would verify them. If they are correct, the server would then provide the users’ browser with a session cookie. A session cookie is needed because web servers use HTTP(S) to communicate which is stateless. Attaching session cookies means that the server will know who is sending what data. The server can then keep track of users' actions. 
+ Weak Session Cookies: Session cookies are how the server keeps track of users. If session cookies contain predictable values, an attacker can set their own session cookies and access users’ accounts

### Working
+ There is an existing user with the name admin and now we want to get access to their account so what we can do is try to re-register that username but with slight modification. 
+ We are going to enter " admin"(space in the starting). Now when you enter that like email id or password and submit that data. 
+ It will actually register a new user but that user will have the same right as normal admin.

### Defense
+ To avoid password guessing attacks, ensure the application enforces a strong password policy. 
+ To avoid brute force attacks, ensure that the application enforces an automatic lockout after a certain number of attempts. This would prevent an attacker from launching more brute force attacks.
+ Implement Multi Factor Authentication - If a user has multiple methods of authentication, for example, using username and passwords and receiving a code on their mobile device, then it would be difficult for an attacker to get access to both credentials to get access to their account.
## Sensitive Data Exposure [Severity 3] 
+ When a webapp accidentally divulges sensitive data, we refer to it as "Sensitive Data Exposure"
+ Attacker would force user connections through a device which they control, then take advantage of weak encryption on any transmitted data to gain access to the intercepted information 
+ My-SQL/MariaDB are "flat-file" databases, as they are stored as a single file on the computer,  format of flat-file database is an sqlite database.
## Open Redirect
![image](https://user-images.githubusercontent.com/51809378/139378895-65f2e490-2a45-43fa-b071-62821797749a.png)

## Filtering and Bypassing Open Redirects
The filtering system may use google.com as to be present in the parameter passed, so how passing google.com in any form like parameter or a subdomain helps us to get the job done!

+ http://sumo.com/login/?nextPage=https://google.com [Allowed]
+ http://sumo.com/login/?nextPage=https://hacker.com [Not Allowed]
+ http://sumo.com/login/?nextPage=https://hacker.com/?google.com [Allowed] 
#### Redirecting pages by passing the filters is the key to open redirect
+ https://sumo.com/redirect=http://www.google.com@nahasec.com   [redirects to nahamsec.com]

## XSS
![image](https://user-images.githubusercontent.com/51809378/139433794-b758d030-8403-400c-a9a3-e9cc371ab4a9.png)

#### Working of XSS
+ The basic working of XSS starts with injecting a javascript onto a browser
+ As soon as the victim encounter/executes the malicious payload
+ The malicious payload can able to perform anything that the Victim can able to do : That consists - Update Email, passw, and many more
+ If the user has some super privs, that as a hacker, will be a diwali for us  
#### Types
+ Reflected
    + Triggers only when victim clicks on the malicious link
    + ![image](https://user-images.githubusercontent.com/51809378/139435920-a1e93037-ebc5-42f4-bb6c-5248319c4358.png)
    + ![image](https://user-images.githubusercontent.com/51809378/139521754-87969834-73d3-4de3-b258-76a5ffed389b.png)

+ Stored
    + Payload will be stored on the server, users accessing the server/site are vulnerable - it pulls the payload from DB and presents to the user browsing the site
    + ![image](https://user-images.githubusercontent.com/51809378/139521749-e4766e78-3079-4adc-b262-8a2a88ec9ad8.png)


+ DOM
    + Usually arise when JavaScript takes data from an attacker-controllable source
## CSRF 
+ ![image](https://user-images.githubusercontent.com/51809378/139521981-32d78262-aa83-4e1d-a439-91d37c9c9287.png)
