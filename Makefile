### Kind Dev environment

setup-kind:
	kind create cluster --name hackdays --config deploy/local/kind-cluster.yaml

deploy-stack:
	$(MAKE) deploy-metrics-server
## Needs to be run multiplce times due to CRDs
	$(MAKE) deploy-monitoring
	sleep 5
	$(MAKE) deploy-monitoring
	sleep 5
	$(MAKE) deploy-monitoring
# Exchange with deploy without build
	$(MAKE) -C camunda-scaling-operator deploy-test-env
	$(MAKE) deploy-camunda

deploy-camunda:
	kustomize build --enable-helm ./deploy/local/camunda | kubectl apply -f -

## Needs to be run twice due to CRDs
deploy-monitoring:
	kustomize build --enable-helm ./deploy/local/monitoring | kubectl apply --server-side -f -

deploy-metrics-server:
	kustomize build ./deploy/local/metrics-server | kubectl apply --server-side -f -

undeploy-camunda:
	kustomize build --enable-helm ./deploy/local/camunda | kubectl delete -f -

teardown:
	kind delete cluster --name hackdays
