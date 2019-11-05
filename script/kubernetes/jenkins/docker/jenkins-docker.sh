cd /home/vagrant/jfeng45/k8sdemo/
# docker build -f ./script/kubernetes/backend/docker/Dockerfile-k8sdemo-backend -t k8sdemo-backend .
# docker run -it --name jenkins-docker jenkins-docker /bin/bash
# docker run -it --name jenkins-test jenkinsci/jenkins:2.154-slim /bin/bash
docker run --name jenkins-docker -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock jenkins-docker
docker build -f ./script/kubernetes/jenkins/docker/Dockerfile-jenkins-docker -t jenkins-docker .

# docker build -f ./script/kubernetes/jenkins/docker/Dockerfile-modified-jenkins -t jfeng45/modified-jenkins:1.0 .