makefile_dir		:= $(abspath $(shell pwd))
HELLGO_API_USER := $(API_USER)
HELLGO_API_URL  := $(API_URL)
HELLGO_API_PORT  := $(API_PORT)



print-vars:
	@echo 'url is $(HELLGO_API_URL)'
	@echo 'port is $(HELLGO_API_PORT)'
	@echo 'user is $(HELLGO_API_USER)'