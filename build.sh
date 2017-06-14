#!/bin/bash

protoc api.proto \
  --go_out=plugins=grpc,import_path=generated:generated \
  -I. -I$HOME/protoc/include

