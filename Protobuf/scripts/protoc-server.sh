#!/bin/bash

ROOT_DIR=$(pwd)

SRC_DIR="${ROOT_DIR}/protobuf/syntax3"
OUT_DIR="${ROOT_DIR}/server"

# rm -r "${OUT_DIR}"
# mkdir "${OUT_DIR}"

protoc \
    -I=${SRC_DIR} \
    --go_out=${OUT_DIR} \
    --go_opt=default_api_level=API_OPAQUE \
    $(find "${SRC_DIR}" -iname "*.proto")
