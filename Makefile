build-wasm:
  GOOS=js\
  GOARCH=wasm\
  go build -o docs/main.wasm docs/wasm_demo.go

build-wasm-tiny:
  tinygo build -o docs/tiny.wasm -target wasm ./docs/wasm_demo.go
