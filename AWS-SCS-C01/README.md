## AWS Config
-  Asses, Audit and evaluate your AWS Assets continously
-  Monitors and records your AWS **resource Configurations** and allows you to evaluate automatically
-  Can trigger an SNS topic upon some changes -> continuous monitoring
-  Can define rules for compaliance checks eg: conformance pack, deviaition can trigger SNS
-  Able to track the relationships among resources and review resource dependencies prior to making changes 
-  Can get History of resource configurations for troubleshooting
-  You can obtain the details of the event API call that invoked the change (e.g., who made the request, at what time, and from which IP address) from the CloudTrail logs.
-  ![image](https://user-images.githubusercontent.com/51809378/155887540-461758d9-09b0-4073-9f43-f1e0feda7e73.png)
-  to monitor access, AWS Config uses CloudTrail
-  to enable it in all region you have to do it manually
-  requires IAM role (with Read only permissions to the all resources, Write access to S3 logging bucket, Publish access to SNS)
-  allows resource tracking, compliance, auditing, security analysis
-  AWS Config sends notifications for the following events:
     - Configuration item change for a resource.
     - Configuration history for a resource was delivered for your account.
     - Configuration snapshot for recorded resources was started and delivered for your account.
     - Compliance state of your resources and whether they are compliant with your rules.
     - Evaluation started for a rule against your resources.
     - AWS Config failed to deliver the notification to your account.
## AWS CloudTrail
- It records all API calls of your accoung and delivers logs to an endpoint
- It contains a full 90-day record of activity
- Great Logging and Monitoring Service
- Doesnt support all the AWS Resources
- By Defualt Cloud trail event Logs are Encrypted using AWS SSE
- Logs are delivered for every 5mins 
- Can prevent logs deletion by configuring s3 MFA delete
- Log file intergrity is enabled by default
- Every hour log files are delivered with 'digest' file to validate the log's integrity [SHA256]
## AWS Cloud Watch
- for metrics and alarms and notifications and storing logs in S3
- Allows us to configure rules based on events
- Logs can be pushed from your on prem an also from AWS services
- Cloud watch support Multi-region and Single Region Api calls
## AWS VPC Flow Logs
- Requires a proprt IAM role to push logs tocloudwatch
- You cannot enable VPC Flow logs outside of your AWS account
- DNS, DHCP, Meta traffic are not recorded in VPC flow logs
- You cant change a configuration after creating a flow log
- Simple Email Service - by default EC2 throttles traffic over port 25. Better to use port 587 or 2587 SMTP
## IAM
- A global user and role based management 
### Policies
- Policies are case senstive
- Inline Policy: Helpful if you want to make a one-to-one relationship between a policy and the principal entity
- Combination of differnt variables in an IAM policy will lead to granular Permissions
- Either a resource policy or Identity based policy is enough to grant acces to a resource
- ![image](https://user-images.githubusercontent.com/51809378/156701911-9006f93b-e770-4c23-92d9-b5c53708bfa4.png)
### AWS STS
- Token service in AWS which provides temprorary access to AWS with a life time of 1-36hrs
- Federation with AD
  - **gives temporary creds to AWS Console (STS API AssumeRoleWithSAML)**
  - Provides SSO for users
  - ADFS is a trusted ID provider
  - A trust relationship must be configured
  - ![image](https://user-images.githubusercontent.com/51809378/156702507-9380089f-d71c-4fa9-986d-ed2439853495.png)
  - Federation with facebook, google and many Open ID providers
### Web Identity Federation
- allows us to login with FB, Linked..
- AWS Congnito
- User pools - user directories; users can login directly to user pools or indirectly via ID providers
- Identity pools - create unique identities, can give temporary credentials to AWS resources
  - OAuth2.0
    - OAuth scope - options to verify identity, e.g. phone mail.
    - implicit grant - you'll get your JWT token
    - Authorization code grant - Cognito will give you authorization code back to process it further on the backend side
### S3
- S3 pre-signed URLs are typically handled via SDK the URLs expires after defined time
- If you use the policy generator at the end of bucket's ARN - without it you can still do actions on any object (only actions against bucket will be denied, e.g. list all objects)!!! On the other hand if you put /* then everyone else can list its content.
- Simple way to grant cross-account to S3 service without using IAM roles
### AWS Organisations
-  Allows us to setup service control policy [ploicies appled to OU's] to block access, these cannot be used to allow access to a certain service
-  They even allows us to specify a bounbdary for root account, overwrites any local policy within OU

