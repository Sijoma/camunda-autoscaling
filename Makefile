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
	helmfile -f demo/deployment/helmfile.yaml destroy
