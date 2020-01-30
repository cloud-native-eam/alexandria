# Alexandria

A cloud native tool for managing the enterprise architecture.

## Dev Setup
1. [Run MicroK8s](/docs/devenv/README.md) or Docker Desktop with K8s 
2. [Run octant](https://github.com/vmware-tanzu/octant) for more cluster visibility
3. execute the /db/deploy-db.sh and wait till all pods are up, then `kubectl apply -f development.yaml`
4. [Install Gloo CLI](https://docs.solo.io/gloo/latest/installation/gateway/kubernetes/) and get the gateway onto your cluster `glooctl install gateway`

