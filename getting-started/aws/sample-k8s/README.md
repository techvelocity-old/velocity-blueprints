# sample-k8s

### What is being deployed?
Here we deploy AWS-Explorer and all of its dependencies, as shown in the following chart:

![](../../../references/aws-explorer/media/chart.png)

[AWS-Explorer](../../../references/aws-explorer) is a reference service that is used to act as an example application in AWS environments.

The k8s manifest code is available under this directory and consists of:
1. [AWS-Explorer](explorer.yaml) which is the example application. It depends on:
    1. [S3 bucket](s3.yaml) (Runs as an actual cloud resource)
    2. [SQS queue](sqs.yaml) (Runs as an actual cloud resource)
    3. [DynamoDB Table](dynamodb.yaml) (Runs as an actual cloud resource)
    4. [PostgreSQL Database](postgresql.yaml) (Runs as a container and not as a cloud managed service)
        1. The k8s manifest code for PostgreSQL was taken from the [reference snippet](../../../references/kubernetes/common-containers/postgresql.yaml).
    5. [DB Migration Seeding Job](seeding.yaml)
        1. Because our app expects a certain DB schema that it can work with, we need to apply DB schema migration scripts.
        2. This k8s seeding Job runs DB migration scripts against the PostgreSQL DB.
        3. A k8s job will run once (considering it finished successfully) just after the PostgreSQL DB has started, and before the application has started.
        4. The migrations scripts it executes can be seen in [/references/db-migrations](../../../references/db-migrations)
