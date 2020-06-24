# TODO :: Alpine variant
# TODO :: CockroachDB instance and link it
# Example: RUN ["apt-get", "install", "python3"]
FROM golang:1.11
EXPOSE 8443
WORKDIR /go/src/github.com/espal-digital-development/espal-core
COPY . .
# CMD ["/bin/sh", "./run.sh", "-dev"]
CMD /bin/sh /go/src/github.com/espal-digital-development/espal-core/run.sh -dev

# CREATE
# docker build -t espal-core:v1 .
# kubectl run espal-core --image=espal-core:v1 --port=8443 --image-pull-policy=Never
    # kubectl get services
    # kubectl get pods
    # kubectl get deployments
    # kubectl get events
# minikube service espal-core

# UPDATE
# docker build -t espa-node:v2 .
# kubectl set image deployment/espal-core espal-core=espal-core:v2
# minikube service espal-core

# kubectl delete service espal-core
# kubectl delete deployment espal-core
# OPTIONAL:
# docker rmi espal-core:v1 espal-core:v2 -f
# EXTRA OPTIONAL:
# minikube stop
# eval $(minikube docker-env -u)
# minikube delete