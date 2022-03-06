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

## Benchmarks

The raw performance results from 2019 are available as an excel spreadsheet as [performance results.xlsx](presentation/performance%20results.xlsx).

## Appendix

The performance benchmarks were presented at [Oslo Go User Group Go Oslo December Meetup @Vipps 🎅🏻 🧡](https://www.meetup.com/Go-Oslo-User-Group/events/266553658/) at Wed, Dec 4 · 6:00 PM 2019. The powerpoint presentation is available at [Interfacing with native code](presentation/Interfacing%20with%20native%20code.pptx).
