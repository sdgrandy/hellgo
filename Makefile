makefile_dir		:= $(abspath $(shell pwd))
#API_USERNAME.collect { it }
HELLGO_API_USER := $(API_USERNAME)
HELLGO_API_URL  := $(API_URL)
HELLGO_API_PORT  := $(API_PORT)
HELLGO_API_PASSWORD  := $(API_PASSWORD)
export

docker_compose := docker-compose
docker_service := hellgo

print-vars:
	@echo 'url is $(HELLGO_API_URL)'
	@echo 'port is $(HELLGO_API_PORT)'
	@echo 'user is $(HELLGO_API_USER)'

vars:
	#touch vars.env
	@echo "HELLGO_API_URL=$(HELLGO_API_URL)" >> vars.env
	@echo "HELLGO_API_PORT=$(HELLGO_API_PORT)" >> vars.env
	@echo "HELLGO_API_USER=$(API_USERNAME)" >> vars.env
	@echo "HELLGO_API_PASSWORD=$(HELLGO_API_PASSWORD)" >> vars.env

push:
	git add .
	git commit -am "$(m)"
	git push origin master

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