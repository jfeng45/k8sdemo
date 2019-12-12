# Deploy: tell Tilt what YAML to deploy
k8s_yaml(['script/kubernetes/backend/backend-deployment.yaml','script/kubernetes/backend/backend-service.yaml'])

# Build: tell Tilt what images to build from which directories
docker_build("k8sdemo-backend:2.0", "./", dockerfile="script/kubernetes/backend/docker/Dockerfile-k8sdemo-backend")
#docker_build('companyname/frontend', 'frontend')
#docker_build('companyname/backend', 'backend')
# ...

# Watch: tell Tilt how to connect locally (optional)
#k8s_resource('frontend', port_forwards=8080)