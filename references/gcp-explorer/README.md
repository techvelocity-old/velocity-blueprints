# gcp-explorer


![](./media/chart.png) 

GCP-Explorer is a reference service that is used to act as an example application in GCP environments.

GCP-Explorer depends on and uses the following services:
1. GCS Bucket
2. Pub/Sub Topic
3. Pub/Sub Subscription (Attached to the Topic)
4. MySQL Database (Runs as a container and not as a cloud managed service)

### TODO
* Add actual logic that relies on the DB schema migration result