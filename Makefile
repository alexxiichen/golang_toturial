export GOPATH=$(shell dirname `pwd`):$(GOLANGPATH)

# $@--目标文件，$^--所有的依赖文件，$<--第一个依赖文件。

$(target): $(target)/*.go
	go build -o ../bin/golang_toturial/$(target) $@

DIRS := $(shell find . -maxdepth 1 -type d)

all : $(DIRS)

$(DIRS) : .
	go build -o ../bin/golang_toturial/$@ $@/*

.PHONY : $(target) clean

clean:
	rm -rf ../bin/golang_toturial/*