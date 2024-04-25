export

GIT_REF := $(shell git rev-parse --short=7 HEAD)

VERSION ?= commit-$(GIT_REF)
ifneq ($(CI),true)
	VERSION := $(USER)-$(VERSION)
endif

.PHONY: version
version:
	@echo $(VERSION)
