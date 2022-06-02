- Pods
  - kubectl run nginx --image nginx
  - kubectl get pods
  - kubectl describe pods
  - kubectl get pods -o wide [short and cute description]
  - kubectl delete pod webapp
  - kubectl run redis --image=redis123 --dry-run=client -o yaml > redis.yaml
  - kubectl apply -f redis.yaml
  - kubectl edit pod podname
- ReplicaSet

- Deployments

- Services

- Imperative Commands

- Declarative Commands

- **Create an NGINX Pod**
  - kubectl run nginx --image=nginx

- Generate POD Manifest YAML file (-o yaml). Don't create it(--dry-run)
  - kubectl run nginx --image=nginx --dry-run=client -o yaml

- Create a deployment
  - kubectl create deployment --image=nginx nginx
- Generate Deployment YAML file (-o yaml). Don't create it(--dry-run)
  - kubectl create deployment --image=nginx nginx --dry-run=client -o yaml
- Generate Deployment YAML file (-o yaml). Don't create it(--dry-run) with 4 Replicas (--replicas=4)
  - kubectl create deployment --image=nginx nginx --dry-run=client -o yaml > nginx-deployment.yaml
- Save it to a file, make necessary changes to the file (for example, adding more replicas) and then create the deployment.
  - kubectl create -f nginx-deployment.yaml

**In k8s version 1.19+, we can specify the --replicas option to create a deployment with 4 replicas.** 
- kubectl create deployment --image=nginx nginx --replicas=4 --dry-run=client -o yaml > nginx-deployment.yaml
