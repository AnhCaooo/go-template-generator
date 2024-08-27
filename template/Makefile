APPLICATION_NAME ?= <application_name>
TAG_VERSION ?= latest

docker: 
	go test ./... && docker build --tag ${APPLICATION_NAME}/${TAG_VERSION} .
