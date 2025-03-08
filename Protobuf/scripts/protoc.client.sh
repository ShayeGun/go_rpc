#!/bin/bash

ROOT_DIR=$(pwd)

PROTOC_GEN_TS_PATH="${ROOT_DIR}/client/node_modules/.bin/protoc-gen-ts_proto"
SRC_DIR="${ROOT_DIR}/protobuf/syntax3"
OUT_DIR="${ROOT_DIR}/client/protobuf"

rm -r "${OUT_DIR}"
mkdir "${OUT_DIR}"

protoc \
    -I=${SRC_DIR} \
    --plugin=${PROTOC_GEN_TS_PATH} \
    --ts_proto_out=${OUT_DIR} \
    $(find "${SRC_DIR}" -iname "*.proto")
