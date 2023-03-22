# minikube k8sv1.23在mac可用，而且istio也支持1.23版本

minikube start \
    --cpus=3 \
    --memory=8900 \
    --kubernetes-version v1.23.14 \
    --image-mirror-country=cn \
    --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers