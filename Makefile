### Kind Dev environment

setup-dev:
	kind create cluster --name hackdays --config deploy/local/kind-cluster.yaml
	$(MAKE) deploy-camunda

deploy-camunda:
	kustomize build --enable-helm ./deploy/local/camunda | kubectl apply -f -

undeploy-camunda:
	kustomize build --enable-helm ./deploy/local/camunda | kubectl delete -f -

# TODO: use camunda-scaling-operator/Makefile:148
# deploy-operator:


teardown:
	kind delete cluster --name hackdays

### Demo stuff

build-demo-starter:
	docker build -t "gcr.io/zeebe-io/starter:autoscaling-demo" -f demo/starter.Dockerfile ./demo/

build-demo-worker:
	docker build -t "gcr.io/zeebe-io/starter:autoscaling-demo" -f demo/starter.Dockerfile ./demo/