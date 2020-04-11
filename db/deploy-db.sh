#!/bin/bash

VERSION='0.4.2'

kubectl apply -f ./adb/namespace.yaml
kubectl apply -f ./adb/arango-crd.yaml
kubectl apply -f ./adb/arango-deployment.yaml

sleep 30
# Install Development DB
kubectl apply -f ./adb/development.yaml

# Install Production DB
# kubectl apply -f ./adb/production.yaml