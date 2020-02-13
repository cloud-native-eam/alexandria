# API - Traefik
We are going to use Traefik as the API Gateway.
Currently Traefik v2.X is not as stable release available so we use the experimental version of containous.

https://github.com/containous/traefik-helm-chart/tree/master/traefik

Deploy and access the web ui of traefik
```
kubectl apply -f traefik-ui-svc.yaml
kubectl apply -f alexandria-ig.yaml
echo "$(minikube ip) traefik-ui.minikube" | sudo tee -a /etc/hosts
```


