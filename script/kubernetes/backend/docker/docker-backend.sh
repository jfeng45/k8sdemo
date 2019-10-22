cd /home/vagrant/jfeng45/k8sdemo/
# docker build -f ./script/kubernetes/backend/docker/Dockerfile-k8sdemo-backend -t k8sdemo-backend .
# docker run -it --name k8sdemo-backend k8sdemo-backend /bin/sh
docker build -f ./script/kubernetes/backend/docker/Dockerfile-k8sdemo-backend-full -t k8sdemo-backend-full .