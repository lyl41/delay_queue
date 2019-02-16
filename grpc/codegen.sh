#!/bin/bash

protoc --go_out=plugins=grpc:.  delayqueue.proto