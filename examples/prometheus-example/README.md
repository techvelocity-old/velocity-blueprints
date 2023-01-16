# Prometheus-Example

Run the following command from the `prometheus-chart` directory to deploy Prometheus and Grafana to a Velocity environment. 
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add grafana https://grafana.github.io/helm-charts
```

```
helm dependency build
```

```
helm template . -f velocity-values.yaml --name-template=velocity | veloctl env create -f -
```

