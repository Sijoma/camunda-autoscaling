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

### Demo stuff

build-demo-starter:
	docker build -t "gcr.io/zeebe-io/starter:autoscaling-demo" -f demo/starter.Dockerfile ./demo/

build-demo-worker:
	docker build -t "gcr.io/zeebe-io/worker:autoscaling-demo" -f demo/worker.Dockerfile ./demo/

build-demo-exporter:
	cd demo/exporter && \
	mvn package -DskipTests -DskipChecks -T1C && \
	docker build -t "gcr.io/zeebe-io/job-metric-exporter:autoscaling-demo" .

push-demo-images:
	docker push gcr.io/zeebe-io/starter:autoscaling-demo && \
 	docker push gcr.io/zeebe-io/worker:autoscaling-demo && \
 	docker push gcr.io/zeebe-io/job-metric-exporter:autoscaling-demo

demo-images: build-demo-starter build-demo-worker build-demo-exporter push-demo-images

lint-demo:
	helmfile -f demo/deployment/helmfile.yaml lint

deploy-demo:
	helmfile -f demo/deployment/helmfile.yaml apply

clean-demo:
	helmfile -f demo/deployment/helmfile.yaml destroy && \
	kubectl get pvc -n demo --no-headers=true -o custom-columns=":metadata.name" | xargs -n1 kubectl -n demo delete pvc
