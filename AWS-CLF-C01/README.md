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
- AWS Organisations

## Server Less Services
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

## Hybrid Services
- AWS Code Deploy
- AWS System Manager
- AWS Storage Gateway
- AWS OpsWork (Cheff and Puppet)
- AWS Outposts*
- AWS Cloud Watch Logs

## Memory Based
- Code Artifact and Artifact are 2 Different Services on AWS
- NAT Gateways are managed by AWS to get internet access to private subnets
- NAT Instances are managed by OWN to get internet access to private subnets
- NACL will have both ALLOW and DENY but security group can only ALLOW a service
- VPC Peering, VPC endpoints, AWS Direct Connect uses AWS Private Network
- At VPC Endpoint Gateway : Dynamo DB and S3
- At VPC Endpoint Interface : Rest AWS Services  
- For having transitive peering between thousands of VPC and on-premises, hub-and-spoke (star) connection use AWS transit Gateway
- Encryption Automatically enabled:
		CloudTrail Logs
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

## Billing
- AWS Control Tower runs on top of AWS Organizations[SCP]
- Reservations are available for EC2 Reserved Instances, DynamoDB Reserved Capacity, ElastiCache Reserved Nodes, RDS Reserved Instance, Redshift Reserved Nodes
* Free Services
        * IAM
        * VPC
        * Consolidated Billing
        * Elastic Beanstalk
        * CloudFormation
        * Auto Scaling Groups
- AWS Compute Optimizer : • Reduce costs and improve performance by recommending optimal AWS resources for your workloads
- TCO Calculator: from on-premises to AWS
- Simple Monthly Calculator / Pricing Calculator: cost of services on AWS
- Billing Dashboard: high level overview + free tier dashboard
- Cost Allocation Tags: tag resources to create detailed reports
- Cost and Usage Reports: most comprehensive billing dataset
- Cost Explorer:View current usage (detailed) and forecast usage
- Billing Alarms: in us-east-1 – track overall and per-service billing
- Budgets: more advanced – track usage, costs, RI, and get alerts
- AWS STS[Security Token Service] - Kerberos

## The 5 pillars
- Operational Excellence
        - Perform Operations as Code
        - Annotate doccumentation
        - Make small changes, which are reversible
        - ChecK Operation Procedures Frequently
        - Accept Failures and learn from failures
- Security
        - Implement a strong identity foundation
        - Apply security at all layers and enable tracebility
        - Protect data in rest nd transist
        - Use Best Security Practices and prepare for security events
- Reliablity
        - Test recovery procedures
        - Automatically recover from failure
        - Scale horizontally to increase availability
        - Stop guessing capacity and manage change in automation
- Performance Efficieny
        - Democratize advanced technologies
        - Go global in minutes and use serverless
        - Experiment more often and try to know all AWS services
- Cost Optimization
        - Adopt a consumption mode
        - Measure overall efficiency
        - Stop spending money on data center operations
        - Analyze and attribute expenditure
        - Use managed and application level services to reduce cost of ownership4
       
 AWS Well Arcitect Tool - Free tool to review your architectures against the 5 pillars Well-Architected Framework and adopt architectural best practices
 
 ## Partners APN
 - APN Technology Partners: providing hardware, connectivity, and software
 - APN Consulting Partners: professional services firm to help build on AWS
 - APN Training Partners: find who can help you learn AWS
 - AWS Competency Program: technical proficiency and proven customer success in specialized solution areas
 - AWS Navigate Program: help Partners become better Partners
