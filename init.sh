#!/bin/bash

echo "의존성 초기화"
wire ./internal
echo "DB 스키마 초기화"
go generate ./pkg/ent