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
  - kubectl create deployment --image=nginx nginx
  - kubectl create deployment nginx --image=nginx --replicas=4
  - 

### Services
  - kubectl get service
  - kubectl describe svc kuberenetes
  - kubectl expose pod nginx --type=NodePort --port=80 --name=nginx-service --dry-run=client -o yaml
  - kubectl create service nodeport nginx --tcp=80:80 --node-port=30080 --dry-run=client -o yaml
  - kubectl create service clusterip redis --tcp=6379:6379 --dry-run=client -o yaml
  - Port and targetport must be mentioned, whereas nodeport is optional for a kubernetes cluster
### Namespaces
  - kubectl get pods --namespace=dev
  - kubectl config set-context $(kubectl config current-context) --namespace=dev
  - kubectl get pods --all-names
  - kubectl run redis --image=redis -n=sumanth

### Imperative Commands
- Create Objects
  - kubectl create -f nginx.yaml
  - kubectl run --image=nginx nginx
  - kubectl expose deployment nginx --port 80
  - kubectl run redis -l tier=db --image=redis:alpine
- Update Objects
  - kubectl edit deployment nginx
  - kubectl replace -f nginx.yaml
  - kubectl replace --force -f nginx.yaml
  - kubectl scale deployment nginx --replicas=5
  - kubectl set image deployment nginx nginx=nginx:1.18
### Declarative Commands
- Create or Update
  - kubectl apply -f /path/config/files or kubectl apply -f nginx.yaml
### Rolling and Rollout Updates
- Update
  - kubectl apply -f delpoyment.yaml
  - kubectl set image deployment/myapp-deployment nginx=nginx:1.19.1
 - Status
   - kubectl rollout status deployment/myapp-deployment
   - kubectl rollout history deployment/myapp-deployment
  - Rollback
    - kubectl rollout undo deployment/myapp-deployment   
 
 
 
