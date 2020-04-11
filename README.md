# Alexandria

A cloud native tool for managing the enterprise architecture.

## Dev Setup
1. [Run MicroK8s](/docs/devenv/README.md) or Docker Desktop with K8s 
2. [Run octant](https://github.com/vmware-tanzu/octant) for more cluster visibility
3. Spin up the DB
    * execute the ./db/deploy-db.sh and wait till all pods are up
    * then `kubectl apply -f ./db/development.yaml`
    * You can reach the DB GUI via https://<`VM-IP`>:<`service/alexandria-db-ea`>
    * `https://192.168.64.5:30696/_db/_system/_admin/aardvark/index.html#login`
    * use `ax-admin` at the `axdb`, the password can be found in the secret

