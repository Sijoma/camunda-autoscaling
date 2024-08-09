# camunda-autoscaling-hackday

Hack it til you make it! ⌨️

This repository contains the following:

- A kubernetes Operator to automate the scaling operations for Camunda inside folder `camunda-scaling-operator`
    - It currently supports scaling brokers up & down
- Local kind setup for testing inside the `deploy` folder.


## How to get started

1. Setup kind cluster & deploy the helm chart
```bash
make setup-dev
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
  namespace: camunda-platform # <-- namespace where the statefulset is in
spec:
  replicas: 3
  zeebeRef:
    name: camunda-platform-zeebe # <-- name of the statefulset
```
4. Scale Camunda Up
```bash
kubectl scale zeebeautoscalers.camunda.sijoma.dev camunda-platform-zeebe --namespace camunda-platform  --replicas 4
```

5. Scale Camunda Down
```bash
kubectl scale zeebeautoscalers.camunda.sijoma.dev camunda-platform-zeebe --namespace camunda-platform  --replicas 3
```

### Autoscaling - based on Prometheus metrics

This requires a proper [Prometheus](https://prometheus.io) & [Prometheus Adapter](https://github.com/kubernetes-sigs/prometheus-adapter) setup. 

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
        # TODO: find a proper metric
        metric:
          name: zeebe_backpressure_inflight_requests_count
        # TODO: find a proper target
        target:
          type: AverageValue
          averageValue: "50"
  scaleTargetRef:
    apiVersion: camunda.sijoma.dev/v1alpha1
    kind: ZeebeAutoscaler
    name: camunda-platform-zeebe

```
