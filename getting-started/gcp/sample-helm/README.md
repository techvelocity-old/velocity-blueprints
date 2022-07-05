# GCP Helm Sample

### What is being deployed?
Here we deploy GCP-Explorer and all of it's dependencies, as shown in the following chart:

![](../../../references/gcp-explorer/media/chart.png)

[GCP-Explorer](../../../references/gcp-explorer) is a reference service that is used to act as an example application in GCP environments.

The k8s manifest code is available under the helm [templates](templates) directory and consists of:
1. [GCP-Explorer](templates/explorer.yaml) which is the example application. It depends on:
   1. [GCS Bucket](templates/gcs.yaml) (Runs as an actual cloud resource)
   2. [Pub/Sub Topic](templates/pubsub.yaml) (Runs as an actual cloud resource)
   3. [Pub/Sub Subscription](templates/pubsub.yaml) (Runs as an actual cloud resource - Attached to the Topic)
   4. [MySQL Database](templates/mysql.yaml) (Runs as a container and not as a cloud managed service)
       1. The k8s manifest code for MySQL was taken from the [reference snippet](../../../references/kubernetes/database-containers/mysql.yaml).


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
* Change `project_id` in `values-velocity.yaml` to your GCP project ID.
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

### TODO

* Add seeding examples