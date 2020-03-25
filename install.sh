#!/usr/bin/env bash

make docker-build

kind load docker-image webhook-image --name webhhok

kubectl apply -f config/crd