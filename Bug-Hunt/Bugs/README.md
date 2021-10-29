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

+ Stored
    + Payload will be stored on the server, users accessing the server/site are vulnerable - it pulls the payload from DB and presents to the user browsing the site
+ DOM
    + Usually arise when JavaScript takes data from an attacker-controllable source
