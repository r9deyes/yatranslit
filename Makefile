build-wasm:
  GOOS=js\
  GOARCH=wasm\
  go build -o demo/main.wasm demo/wasm_demo.go

build-wasm-tiny:
  tinygo build -o demo/tiny.wasm -target wasm ./demo/wasm_demo.go
