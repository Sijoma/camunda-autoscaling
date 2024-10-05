### Kind Dev environment

setup-kind:
	kind create cluster --name hackdays --config deploy/local/kind-cluster.yaml

deploy-stack:
	$(MAKE) deploy-metrics-server
## Needs to be run multiplce times due to CRDs
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

deploy-keda:
	kustomize build ./deploy/local/keda | kubectl apply --server-side -f -

deploy-demo:
	kubectl apply -f ./deploy/local/demo

undeploy-camunda:
	kustomize build --enable-helm ./deploy/local/camunda | kubectl delete -f -

undeploy-demo:
	kubectl delete -f ./deploy/local/demo

undeploy-keda:
	kustomize build --enable-helm ./deploy/local/keda | kubectl delete -f -

teardown:
	kind delete cluster --name hackdays

### Demo stuff

demo-app:
	cd ./demo/app && mvn package -Ddocker.goal=dockerBuild

demo-exporter:
	cd demo/exporter && \
	mvn package -DskipTests -DskipChecks -T1C && \
	docker build -t "ghcr.io/sijoma/camunda-autoscaling-hackday/job-metric-exporter:SNAPSHOT" .

deploy-demo-kind:
	kind load docker-image ghcr.io/sijoma/camunda-autoscaling-hackday/starter:SNAPSHOT --name hackdays
	kind load docker-image ghcr.io/sijoma/camunda-autoscaling-hackday/worker:SNAPSHOT --name hackdays
	kind load docker-image ghcr.io/sijoma/camunda-autoscaling-hackday/job-metric-exporter:SNAPSHOT --name hackdays
	kind load docker-image camunda/zeebe:8.6.0 --name hackdays
	helmfile -f demo/deployment/helmfile.yaml apply

deploy-operator:
	cd camunda-scaling-operator/deploy/local && \
	kustomize edit set image controller=ghcr.io/sijoma/camunda-scaling-operator:v0.0.3 && \
    kustomize build . | kubectl apply -f -