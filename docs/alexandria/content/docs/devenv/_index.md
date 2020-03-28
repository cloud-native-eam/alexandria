---
weight: 1
bookFlatSection: true
title: "Development Setup"
---
# Using MicroK8s
(for Mac user)

Get and install Multipass https://microk8s.io/docs/install-alternatives#macos

Spin up a local machine and exec into it
```sh
multipass launch --name microk8s --mem 4G --disk 40G
multipass shell microk8s
```
Next wie need to install microk8s and set an alias to make our life easier
```sh
sudo snap install microk8s --classic --channel=1.17/stable
vim ~/.bash_aliases
alias k='sudo microk8s.kubectl' #save & close it
. ~/.bash_aliases
k cluster-info
```
Now we have a K8s running, next we need to spin up some basic services
```sh
sudo microk8s.enable dns dashboard storage helm3 rbac
k get all -A
```
Next we need to be sure that we can do access the K8s in the VM
```sh
sudo iptables -P FORWARD ACCEPT
```

## Accessing Service on K8s
To access we need actually two things 1. the IP of the VW and 2. the API for the service
```sh
multipass list 
#take the IP from your host/client
#and the api from K8s
k cluster-info
#https://127.0.0.1:16443/api/v1/namespaces/kube-system/services/monitoring-grafana/proxy
```

In your Webbrowser you now need to replace the local IP with the IP of the VM, then you will get asked for username and password. This you get in the VM by running
```sh
sudo microk8s.config
#in the end of the config you find the admin and PW
```

Now you should be able to log into the Dashboard! 
This way you can reach all exposed workload.

## Configure you local kubeconfig
At the moment you run K8s within the VM and you can reach the API, but you can't use kubectl from a local machine.
Therefore you need to update your kubeconfig which you can mostly found in your root under .kube/config
In a first step you need from within the VM the Kubeconfig
```sh
sudo microk8s.kubectl config view --raw
```

Then you open the kubeconfig and add cluster, context and user. Don't forget to replace the local IP with the VM IP. In the End your kubeconfig should looks like this:
```yaml
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: SUPERSECRET_CA_AUTH
    server: https://127.0.0.1:16443 
  name: microk8s-cluster
contexts:
- context:
    cluster: docker-desktop
    user: docker-desktop
  name: docker-desktop
- context:
    cluster: microk8s-cluster
    user: admin
  name: microk8s
current-context: microk8s
kind: Config
preferences: {}
users:
- name: admin
  user:
    password: SUPERSECRETPASSWORT
    username: admin
```

Save the config and run on you host/client
```sh
kubectl get all -A
```
If you see all the resources within the K8s from the VM you are done!


## About Multipass
Start VM
```sh
multipass start microk8s
```
Stop VM
```sh
multipass stop microk8s
```
Exec/Shell into VM
```sh
multipass shell microk8s
```
Delete & Clean up VMs
```sh
multipass delete microk8s
multipass purge
```
