### Kind Dev environment

setup-dev:
	kind create cluster --name hackdays --config deploy/local/kind-cluster.yaml
	$(MAKE) deploy-camunda

deploy-camunda:
	kustomize build --enable-helm ./deploy/local/camunda | kubectl apply -f -

teardown:
	kind delete cluster --name hackdays
