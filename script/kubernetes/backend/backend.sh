cd /home/vagrant/jfeng45/k8sdemo/script/kubernetes/backend
kubectl apply -f backend-volume.yaml
kubectl apply -f backend-deployment.yaml
kubectl apply -f backend-service.yaml

# kubectl get pv
# kubectl get pvc
# kubectl get service
# kubectl get deployment
# kubectl describe configMap
# kubectl describe secret
# kubectl exec -ti k8sdemo-backend-deployment-6b99dc6b8c-tbl4v -- /bin/sh
# kubectl exec -ti k8sdemo-backend-deployment-df5b977b5-skkc6 -- /bin/bash
