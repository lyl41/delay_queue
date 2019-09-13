#!/usr/bin/env bash

all: dev run

fmt:
	goimports -l -w  ./app

install: fmt clean

clean:
	rm -rf output/conf/

dev: install
	go build -o output/bin/bin_delay_queue ./app

