# AWS Cloud Practioner Important Points AWS CLF-C01

Hi there!
This is my personal notes which i wrote in hurry prep of AWS Cloud Practioner CLF-C01
**Security** 
- AWS [**OF**]
- Customer[**IN**]
- Shared: Patch Management, Configuration Management, Awareness & Training


## Cloud Types and their Computing Types:
There are 3 types of cloud : Private Cloud, Public Cloud and Hybrid Cloud
**Private Cloud:**  Cloud services used by a single organization, not exposed to the public. Ex: Rackspace

**Public Cloud:** Cloud resources owned and operated by a third-party cloud service provider delivered over the Internet. Ex: GCP, AWS

**Hybrid Cloud:** Keep some servers on premises and extend some capabilities to the Cloud Ex: OnPrem+Public

**Infrastructure as a Service:**
• Amazon EC2 (on AWS)
• GCP, Azure, Rackspace, Digital Ocean, Linode

**Platform as a Service:**
• Elastic Beanstalk (on AWS)
• Heroku, Google App Engine (GCP), Windows Azure (Microsoft)

**Software as a Service:**
• Many AWS services (ex: Rekognition for Machine Learning) Google Apps (Gmail), Dropbox, Zoom

**AWS Global Services:**
- Identity and Access Management (IAM)
- Route 53 (DNS service)
- CloudFront (Content Delivery Network)
- WAF (Web Application Firewall)

**Server Less Services**
- Dynamo DB
- S3 Bucket
- Redshift
- AWS Anthena	
- Amazon Quick Sight
- AWS Glue
- AWS Lambda
- AWS API Gateway	
- AWS Code Build
- AWS Fargate
- AWS SQS
- AWS Comprehend

**Hybrid Services**
- AWS Code Deploy
- AWS System Manager
- AWS Storage Gateway
- AWS OpsWork (Cheff and Puppet)
- AWS Outposts*
- AWS Cloud Watch Logs

**Memory Based**
- Code Artifact and Artifact are 2 Different Services on AWS
- NAT Gateways are managed by AWS to get internet access to private subnets
- NAT Instances are managed by OWN to get internet access to private subnets
- NACL will have both ALLOW and DENY but security group can only ALLOW a service
- VPC Peering, VPC endpoints, AWS Direct Connect uses AWS Private Network
- At VPC Endpoint Gateway : Dynamo DB and S3
- At VPC Endpoint Interface : Rest AWS Services  
- For having transitive peering between thousands of VPC and on-premises, hub-and-spoke (star) connection use AWS transit Gateway
- Encryption Automatically enabled:
_        CloudTrail Logs
        S3 Glacier
        Storage Gateway_
- AWS Artifact - Can be used to support internal audit or compliance, Portal that provides customers with on-demand access to AWS compliance documentation and AWS agreements
- AWS GuardDuty - Intelligent Threat discovery to Protect AWS Account using ML
- AWS Insceptor - Automated Security Assessments for EC2 instances
- AWS Config - auditing and recording compliance of your AWS resources
- AWS Security Hub - Central security tool to manage security across several AWS accounts and automate security checks 
- Amazon Transcribe - convert speech to text
- Amazon **P**olly - Convert **T**ext to **s**peech [_TSP_]
- Amazon Lex - Automatic Speech Recognition (ASR) to convert speech to text - Chatbots
- SageMaker: machine learning for every developer and data scientist 
- Forecast: build highly accurate forecasts
- Kendra: ML-powered search engine 
- Personalize: real-time personalized recommendations
**Billing**
