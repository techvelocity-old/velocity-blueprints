```
helm repo add elastic https://helm.elastic.co
helm dependency build

helm template . --name-template=velocity -f velocity-values.yaml | veloctl env create -f -
```