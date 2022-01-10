## Practice Tests Review and Important points to remember for SAA-C02


+ Amazon Redshift is a fully-managed petabyte-scale cloud-based data warehouse product designed for large scale data set storage and analysis.
+ Amazon DMS + Event Bridge + Kinesis Data Stream - streams the existing data files as well and ongoing file updates from Amazon S3 to Amazon Kinesis Data Streams
+ S3 cannot directly write data into SNS, SNS cannot directly send messages to Kinesis Data Streams
+ ElastiCache is a popular choice for real-time use cases like Caching, Session Stores, Gaming, Geospatial Services, Real-Time Analytics, and Queuing. 
+ Elasticache is used as a caching layer in front of relational databases. It is not a good fit to store data in key-value pairs from the IoT sources
+ With CloudTrail, you can log, continuously monitor, and retain account activity related to actions across your AWS infrastructure
+  VPN CloudHub operates on a simple hub-and-spoke model that you can use with or without a VPC. This design is suitable if you have multiple branch offices and existing internet connections and would like to implement a convenient, potentially low-cost hub-and-spoke model for primary or backup connectivity between these remote offices
+  A VPC endpoint enables you to privately connect your VPC to supported AWS services and VPC endpoint services powered by AWS PrivateLink without requiring an gateways
+  Amazon FsX - used for workloads such as machine learning, high-performance computing (HPC), video processing, and financial modeling, hot data and cold data
+  Cluster placement group -  packs instances close together inside an Availability Zone. This strategy enables workloads to achieve the low-latency network performance necessary for tightly-coupled node-to-node communication
+  Partition placement group - spreads your instances across logical partitions such that groups of instances in one partition do not share the underlying hardware with groups of instances in different partitions, hadoop, kafka
+  Spread placement group - A Spread placement group is a group of instances that are each placed on distinct racks, with each rack having its own network and power source.
+  Global Accelerator is a good fit for non-HTTP use cases, such as gaming (UDP), IoT (MQTT), or Voice over IP.
