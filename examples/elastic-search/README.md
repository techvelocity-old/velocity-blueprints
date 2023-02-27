```
helm repo add elastic https://helm.elastic.co
helm dependency build
helm template . --namespace="{velocity.v1.envName}" --name-template=velocity -f values.yaml > output.yaml
```

Manual steps:

1. add annotation to Kibana deployment in `output.yaml`:
```
velocity.tech.v1/dependsOn: pre-install-velocity-kibana
```

2. comment out `post-delete-velocity-kibana` job in `output.yaml`

```
veloctl env create -f output.yaml
```