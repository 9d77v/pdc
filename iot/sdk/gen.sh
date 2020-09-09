#!/bin/sh
protoc -I=. -I=/usr/local/include/google/protobuf/*.proto --go_out=./pb ./*.proto

