# camunda-autoscaling

A proof of concept around autoscaling a Camunda orchestration cluster & Camunda applications. 

This repository contains the following:

- A kubernetes Operator to automate the scaling operations for Camunda inside folder `camunda-scaling-operator`
    - It currently supports scaling brokers up & down
- A custom exporter that:
    - tracks the count of available jobs per type as a gauge
    - providing an absolute count per job type
- a demo application (job worker) that reports the workers capacity with the metric: `zeebe_client_worker_job_capacity`
- Local kind setup for testing inside the `deploy` folder.


## Scaling Brokers

1. Setup kind cluster & deploy the helm chart
```bash
make setup-kind
make deploy-stack
```
2. Go to Operator directory && build and deploy it to the current kubernetes context
```
cd camunda-scaling-operator && make deploy-test-env
```
3. Apply the ZeebeAutoscaler object to kubernetes
```bash
kubectl apply -f <name of your yaml>
```
```yaml
apiVersion: camunda.sijoma.dev/v1alpha1
kind: ZeebeAutoscaler
metadata:
  name: camunda-platform-zeebe
spec:
  replicas: 3
  zeebeRef:
    # Name of the Zeebe StatefulSet to scale in the same namespace
    name: camunda-platform-zeebe # Default (can be omitted)
    gateway:
      port: 9600
      # Name of the Service of the Zeebe Gateway
      serviceName: camunda-platform-zeebe-gateway # Default (can be omitted)
```
4. Scale Camunda Up
```bash
kubectl scale zeebeautoscalers.camunda.sijoma.dev camunda-platform-zeebe --namespace camunda-platform  --replicas 4
```

5. Scale Camunda Down
```bash
kubectl scale zeebeautoscalers.camunda.sijoma.dev camunda-platform-zeebe --namespace camunda-platform  --replicas 3
```

### Autoscaling Brokers with HPA- based on Prometheus metrics - WIP

This requires a [Prometheus](https://prometheus.io) & [Prometheus Adapter](https://github.com/kubernetes-sigs/prometheus-adapter) setup to expose the metrics to Kubernetes. 

If you want to learn more about this the following article may help to understand the basics: https://learnk8s.io/autoscaling-apps-kubernetes

```yaml 
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: camunda-hpa
  namespace: camunda-platform
spec:
  behavior:
    scaleDown:
      stabilizationWindowSeconds: 0
  minReplicas: 3
  maxReplicas: 6
  metrics:
    - type: Pods
      pods:
        # TODO: Needs proper metric
        metric:
          name: zeebe_backpressure_inflight_requests_count
        # TODO: Needs proper target
        target:
          type: AverageValue
          averageValue: "50"
  scaleTargetRef:
    apiVersion: camunda.sijoma.dev/v1alpha1
    kind: ZeebeAutoscaler
    name: camunda-platform-zeebe
```


### Autoscaling Workers from Zero

More details can be found [here](demo/README.md)
