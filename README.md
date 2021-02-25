# dapr-examples

## Overall Architecture

![dapr architecture](dapr.png)

### Install Dapr To K8S Cluster

![dapr install](dapr-init.jpeg)


### Apply Dapr Rabbitmq Component

```
kubectl apply -f dapr-rabbitmq.yaml
```

### Deploy Application (kind cluster)

```
docker build -t dapr-test .
kind load docker-image dapr-test
kubectl apply -f app.yaml
````

### View Rabbitmq & Dapr Dashboards

![dapr dashboard](dapr-dashboard.jpeg)
![rabbitmq dashboard](rabbitmq-dashboard.png)
