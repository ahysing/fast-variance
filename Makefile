
CGO_ENABLED=1
CLANG=clang

BUILD := build-portable
CLEAN := clean-portable
PREPARE_COMPILE_TARGETS := 

ifneq ($(OS),Windows_NT)
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		BUILD := build-native
		CLEAN := clean-native
	endif

	ifeq ($(UNAME_S),Darwin)
		BUILD := build-native-darwin
		CLEAN := clean-native
		LEAF7_FEATURES := $(shell sysctl -a | grep machdep.cpu.leaf7_features)
		ifeq ($(LEAF7_FEATURES),)
			PREPARE_COMPILE_TARGETS := cgo-sources-without-avx2
		endif
	endif
endif

.PHONY: build-native-darwin build-native build test build build-portable clean clean-portable

build: $(BUILD)
clean: $(CLEAN)

build-portable:
	$(info building portable binaries)
	$(info Operating System: $(OS))
	$(info C compiler: $(shell $(CC) -v))
	go build .

clean-portable:
	rm word_cloud
	@echo "done"



build-native:
	@echo "building Linux binary"
	$(info Operating System: $(shell uname -s))
	$(info C compiler: )$(shell cc -v)
	CGO_ENABLED=$(CGO_ENABLED) go build .

cgo-sources-without-avx2:
	$(info WARN downgrading c source files to not use AVX2 instruction set)
	sed -i .bak 's/ -mavx2 / /g' stats_nocgo_darwin.go
	sed -i .bak 's/ -mavx2 / /g' stats_darwin.go

build-native-darwin: $(PREPARE_COMPILE_TARGETS)
	@echo "building Darwin binary"
	$(info Operating System: $(shell uname -s))
	$(info C compiler: )$(shell cc -v)
	CGO_ENABLED=$(CGO_ENABLED) go build .

clean-native:
	rm word_cloud
	@echo "done"

test: $(BUILD)
	CGO_ENABLED=$(CGO_ENABLED) go test . -test.v
