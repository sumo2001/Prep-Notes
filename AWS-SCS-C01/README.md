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
- You can authenticate to your DB instance using AWS Identity and Access Management (IAM) database authentication. IAM database authentication works with MySQL and PostgreSQL.
- With this authentication method, you don’t need to use a password when you connect to a DB instance. Instead, you use an authentication token.
- IAM database authentication provides the following benefits:
  - Network traffic to and from the database is encrypted using Secure Sockets Layer (SSL).
  - You can use IAM to centrally manage access to your database resources, instead of managing access individually on each DB instance.
  - For applications running on Amazon EC2, you can use profile credentials specific to your EC2 instance to access your database instead of a password
- Classic Load Balancer does not support Server Name Indication (SNI). You have to use an Application Load Balancer instead or a CloudFront web distribution to use SNI.
-  The public-facing Application Load Balancer can route the traffic to these web servers hosted in private subnets
-  You can turn off access to your instance metadata by disabling the HTTP endpoint of the instance metadata service, regardless of which version of the instance metadata service you are using. You can reverse this change at any time by enabling the HTTP endpoint. Use the modify-instance-metadata-options CLI command and set the http-endpoint parameter to disabled.
-  AWS Systems Manager agent simply makes it possible for Systems Manager to update, manage, and configure the AWS resources.
-  ![image](https://user-images.githubusercontent.com/51809378/156792179-079fbcc6-238c-4ca6-b353-2c8e4adb9bb4.png)
- Your VPC has attributes that determine whether instances launched in the VPC receive public DNS hostnames that correspond to their public IP addresses, and whether DNS resolution through the Amazon DNS server is supported for the VPC.
- Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC
- Perfect Forward Secrecy is a feature that provides additional safeguards against the eavesdropping of encrypted data, through the use of a unique random session key. This prevents the decoding of captured data, even if the secret long-term key is compromised.
- ![image](https://user-images.githubusercontent.com/51809378/156797742-eb082735-eebe-4560-8b83-22aa7965adcb.png)
- ![image](https://user-images.githubusercontent.com/51809378/156797882-45fa40d6-b24e-4e84-b887-6246aeec765c.png)
- CloudFront and Elastic Load Balancing are the two AWS services that support Perfect Forward Secrecy
- Amazon ECS enables you to inject sensitive data into your containers by storing your sensitive data in either AWS Secrets Manager secrets or AWS Systems Manager Parameter Store parameters and then referencing them in your container definition. Fargate also supports this feature
- Users of a master account can configure GuardDuty:
  - Users from a master account can generate sample findings in their own account. Users from a master account CANNOT generate sample findings in members’ accounts.
  - Users from a master account can archive findings in their own accounts and in all member accounts.
  - Users from a master account can upload and further manage trusted IP lists and threat lists in their own account.
- Users of a member account can configure GuardDuty:
  - Users from a member account can generate sample findings in their own member account. Users from a member account can’t generate sample findings in the master or other member accounts.
  - Users from a member account can’t archive findings either in their own account or in their master’s account, or in other member accounts.
  - Users from a member account can’t upload and further manage trusted IP lists and threat lists.
  - Trusted IP lists and threat lists that are uploaded by the master account are imposed on GuardDuty functionality in its member accounts. 
 - AWS Firewall Manager simplifies your AWS WAF, AWS Shield Advanced, and Amazon VPC security groups administration and maintenance tasks across multiple accounts and resources.
 - Firewall Manager provides these benefits:
   - Helps to protect resources across accounts
   - Helps to protect all resources of a particular type, such as all Amazon CloudFront distributions
   - Helps to protect all resources with specific tags
   - Automatically adds protection to resources that are added to your account
   - Allows you to subscribe all member accounts in an AWS Organizations organization to AWS Shield Advanced, and automatically subscribes new in-scope accounts that join the organization
   - Allows you to apply security group rules to all member accounts or specific subsets of accounts in an AWS Organizations organization, and automatically applies the rules to new in-scope accounts that join the organization
   - Lets you use your own rules, or purchase managed rules from AWS Marketplace
  - AWS Resource Access Manager (RAM) is a simple service that lets you share your resources with any AWS account or through AWS Organizations.
  - Each VPC endpoint is represented by one or more Elastic Network Interfaces (ENIs) with private IP addresses in your VPC subnets
  - AWS Systems Manager Parameter Store provides secure, hierarchical storage for configuration data management and secrets management. You can store data such as passwords, database strings, and license codes as parameter values.
### Identity and Access Management 
- A best practice they recommend is to delete root user access keys
- Inline policies are just regular policies that you create, manage, and embed directly into a single user, group, or role.
- You should not use the root account for managing any resource. Create an IAM user with the necessary permissions instead to perform these tasks.
- ![image](https://user-images.githubusercontent.com/51809378/156867467-16d81579-5dcd-44bd-9090-0e29c08675f5.png)
-  User pools are user directories that provide sign-up and sign-in options for your app users. Identity pools enable you to grant your users access to other AWS services. You can use identity pools and user pools separately or together.
-  ![image](https://user-images.githubusercontent.com/51809378/156869500-6927d30a-2e8c-482d-bba6-2f9b85a3daa7.png)
-  To use an external ID, update a role trust policy with the external ID of your choice. Then, when someone uses the AWS CLI or AWS API to assume that role, they must provide the external ID.
-  Amazon Cognito identity pools enable you to create unique identities and assign permissions for users. Your identity pool can include:
  - Users in an Amazon Cognito user pool
  - Users who authenticate with external identity providers such as Facebook, Google, or a SAML-based identity provider
  - Users authenticated via your own existing authentication process
- Trust relationships enable access to various resources that can be either one-way or two-way. A one-way trust is a unidirectional authentication path created between two domains
  - One-way:incoming – Users in the specified realm will not be able to access any resources in this domain.
  - One-way:outgoing – Users in this domain will not be able to access any resources in the specified realm.
  - Two-way (Bi-directional) – Users in this domain and users in the specified realm will be able to access resources in either domain or realm.
### Data Protection 
- Envelope encryption is the practice of encrypting plaintext data with a data key, and then encrypting the data key under another key
- ![image](https://user-images.githubusercontent.com/51809378/156876530-978ff8b3-6f23-486c-9d57-03a367dc0e39.png)
- A key policy document cannot exceed 32 KB (32,768 bytes). Key policy documents use the same JSON syntax
- "Principal": {"AWS": "arn:aws:iam::111122223333:root"} ->  Allows the AWS IAM service of the 111122223333 AWS Account to delegate permissions and KMS actions
- ![image](https://user-images.githubusercontent.com/51809378/156889143-aaf8f4b2-14d1-45c7-834a-805c3e097cdc.png)
- By default, the log files delivered by CloudTrail to your bucket are encrypted by Amazon server-side encryption with Amazon S3-managed encryption keys (SSE-S3).
- Enabling server-side encryption encrypts the log files but not the digest files with SSE-KMS. Digest files are encrypted with Amazon S3-managed encryption keys (SSE-S3).
- Automatic key rotation is not supported on the following types of CMKs, but you can rotate these CMKs manually.
  - 1. Asymmetric CMKs
  - 2. CMKs in custom key stores
  - 3. CMKs that have imported key material
- AWS KMS does not delete any rotated key material until you delete the CMK. Key rotation changes only the CMK’s backing key, which is the cryptographic material that is used in encryption operations
- AWS KMS supports optional automatic key rotation only for customer-managed CMKs.
- You must also have permissions to kms:Encrypt, kms:ReEncrypt*, kms:GenerateDataKey*, and kms:DescribeKey actions. to complete a KMS Action
- ![image](https://user-images.githubusercontent.com/51809378/156926272-64e311ac-3764-4e69-b368-8a752e9f384d.png)
- All AWS KMS cryptographic operations with symmetric CMKs accept an encryption context, an optional set of key-value pairs that can contain additional contextual information about the data. AWS KMS uses the encryption context as additional authenticated data (AAD) to support authenticated encryption.
- AWS Security Hub provides you with a comprehensive view of your security state in AWS and helps you check your environment against security industry standards including AWS Foundational Security Best Practices, CIS AWS Foundations Benchmark, and Payment Card Industry Data Security Standard (PCI DSS)
- 





