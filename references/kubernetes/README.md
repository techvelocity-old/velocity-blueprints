# kubernetes

[cloud-resources](cloud-resources) contains example snippets for when you want to use real cloud resources in your Velocity environments.

[common-containers](common-containers) contains example snippets for when you want to use a database/other common services as a container in your Velocity environments.

### When should you use cloud resources?
If you can't find a "fake" replacement, and you don't want to invest on mocking/faking it - use the real thing. That will provide you with the best "production-like" experience.

Example use cases:
* SQS/Pubsub
* Storage Bucket
* DynamoDB (Serverless)

### When should you use a container replacement?
In many cases that involves a database, provisioning a separate cloud resource for each developer environments can be costly (in both time, that is provisioning time, and money).
In that sense it can make more sense to use a container, which is more lightweight and quicker to bring up, than using an actual cloud resource.

Example use cases:
* MongoDB
* PostgreSQL
* MySQL
* Redis
* RabbitMQ