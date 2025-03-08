# Installation
### Commands
1. download it's zip package from Github _(from release section)_

2. unzip it and move it to desired directory

3. unzipped file consists of a bin directory which includes an executable

4. add bin directory that PATH to .bashrc
    > export PATH="$PATH:{CUSTOM_PATH}/bin"

5. Golang plugin:
    - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest` => installs Go plugin for protoc
        > It must be in your $PATH for the protocol buffer compiler to find it.

6. TS plugin:
    - make a new directory
    - `npm init -y`
    - `npx tsc --init`
    - `npm install ts-proto`

### Links
- [protoc](https://protobuf.dev/installation/)
- [Go Plugin](https://protobuf.dev/reference/go/go-generated-opaque/)
- [TS Plugin](https://github.com/stephenh/ts-proto?tab=readme-ov-file)

<br>

# Initial Usage
### Commands
#### Go
- `protoc -I={SRC_DIR} --go_out={DIST_DIR} {SRC_DIR}/{FILE_NAME}.proto`

#### TS
- `protoc -I={SRC_DIR} --plugin={PROJECT_DIR}/node_modules/.bin/protoc-gen-ts_proto --ts_proto_out={DIST_DIR}  {SRC_DIR}/{FILE_NAME}.proto`

### Explanations
- `go_package` option inside `.proto` file is relative to `--go_out` path which is specified
- generated package name is entire path specified in `go_package`

### Links
- [Go Quick start](https://protobuf.dev/getting-started/gotutorial/)
- [TS Quick start](https://github.com/stephenh/ts-proto?tab=readme-ov-file#quickstart)

<br>

# Issues

### Importing multiple .proto files
- MUST be in same directory
- MUST be in same proto package
- MUST have `go_package` option _(for Golang)_
- `protoc -I={SRC_DIR} --go_out={DIST_DIR} {SRC_DIR}/{FILE_NAME_1}.proto {SRC_DIR}/{FILE_NAME_2}.proto`
    > Looks like these options are ignored by `TS plugin`

### Opaque Usage With Edition
1. edition = "2023";
2. import "google/protobuf/go_features.proto";
3. option features.(pb.go).api_level = API_OPAQUE;
- [Link](https://protobuf.dev/reference/go/go-generated-opaque/)

### Opaque Usage With Syntax3
just add this option to the end of protoc command `--go_opt=default_api_level=API_OPAQUE`
no need to change any thing in .proto file  unlike **edition** :

```bash
protoc -I={SRC_DIR} --go_out={DIST_DIR} {SRC_DIR}/{FILE_NAME}.proto --go_opt=default_api_level=API_OPAQUE
```

- [Link](https://protobuf.dev/reference/go/go-generated-opaque/)


<br>

# Notes
- There is **no correlation** between the `Go import path` and the `package specifier in the .proto file`. The latter is only relevant to the `protobuf namespace`, while the former is only relevant to the `Go namespace`. Also, there is no correlation between the Go import path and the .proto import path.

</br>

# Other Useful Links

- [Proto Buffer best Practices](https://protobuf.dev/best-practices/dos-donts/)

- [Protobuf edition vs syntax](https://protobuf.dev/editions/overview/)

- [Protobuf encoding](https://protobuf.dev/programming-guides/encoding/)

- [Golang Open struct vs Opaque](https://go.dev/blog/protobuf-opaque)

- [writing scripts and better protobuf structure](https://blog.logrocket.com/using-protobuf-typescript-data-serialization/)
