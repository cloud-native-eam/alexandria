#!/bin/bash

#VERSION='v1.1.0'

glooctl install gateway

#add gloo repo to helm
#helm repo add gloo https://storage.googleapis.com/solo-public-helm

#helm install gloo.api gloo/gloo --namespace alexandria --version 1.1.0