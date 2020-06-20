makefile_dir		:= $(abspath $(shell pwd))
HELLGO_API_USER := $(API_USER)
#HELLGO_API_URL  := $(API_URL)
HELLGO_API_URL  := google.com
HELLGO_API_PORT  := $(API_PORT)
export

buildvars := "-X main.HELLGO_API_URL=yahoo.com -X main.HELLGO_API_PORT=7500"

print-vars:
	@echo 'url is $(HELLGO_API_URL)'
	@echo 'port is $(HELLGO_API_PORT)'
	@echo 'user is $(HELLGO_API_USER)'

vars:
	touch env.sh
	chmod +x env.sh
	@echo export HELLGO_API_URL=$(HELLGO_API_URL) >> env.sh
	@echo export HELLGO_API_PORT=$(HELLGO_API_PORT) >> env.sh
	@echo export HELLGO_API_USER=$(HELLGO_API_USER) >> env.sh

push:
	git add .
	git commit -am "$(m)"
	git push origin master

run:
	go build -ldflags $(buildvars)
	./hellgo

docker-build:
	docker build --no-cache --build-arg PORT=$(HELLGO_API_PORT) --tag build:2.0 .

docker-run:
	docker run --name bb build:2.0
	docker stop bb
	docker rm --force bb