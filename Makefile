makefile_dir		:= $(abspath $(shell pwd))
HELLGO_API_USER := $(API_USER)
HELLGO_API_URL  := $(API_URL)
HELLGO_API_PORT  := $(API_PORT)



print-vars:
	@echo 'url is $(HELLGO_API_URL)'
	@echo 'port is $(HELLGO_API_PORT)'
	@echo 'user is $(HELLGO_API_USER)'

all:
	touch env.sh
	@echo HELLGO_API_URL=$(HELLGO_API_URL) >> env.sh

push:
	git add .
	git commit -am "$(m)"
	git push origin master