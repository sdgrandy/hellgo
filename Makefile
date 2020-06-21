makefile_dir		:= $(abspath $(shell pwd))
HELLGO_API_URL  := $(API_URL)
HELLGO_API_PORT  := $(API_PORT)
HELLGO_API_USER  := $(API_USER)
HELLGO_API_PASSWORD  := $(HELLGO_API_PASSWORD)
ENVIRONMENT := $(ENVIRONMENT)
export

docker_compose := docker-compose
docker_service := hellgo

print-vars:
	@echo 'url is $(HELLGO_API_URL)'
	@echo 'port is $(HELLGO_API_PORT)'
	@echo 'user is $(HELLGO_API_USER)'

vars:
	touch vars.env
	#@echo "HELLGO_API_URL=$(HELLGO_API_URL)" >> vars.env
	#@echo "HELLGO_API_PORT=$(HELLGO_API_PORT)" >> vars.env
	
push:
	git add .
	git commit -am "$(m)"
	git push origin $(ENVIRONMENT)

run:
	go build 
	./hellgo

docker-build:
	$(docker_compose) build $(docker_service)
	#docker build --no-cache --build-arg PORT=$(HELLGO_API_PORT) \
	#--tag build:2.0 .
	#@make vars
	#docker build --no-cache --tag build:2.0 .


docker-up:
	$(docker_compose) up $(docker_service)
	#docker run --name bb build:2.0
	#docker stop bb
	#docker rm --force bb