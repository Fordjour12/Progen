# Simple Makefile for a simple program
#
# Author: Bobie Fordjour McCamble Kofi
# Date: 2024-06-23
#

all: build

build:
	@go build -o Progen main.go

.PHONY: all build
