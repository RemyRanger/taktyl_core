#!/bin/bash

protoc src/api/rpc/event.proto --go_out=plugins=grpc:.