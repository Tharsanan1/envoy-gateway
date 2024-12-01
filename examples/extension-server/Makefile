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
	kubectl apply -f /home/tharsanan/Projects/gateway/examples/extension-server/settings.yaml
	kubectl rollout restart deployment/envoy-gateway -n envoy-gateway-system

apis:
	kubectl apply -f /home/tharsanan/Projects/gateway/examples/extension-server/tapi.yaml

delete-apis:
	kubectl delete -f /home/tharsanan/Projects/gateway/examples/extension-server/tapi.yaml

qs:
	kubectl apply -f /home/tharsanan/Projects/gateway/examples/extension-server/qs.yaml

.PHONY: build image manifests generate tools.clean clean
