cd /home/vagrant/jfeng45/k8sdemo/script/kubernetes/jenkins
kubectl apply -f service-account.yaml
kubectl create clusterrolebinding service-reader-pod --clusterrole=service-reader  --serviceaccount=default:default
kubectl apply -f jenkins-volume.yaml
kubectl apply -f jenkins-deployment.yaml
kubectl apply -f jenkins-service.yaml

# kubectl delete clusterrolebinding service-reader-pod