#!/bin/bash

VERSION= 0.0.1

docker build -t alexandriaeam/axdb-init:$VERSION .
docker push alexandriaeam/axdb-init:$VERSION