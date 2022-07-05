# AWS Helm Sample

### What is being deployed?
Here we deploy AWS-Explorer and all of its dependencies, as shown in the following chart:

![](../../../references/aws-explorer/media/chart.png)

[AWS-Explorer](../../../references/aws-explorer) is a reference service that is used to act as an example application in AWS environments.

The k8s manifest code is available under the helm [templates](templates) directory and consists of:
1. [AWS-Explorer](templates/explorer.yaml) which is the example application. It depends on:
   1. [S3 bucket](templates/s3.yaml) (Runs as an actual cloud resource)
   2. [SQS queue](templates/sqs.yaml) (Runs as an actual cloud resource)
   3. [DynamoDB Table](templates/dynamodb.yaml) (Runs as an actual cloud resource)
   4. [PostgreSQL Database](templates/postgresql.yaml) (Runs as a container and not as a cloud managed service)
      1. The k8s manifest code for PostgreSQL was taken from the [reference snippet](../../../references/kubernetes/database-containers/postgresql.yaml).
   5. [DB Migration Seeding Job](templates/seeding.yaml)
      1. Because our app expects a certain DB schema that it can work with, we need to apply DB schema migration scripts.
      2. This k8s seeding Job runs DB migration scripts against the PostgreSQL DB.
      3. A k8s job will run once (considering it finished successfully) just after the PostgreSQL DB has started, and before the application has started.
      4. The migrations scripts it executes can be seen in [/references/db-migrations](../../../references/db-migrations)


### Helm

`velues-prod.yaml` is the file the contains the values that fit your production environment.

`values-velocity.yaml` is the file the contains the values that fit your velocity environments.


### HOWTO - Dry run for Production

To dry run the helm chart for prod environment, run the following command:
```shell
helm template --values values-prod.yaml . > prod.yaml
```

The resulting `prod.yaml` file reflects the resulting artifacts of the helm charts that are to be applied on your production environment.

### HOWTO - Dry run for Velocity Environment

To dry run the helm chart for velocity environments, follow these instructions:
* Modify `aws_access_key_id`, `aws_secret_access_key`, `aws_session_token` in `values-velocity.yaml` to reflect your AWS credentials for the AWS account you are working on.
* run the following command:
   ```shell
   helm template --set provision_resources=true --values values-velocity.yaml . > velocity.yaml
   ```
The resulting `velocity.yaml` file reflects the resulting artifacts of the helm charts that are to be applied on your velocity environments.


### HOWTO - Create a Velocity Environment
To create a Velocity Environment, first you need to run the instructions on the previous section (Dry run for Velocity Environment).

After doing so, you shall have a `velocity.yaml` result artifact.

To create an environment, pass the `velocity.yaml` file to a `veloctl env create` command as such:
```shell
veloctl env create -f velocity.yaml
```

NOTE: You can also choose to create the `velocity.yaml` file and running the `veloctl env create` command together by using pipes, as such
```shell
helm template --set provision_resources=true --values values-velocity.yaml . | veloctl env create -f -
```