# Variables
IMAGENAME=pi-temp
VERSION=v0.1.0
PREFIX=sebastiangaiser
REGISTRY=docker.io

# Generic variables
IMAGEFULLNAME=${REGISTRY}/${PREFIX}/${IMAGENAME}:${VERSION}

.PHONY: help push login

help:
	    @echo "Makefile commands:"
	    @echo ""
	    @echo "push"
	    @echo "login"

.DEFAULT_GOAL := push

push:
	    @docker buildx build --platform linux/amd64,linux/arm64 -t ${IMAGEFULLNAME} --push .

login:
		@docker login -u ${PREFIX}