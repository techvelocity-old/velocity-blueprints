# Prometheus-Example

Run the following command from the `prometheus-chart` directory to deploy Prometheus and Grafana to a Velocity environment. 

```
helm template . -f velocity-values.yaml --name-template=velocity | veloctl env create -f -
```

