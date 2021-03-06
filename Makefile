BINARY_NAME=relayctl
BINARY_TARGET=$(BINARY_NAME).target
CONFIG=pins.toml

include Makefile.target

all: test build deploy run

test:
	ginkgo

cover:
	ginkgo -cover

build:
	GOARCH=$(GOARCH) GOARM=$(GOARM) GOOS=$(GOOS) go build -o $(BINARY_TARGET)

deploy:
	scp $(BINARY_TARGET) $(USER)@$(HOST):$(DEST)

run:
	-ssh -t $(USER)@$(HOST) $(DEST)/$(BINARY_TARGET) -config $(DEST)/$(CONFIG)

.PHONY: all test cover build deploy run
