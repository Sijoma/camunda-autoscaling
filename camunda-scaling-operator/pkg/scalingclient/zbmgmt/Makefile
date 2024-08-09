OPENAPI_GENERATOR_IMAGE = openapitools/openapi-generator-cli:v7.1.0
OPENAPI_SPEC_FILE = cluster-api.yaml

.PHONY: all $(OPENAPI_SPEC_FILE) clean generate test

all:
	$(MAKE) $(OPENAPI_SPEC_FILE) clean generate

$(OPENAPI_SPEC_FILE):
	curl https://raw.githubusercontent.com/camunda/camunda/main/dist/src/main/resources/api/cluster/cluster-api.yaml  > cluster-api.yaml
	curl https://raw.githubusercontent.com/camunda/camunda/main/dist/src/main/resources/api/cluster/components.yaml  > components.yaml



clean:
	cat .openapi-generator/FILES | xargs rm -f

generate:
	docker run --rm \
		--user $(shell id -u) \
		-v ${PWD}:/local \
		--workdir /local \
		$(OPENAPI_GENERATOR_IMAGE) generate --config openapi-generator.yaml
	go fmt .
