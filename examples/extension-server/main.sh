kubectl create clusterrole listener-context-example-viewer \
           --verb=get,list,watch,update  \
           --resource=ListenerContextExample

kubectl create clusterrolebinding envoy-gateway-listener-context \
           --clusterrole=listener-context-example-viewer \
           --serviceaccount=envoy-gateway-system:envoy-gateway


kubectl create clusterrole api-viewer \
           --verb=get,list,watch,update  \
           --resource=API

kubectl create clusterrolebinding api-context \
           --clusterrole=api-viewer \
           --serviceaccount=envoy-gateway-system:envoy-gateway