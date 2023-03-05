# Dump data from ElasticSearch

```
npm install elasticdump -g
```

```
elasticdump --input=https://<user>:<password>@<es_host>:9200/index_name --output=index_data.json --type=data
```

# Seed ElasticSearch
```
veloctl snapshot put --target es-seeding-example -f index_data.json --default
```

```
helm repo add elastic https://helm.elastic.co

{ helm template elastic/elasticsearch --set replicas=1 --set tests.enabled=false && cat elasticsearch-seed-job.yaml } | veloctl env create -f -
```
