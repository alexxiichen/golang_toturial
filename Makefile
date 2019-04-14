# export GOPATH=$(shell dirname `pwd`):$(GOLANGPATH)

# --目标文件，$^--所有的依赖文件，$<--第一个依赖文件。

all : helloworld basic_data_type printf

helloworld: day01/hellworld.go
	go build -o ../../bin/golang_toturial/day01/$@ $^

basic_data_type: day02/basic_data_type.go
	go build -o ../../bin/golang_toturial/day02/$@ $^

printf: day03/printf.go
	go build -o ../../bin/golang_toturial/day03/$@ $^

.PHONY : $(target) clean

clean:
	rm -rf ../../bin/golang_toturial/*