### Pods
  - kubectl run nginx --image nginx
  - kubectl get pods
  - kubectl describe pods
  - kubectl get pods -o wide [short and cute description]
  - kubectl delete pod webapp
  - kubectl run redis --image=redis123 --dry-run=client -o yaml > redis.yaml
  - kubectl apply -f redis.yaml
  - kubectl edit pod podname
### ReplicaSet and Deployments
  - kubectl explain replicaset
  - kubectl create -f rs.yaml
  - kubectl get ReplicaSet
  - kubectl get replicaset new-replica-set
  - kubectl delete replicaset myapp-replicaset
  - kubectl delete pod new-replica-a3sdnn new-replica-rnf4f 
  - kubectl delete replicaset myapp-rs  [deletes all the underlying pods] 
  - kubectl replace -f replica-def.yaml
  - kubectl scale -replicas=6 -f replicaset.yaml
  - kubectl scale rs new-replica-set --replicas=5
  - kubectl edit rs new-replica-se
  - kubectl get all
### Services
  - kubectl get service
  - kubectl describe svc kuberenetes
  - kubectl expose ppod nginx --type=NodePort --port=80 --name=nginx-service --dry-run=client -o yaml
  - kubectl create service nodeport nginx --tcp=80:80 --node-port=30080 --dry-run=client -o yaml
  - kubectl create service clusterip redis --tcp=6379:6379 --dry-run=client -o yaml
  - Port and targetport must be mentioned, whereas nodeport is optional for a kubernetes cluster
### Namespaces
  - kubectl get pods --namespace=dev
  - kubectl config set-context $(kubectl config current-context) --namespace=dev
  - kubectl get pods --all-names
  - kubectl run redis --image=redis -n=sumanth

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
