tools.dir = tools
tools.bindir = tools/bin
tools.srcdir = tools/src

tools/controller-gen     = $(tools.bindir)/controller-gen
$(tools.bindir)/%: $(tools.srcdir)/%/pin.go $(tools.srcdir)/%/go.mod
	cd $(<D) && GOOS= GOARCH= go build -o $(abspath $@) $$(sed -En 's,^import _ "(.*)".*,\1,p' pin.go)

build: generate manifests
	mkdir -p bin
	CGO_ENABLED=0 go build -o bin/extension-server ./cmd/extension-server

image: build
	docker build -t extension-server:latest1 -f tools/docker/extension-server/Dockerfile .
	docker tag extension-server:latest1 tharsanan/extension-server:latest1
	docker push tharsanan/extension-server:latest1

manifests: $(tools/controller-gen)
	$(tools/controller-gen) crd:allowDangerousTypes=true paths="./..." output:crd:artifacts:config=charts/extension-server/crds/generated

generate: $(tools/controller-gen)
	$(tools/controller-gen) object:headerFile="$(tools.dir)/boilerplate.generatego.txt",year=2024 paths="{./api/...}"

tools.clean: # Remove all tools
	@$(LOG_TARGET)
	rm -rf $(tools.bindir)

clean: tools.clean
	rm -fr bin

settings: 
	-/home/tharsanan/Projects/gateway/examples/extension-server/main.sh || true
	-kubectl apply -f /home/tharsanan/Projects/gateway/examples/extension-server/settings.yaml || true
	-kubectl rollout restart deployment/envoy-gateway -n envoy-gateway-system || true

apis:
	kubectl apply -f /home/tharsanan/Projects/gateway/examples/extension-server/tapi.yaml

delete-apis:
	kubectl delete -f /home/tharsanan/Projects/gateway/examples/extension-server/tapi.yaml

qs:
	kubectl apply -f /home/tharsanan/Projects/gateway/examples/extension-server/qs.yaml

qs-api:
	kubectl apply -f /home/tharsanan/Projects/gateway/examples/extension-server/qs-api.yaml

install-envoy:
	helm install eg oci://docker.io/envoyproxy/gateway-helm --version v1.2.2 -n envoy-gateway-system --create-namespace --values /home/tharsanan/Projects/gateway/gateway-helm/values.yaml

uninstall-envoy:
	helm uninstall eg  -n envoy-gateway-system

delete-qs:
	kubectl delete -f /home/tharsanan/Projects/gateway/examples/extension-server/qs.yaml

test-qs:
	GATEWAY_HOST=$$(kubectl get gateway/eg -o jsonpath='{.status.addresses[0].value}') && echo host: $$GATEWAY_HOST  && \
	curl --verbose --header "Host: www.example.com" http://$${GATEWAY_HOST}/hello

test-auth:
	GATEWAY_HOST=$$(kubectl get gateway/eg -o jsonpath='{.status.addresses[0].value}') && echo host: $$GATEWAY_HOST  && \
	curl --verbose --header "Host: www.example.com" http://$${GATEWAY_HOST}/hello  --user 'user:p@ssw0rd'

test-api:
	GATEWAY_HOST=$$(kubectl get gateway/eg -o jsonpath='{.status.addresses[0].value}') && echo host: $$GATEWAY_HOST  && \
	curl --verbose --header "Host: www.example.com" http://$${GATEWAY_HOST}/example-context-1/hello  --user 'user:p@ssw0rd'


install-extension-server:
	helm install -n envoy-gateway-system extension-server /home/tharsanan/Projects/gateway/examples/extension-server/charts/extension-server

uninstall-extension-server:
	helm uninstall extension-server -n envoy-gateway-system

clean-all:
	-make delete-apis || true
	-make delete-qs || true
	-make uninstall-envoy || true
	-make uninstall-extension-server || true

install-all:
	make install-envoy
	make install-extension-server
	make settings
	make qs
	# make apis
	# make qs-api

.PHONY: build image manifests generate tools.clean clean
