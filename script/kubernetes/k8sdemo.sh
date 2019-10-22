# sudo minikube start
kubectl apply -f k8sdemo-config.yaml
kubectl apply -f k8sdemo-secret.yaml

# kubectl get service
# kubectl get deployment
# kubectl describe configMap
# kubectl describe secret
# kubectl get pv
# kubectl get pvc