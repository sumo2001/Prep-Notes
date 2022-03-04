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
## Section based review and key points from John Bonso
### Incident Response: 
- Amazon GuardDuty monitors the security of your AWS environment by analyzing and processing VPC Flow Logs, AWS CloudTrail event logs, and DNS logs
- Amazon GuardDuty is a threat detection service that continuously monitors for malicious activity and unauthorized behavior to protect your AWS accounts
- You can customise your threat lists and add trusted IPs in Guard Duty
- ![image](https://user-images.githubusercontent.com/51809378/156748907-9fa9334d-de92-4acd-886e-8ad3c48e544d.png)
- GuardDuty can send notifications based on Amazon CloudWatch Events when any changes in the findings take place. These changes include newly generated findings or subsequent occurrences of existing findings.
- In order to receive notifications about GuardDuty findings based on CloudWatch Events, you must create a CloudWatch Events rule and a target for GuardDuty. This rule enables CloudWatch to send events for all findings that GuardDuty generates to the target that is specified in the rule
-  Amazon Macie is just a security service that uses machine learning to automatically discover, classify, and protect sensitive data in AWS, specifically in Amazon S3
-  Amazon Macie has the ability to detect global access permissions inadvertently being set on sensitive data, detect uploading of API keys inside source code, and verify sensitive customer data is being stored and accessed in a manner that meets their compliance standard
-  Enabling rotation in AWS Secrets Manager causes the secret to rotate immediately, uses a Lambda function Secrets Manager provides
-  If your RDS instances or on a private VPC using AWS SM,  Secrets Manager also configures the Lambda function to run within that VPC
-  Secrets Manager encrypts the protected text of a secret by using the AWS Key Management Service (AWS KMS).
-  ![image](https://user-images.githubusercontent.com/51809378/156751844-26df5a55-6996-4bd9-af88-6b46f760c928.png)
-  Amazon GuardDuty can generate findings based on suspicious activities such as requests coming from known malicious IP addresses, changing of bucket policies/ACLs to expose an S3 bucket publicly, or suspicious API call patterns that attempt to discover misconfigured bucket permissions.
-  Macie cannot detect usage patterns on S3 data. Although Amazon Macie is capable of detecting policy changes in S3 buckets, this is not enough to detect the potential threats
-  Amazon SQS has its own resource-based permissions system that uses policies written, Access to Amazon SQS requires credentials that AWS can use to authenticate your requests. 
-  Amazon Inspector is an automated security assessment service that helps improve the security and compliance of applications deployed on AWS. Amazon Inspector automatically assesses applications for exposure, vulnerabilities, and deviations from best practices
- You can assign up to five security groups to the instance. Security groups act at the instance level, not the subnet level.
-  NAT Gateway is primarily used to enable instances in a private subnet to connect to the Internet or other AWS services, but prevents the Internet from initiating a connection with those instances
-   You have to use AWS Config to check whether AWS CloudTrail is enabled on your AWS accounts.
-   ![image](https://user-images.githubusercontent.com/51809378/156755732-07629088-8c11-426e-adcc-e7fb4ed91670.png)
### Logging and Monitoring:
- AWS Security Hub simply gives you a comprehensive view of your high-priority security alerts and security posture across your AWS accounts. It is not capable of tracking the activities and resources of all AWS Regions, unlike AWS CloudTrail
- Use AWS services and third-party IDS/IPS solutions offered in AWS Marketplace to stay one step ahead of potential attackers. 
- When AWS Config detects changes that are non-compliant, you can create a CloudWatch Events rule that will trigger an AWS Lambda function that can perform actions
- Amazon Kinesis Data Streams enables real-time processing of streaming big data. It provides ordering of records, as well as the ability to read and/or replay records in the same order to multiple Amazon Kinesis Applications
- The Amazon Kinesis Client Library (KCL) -> to perform counting, aggregation, and filtering
- Elasticsearch is a popular open-source search and analytics engine from Elastic that provides a quick time to value and is well supported by a vibrant open-source community
- Use AWS Config to evaluate the configuration settings of your AWS resources. You do this by creating AWS Config rules, which represent your ideal configuration settings.
- You can configure your environment to stream logs to Amazon CloudWatch Logs in the Elastic Beanstalk console or by using configuration options
- If logs are only pushed for a short time after the awslogs agent is restarted, check for duplicates in the [logstream] section of the agent configuration file. Each section must have a unique name.
- If the awslogs.log log file takes up too much disk space, check the log file for errors and correct them. If the log file only contains informational messages, specify a lower logging level for the logging_config_file option in the agent configuration file.
- In s3, The resource owner can grant access permissions to other resources and users by writing an access policy
- . With Requester Pays buckets, the requester instead of the bucket owner pays the cost of the request and the data download from the bucket
### Infrastructure Security
- 
