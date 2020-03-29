# Alexandria
[![Netlify Status](https://api.netlify.com/api/v1/badges/16c88884-3a57-4823-bcfe-89f527c62e2e/deploy-status)](https://app.netlify.com/sites/competent-elion-26fee3/deploys)

[Documentation](https://competent-elion-26fee3.netlify.com/)

A cloud native tool for managing the enterprise architecture.

## Dev Setup
1. [Run MicroK8s](/docs/devenv/README.md) or Docker Desktop with K8s 
2. [Run octant](https://github.com/vmware-tanzu/octant) for more cluster visibility
3. Spin up the DB
    * execute the ./db/deploy-db.sh and wait till all pods are up
    * then `kubectl apply -f ./db/development.yaml`
    * You can reach the DB GUI via https://<`VM-IP`>:<`service/alexandria-db-ea`>
    * use `root` as user with an empty password and change it soon
4. Spin Up API Gateway - Traefik
    *  'helm install ./api/traefik'

