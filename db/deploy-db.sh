#!/bin/bash

VERSION='0.4.2'

kubectl apply -f https://raw.githubusercontent.com/arangodb/kube-arangodb/${VERSION}/manifests/arango-crd.yaml
kubectl apply -f https://raw.githubusercontent.com/arangodb/kube-arangodb/${VERSION}/manifests/arango-deployment.yaml

#sleep 30s
# Install Development DB
#kubectl apply -f development.yaml

# Install Production DB
# kubectl apply -f production.yaml

https://192.168.64.5:16443/api/v1/namespaces/default/services/alexandria-db/proxy/dashboard/