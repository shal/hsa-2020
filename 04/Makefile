.PHONY: default build clean
BLDDIR ?= bin
# VERSION     ?= $(if $(TAG),$(TAG),$(COMMIT_SHA))
# BUILD_DATE  := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
# LDFLAGS     := -ldflags "-X $(IMPORT_BASE)/pkg/version.Version=$(VERSION) -X $(IMPORT_BASE)/pkg/version.BuildDate=$(BUILD_DATE)"

.EXPORT_ALL_VARIABLES:
GO111MODULE  = on

default: clean build

build: clean
	@mkdir -p $(BLDDIR)
	@go build $(LDFLAGS) -o $(BLDDIR)/server ./cmd/server

clean:
	@rm -rf $(BLDDIR)