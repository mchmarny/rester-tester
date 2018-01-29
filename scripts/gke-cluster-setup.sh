#!/bin/bash

# Assumes env vars
# export PROJECT_ID=f9s-lab
# export CLUSTER_NAME=demo
# export DOCKER_REPO_OVERRIDE="us.gcr.io/${PROJECT_ID}"
# export K8S_CLUSTER_OVERRIDE="gke_${PROJECT_ID}_us-east1-d_${CLUSTER_NAME}"

echo "Setting project..."
gcloud config set project $PROJECT_ID

echo "Deleteing previous cluster..."
gcloud --quiet container clusters delete $CLUSTER_NAME

echo "Deleteing new cluster..."
 gcloud alpha --project=$PROJECT_ID container clusters create \
    --enable-kubernetes-alpha \
    --cluster-version=1.9.1-gke.0 \
    --zone=us-east1-d \
    --scopes=cloud-platform \
    --enable-autoscaling --min-nodes=1 --max-nodes=10 \
    $CLUSTER_NAME

echo "Getting credentials to new cluster..."
gcloud container clusters get-credentials $CLUSTER_NAME

echo "Setting self-admin to new cluster..."
kubectl create clusterrolebinding self-cluster-admin \
	--clusterrole=cluster-admin \
	--user=$(gcloud config get-value core/account)