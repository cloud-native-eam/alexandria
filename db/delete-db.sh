#!/bin/bash

VERSION='0.4.2'

# Install Development DB
kubectl delete -f ./adb/development.yaml

# Install Production DB
# kubectl apply -f ./adb/production.yaml

sleep 30s
kubectl delete -f ./adb/arango-deployment.yaml
kubectl delete -f ./adb/arango-crd.yaml


