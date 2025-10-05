#!/bin/sh

tinygo \
	build \
	-o ./parquets2count.wasm \
	-target=wasip1 \
	-opt=z \
	-no-debug \
	./main.go
