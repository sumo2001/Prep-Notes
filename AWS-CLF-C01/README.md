# AWS Cloud Practioner Important Points AWS CLF-C01

Hi there!
This is my personal notes which i wrote in hurry prep of AWS Cloud Practioner CLF-C01
+ **Security** 
     + AWS [**OF**]
     + Customer[**IN**]
     + Shared: Patch Management, Configuration Management, Awareness & Training


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
* Identity and Access Management (IAM)
* Route 53 (DNS service)
* CloudFront (Content Delivery Network)
* WAF (Web Application Firewall)
* AWS Organisations

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
- Step Function

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
     + AWs CloudTrail Logs
     + AWS S3 Glacier
     + AWS Storage Gateway
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
- Free Services
  + IAM
  + VPC
  + Consolidated Billing
  + Elastic Beanstalk
  + CloudFormation
  + Auto Scaling Groups
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
- ![image](https://user-images.githubusercontent.com/51809378/135027384-6baa67a0-1963-4961-8c90-26ebab95e554.png)



## The 5 pillars
- Operational Excellence 
   + Perform Operations as Code
   + Annotate documentation
   + Make small changes, which are reversible
   + Check Operation Procedures Frequently
   + Accept Failures and learn from failures
- Security
   + Implement a strong identity foundation
   + Apply security at all layers and enable traceability
    + Protect data in rest nd transist
   + Use Best Security Practices and prepare for security events
- Reliability
   + Test recovery procedures
   + Automatically recover from failure
   + Scale horizontally to increase the availability
   + Stop guessing capacity and manage change in automation
- Performance Efficiency
    + Democratize advanced technologies
    + Go global in minutes and use serverless
  + Experiment more often and try to know all AWS services
- Cost Optimization
  + Adopt a consumption mode
  + Measure overall efficiency
  + Stop spending money on data center operations
  + Analyze and attribute expenditure
  + Use managed and application-level services to reduce the cost of ownership4
       
 AWS Well Architect Tool - Free tool to review your architectures against the 5 pillars Well-Architected Framework and adopt architectural best practices
 
 ## Partners APN
 - APN Technology Partners: providing hardware, connectivity, and software
 - APN Consulting Partners: professional services firm to help build on AWS
 - APN Training Partners: find who can help you learn AWS
 - AWS Competency Program: technical proficiency and proven customer success in specialized solution areas
 - AWS Navigate Program: help Partners become better Partners

## Practice Exams Review
- Agility - Agility refers to new IT resources being only a click away
- ![image](https://user-images.githubusercontent.com/51809378/135115999-7b1eb1a5-f18d-41ce-8ff6-079523bc8bcc.png)
- **AWS Trusted Advisor - AWS Trusted Advisor is an online tool that provides you real-time guidance to help you provision your resources following AWS best practices on cost optimization, security, fault tolerance, service limits, and performance improvement**
- Inspector means EC2- Ec2 means Inspector
- Chef and Puppet - AWS OpsWork
- Aws Organisations - Can Manage all AWS accounts and can be used to share reserved Instances amongst all units
- Systems Manager provides a unified user interface so you can view operational data from multiple AWS services and allows you to automate operational tasks across your AWS resources [Grouping]
- Service Catalog - AWS Service Catalog allows organizations to create and manage catalogs of IT services that are approved for use on AWS
- Enhance Database Availability is the main reason deploying an RDS database in a Multi-AZ configuration
- ![image](https://user-images.githubusercontent.com/51809378/134817783-b7dce802-e816-42d1-ab29-b9e388bc90cd.png)
- AWS Cloud trail logs, AWS S3 Glacier and AWS Storage Gateway by default encryption is enabled on them
- Auto Scaling helps you ensure that you have the correct number of Amazon EC2 instances available to handle the load for your application.
- AWS Compute Optimizer helps you identify the optimal AWS resource configurations, such as **Amazon EC2 instance types, Amazon EBS volume configurations, and AWS Lambda function memory size**
- Step Function - AWS Step Function lets you coordinate multiple AWS services into serverless workflows. You can design and run workflows that stitch together services such as AWS Lambda, AWS Glue and Amazon SageMaker
-  **Cloud Trail -   service enables risk auditing by continuously monitoring and logging account activity, including user actions in the AWS Management Console and AWS SDKs**
- The AMI must be in the same region as that of the EC2 instance to be launched. If the AMI exists in a different region, you can copy that AMI to the region where you want to launch the EC2 instance 
- Developer plan will give you general architectural guidance as you build and test
- AWS Shield Advanced provides expanded DDoS attack protection for
    + Amazon CloudFront distributions
    + Amazon Route 53 hosted zones
    + AWS Global Accelerator accelerators
    + ALB and ELB
    + Amazon Elastic Compute Cloud (Amazon EC2) Elastic IP addresses
- The AWS account must be able to operate as a standalone account. Only then it can be removed from AWS organizations
- ![image](https://user-images.githubusercontent.com/51809378/134818982-70c28160-97ae-4e5f-b22a-5efe97964453.png)
- **CloudWatch: CloudWatch provides data and actionable insights to monitor applications, respond to system-wide performance changes, optimize resource utilization, and get a unified view of operational health**
- Security Group acts as a firewall at the instance level whereas Network Access Control List acts as a firewall at the subnet level
- AWS Storage Gateway service provides three different types of gateways – Tape Gateway, File Gateway, and Volume Gateway
-  Outbound data to the internet from all AWS regions is billed at region-specific, tiered data transfer rates. Inbound data transfer into all AWS regions from the internet is free.
-  A Security Group is stateful, that is, it automatically allows the return traffic
-  A NACL contains a numbered list of rules and evaluates these rules in the increasing order while deciding whether to allow the traffic
-  ![image](https://user-images.githubusercontent.com/51809378/134849542-7c86fcb0-23bb-4555-9cca-bc1bc8c684f7.png)
-  Amazon EBS Snapshots are stored incrementally, which means you are billed only for the changed blocks stored
-  You will pay a fee each time you read from or write data stored on the EFS - Infrequent Access storage class
-  AWS Budgets - Cost budget, Usage budget, Reservation budget
-  EC2 instances can access files on an EFS file system across many Availability Zones, Regions and VPCs
-  AWS Cloud recommends microservices as an architectural and organizational approach to software development where software is composed of small independent services that communicate over well-defined APIs
-  EC2 instance user data is the data that you specified in the form of a bootstrap script or configuration parameters while launching your instance.
![image](https://user-images.githubusercontent.com/51809378/134890859-17e3817a-f551-44df-93ed-48a46f0d141a.png)
- Agility - A startup is working on a new application that needs to go to market quickly. The application requirements may need to be adjusted in the near future
- CloudWatch monitors performance, whereas CloudTrail monitors actions in your AWS environment.
- ![image](https://user-images.githubusercontent.com/51809378/135115436-e895b210-9ea2-4fcf-a790-8eb31adba8b1.png)





