# Traefik
We are going to use traefik as Ingress/API Gateway. As there is no stable Helm release available we will do it by our own.

```
kubectl apply -f traefik/
kubectl get all -A
```
find the external IP and got to your browser, use the external IP + 8100

