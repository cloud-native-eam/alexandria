#!/bin/bash

VERSION=0.0.2

docker build -t alexandriaeam/axdb-init:"$VERSION" .
docker push alexandriaeam/axdb-init:"$VERSION"