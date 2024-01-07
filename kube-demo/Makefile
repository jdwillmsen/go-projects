IMAGE=jdwillmsen/hellogopher
TAG=latest
VERSION=5.0

mk-start:
	minikube start

k8-get-nodes:
	kubectl get nodes

mk-docker-env:
	minikube docker-env
mk-docker-daemon:
	eval $(minikube -p minikube docker-env)

docker-build:
	docker build -t "${IMAGE}:${TAG}" -t "${IMAGE}:${VERSION}" .

k8-create:
	kubectl create -f hellogopher.yml
k8-get-all:
	kubectl get all
k8-deploy-hg:
	kubectl get deployment/hellogopher -o yaml
k8-expose:
	kubectl expose deployment hellogopher --type=NodePort --port=8080

mk-service:
	minikube service hellogopher

k8-delete-deploy:
	kubectl delete deployment hellogopher
k8-delete-service:
	kubectl delete service hellogopher
k8-get-all-labels:
	kubectl get all --show-labels
k8-get-pods:
	kubectl get pods --show-labels
k8-label-pod:
	kubectl label po/hellogopher-567d594448-nbhvx author=jake
k8-label-pod-update:
	kubectl label po/hellogopher-567d594448-nbhvx author=ironman --overwrite
k8-label-pod-delete:
	kubectl label po/hellogopher-567d594448-nbhvx author-
k8-pod-selector:
	kubectl get pods --selector app=hellogopher
k8-pod-selector-2:
	kubectl get pods --selector app!=hellogopher
k8-delete-pod:
	kubectl delete pods -l app=hellogopher
k8-describe-pod:
	kubectl describe pod/hellogopher-764fcc47c-hg4dg
k8-update-image:
	kubectl set image deployment/hellogopher hellogopher=jdwillmsen/hellogopher:2.0
k8-apply:
	kubectl apply -f hellogopher.yml
k8-pod-logs:
	kubectl logs pod/hellogopher-5bcff6f69f-tp5sz
k8-exec:
	kubectl exec -it pod/hellogopher-5fb8947fdb-grbxc -- /bin/bash

mk-addons:
	minikube addons
mk-addons-list:
	minikube addons list
mk-enable-dashboard:
	minikube addons enable dashboard
mk-enable-metrics-server:
	minikube addons enable metrics-server
mk-dashboard:
	minikube dashboard

k8-apply-db:
	kubectl apply -f kubernetes/database.yml
k8-apply-all:
	kubectl apply -f kubernetes/
k8-exec-db:
	kubectl exec pod/mysql-7477676875-gmtf2 -it -- bash
	# mysql --user=root --password=$MYSQL_ROOT_PASSWORD
	# echo $MYSQL_SERVICE_HOST
k8-config:
	kubectl config set-context --current --namespace=hellogopher

mk-service-namespace:
	minikube service hellogopher -n hellogopher

k8-configmap-create:
	kubectl create configmap database-configs --from-env-file=dbConfig.properties
k8-configmap-delete:
	kubectl delete configmap database-configs
k8-configmap-get:
	kubectl get configmaps
k8-configmap-get-yaml:
	kubectl get configmaps -o yaml
k8-secrets-get:
	kubectl get secrets
k8-secrets-get-yaml:
	kubectl get secrets -o yaml
k8-patch:
	kubectl patch secret database-secrets \
	--type='json' \
	-p='[{"op":"replace","path":"/data/DATABASE_PASSWORD","value":"test"}]'
k8-top-pods:
	kubectl top pods