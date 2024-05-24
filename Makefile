GOCMD=GO111MODULE=on GOPROXY=https://goproxy.io CGO_ENABLED=0 go
GOBUILD=$(GOCMD) build .

all:
	$(GOBUILD)