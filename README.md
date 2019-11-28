# Fast variance in Go

This project implements population variance in go.
The project demonstrates performance speed ups with native c-code built for different systems.

## Prerequisites

* [clang](https://clang.llvm.org/) or [gcc](https://gcc.gnu.org/)
* [go 1.13.1](https://golang.org/)
* [make](https://www.gnu.org/software/make/)

## Usage

```shell
go test . -test.v
````

This software can be run without c-code and c-compiler. Just do

```shell
CGO_ENABLED=0 go test . -test.v
```
