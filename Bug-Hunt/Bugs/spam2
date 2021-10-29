![image](https://user-images.githubusercontent.com/51809378/139378895-65f2e490-2a45-43fa-b071-62821797749a.png)

## Filtering and Bypassing Open Redirects
The filtering system may use google.com as to be present in the parameter passed, so how passing google.com in any form like parameter or a subdomain helps us to get the job done!

+ http://sumo.com/login/?nextPage=https://google.com [Allowed]
+ http://sumo.com/login/?nextPage=https://hacker.com [Not Allowed]
+ http://sumo.com/login/?nextPage=https://hacker.com/?google.com [Allowed] 
#### Redirecting pages by passing the filters is the key to open redirect
+ https://sumo.com/redirect=http://www.google.com@nahasec.com   [redirects to nahamsec.com]
