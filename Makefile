#!/usr/bin/make -f
SRC := $(wildcard cmd/gfm-serve/*.go */*/*.go )
BIN := bin/gfm-serve
TARGETS := darwin-arm64 darwin-amd64 linux-amd64

all: clean $(TARGETS)

$(TARGETS): $(SRC)
	GOOS=$(word 1,$(subst -, ,$@)) \
GOARCH=$(word 2,$(subst -, ,$@)) \
go build -o out/$@/$(BIN) $<

clean:
	$(foreach target,$(TARGETS),rm -rf out/$(target)/)

.PHONY: clean all
