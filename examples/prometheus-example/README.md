# Prometheus-Example

Run the following command from the `prometheus-chart` directory to deploy Prometheus to a Velocity environment. 

```
helm template . -f velocity-values.yaml | veloctl env create -f -
```

