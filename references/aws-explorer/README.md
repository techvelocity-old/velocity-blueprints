# aws-explorer


![](./media/chart.png) 

AWS-Explorer is a reference service that is used to act as an example application in AWS environments.

AWS-Explorer depends on and uses the following services:
1. S3 bucket
2. SQS queue
3. DynamoDB Table
4. PostgreSQL Database (Runs as a container and not as a cloud managed service)

The AWS-Explorer service prints details about its dependencies upon booting.

The AWS-Explorer service also exposes an HTTP server that prints the contents of the S3 bucket when called in endpoint `/`

### TODO
* Add actual logic that relies on the DB schema migration result