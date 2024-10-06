

# Get Prom Adapter metrics:

Prom Adapter is needed when scaling with k8s HPA.

> kubectl get --raw /apis/custom.metrics.k8s.io/v1beta1

# Get Keda Metrics

Keda can be used in more powerful scenarios & does not require prom-adapter as it can query prometheus diretly.

https://keda.sh/docs/2.15/operate/metrics-server/

> kubectl get --raw "/apis/external.metrics.k8s.io/v1beta1"
