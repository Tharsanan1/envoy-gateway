git add /home/tharsanan/Projects/gateway/internal/provider/kubernetes/routes.go
git commit -m "test"
commit_hash=$(git log -1 --format="%h")
echo $commit_hash

IMAGE=docker.io/tharsanan/gateway-dev make image
docker push docker.io/tharsanan/gateway-dev:$commit_hash
kubectl set image deployment/envoy-gateway -n envoy-gateway-system envoy-gateway=docker.io/tharsanan/gateway-dev:$commit_hash