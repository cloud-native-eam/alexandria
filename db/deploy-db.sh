#!/bin/bash

VERSION='0.4.2'

kubectl apply -f ./adb/arango-crd.yaml
kubectl apply -f ./adb/warango-deployment.yaml

sleep 30
# Install Development DB
kubectl apply -f ./adb/development.yaml

# install the ingress route
# kubectl apply -f ./adb/ingress.yaml

# Install Production DB
# kubectl apply -f ./adb/production.yaml

#https://192.168.64.5:16443/api/v1/namespaces/default/services/alexandria-db/proxy/dashboard/