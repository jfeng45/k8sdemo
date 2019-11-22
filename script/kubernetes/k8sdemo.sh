# sudo minikube start
kubectl apply -f k8sdemo-config.yaml
kubectl apply -f k8sdemo-secret.yaml

# cd /home/vagrant/jfeng45/k8sdemo/script/kubernetes
# kubectl get service
# kubectl get deployment
# kubectl describe configMap
# kubectl describe secret
# kubectl get pv
# kubectl get pvc