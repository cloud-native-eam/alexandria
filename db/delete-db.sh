#!/bin/bash

VERSION='0.4.2'

kubectl delete -f ./adb/arango-crd.yaml
kubectl delete -f ./adb/arango-deployment.yaml

#sleep 30s
# Install Development DB
#kubectl apply -f ./adb/development.yaml

# Install Production DB
# kubectl apply -f ./adb/production.yaml