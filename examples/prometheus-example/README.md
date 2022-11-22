# Prometheus-Grafana-Example

Run the following command from the `prometheus-example` directory to deploy pre-configured instances of Prometheus and Grafana to a Velocity environment. 

```
{ helm template grafana-chart --values=grafana-chart/velocity-values.yaml --name-template=velocity && helm template prometheus-chart --values=prometheus-chart/velocity-values.yaml --name-template=velocity } | veloctl env create -f -
```

## Default Grafana Credentials:
**username:** admin
**password:** admin

