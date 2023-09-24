package main

import (
	"syscall/js"

	"github.com/r9deyes/yatranslit/pkg/yatranslit"
)

var t = yatranslit.NewTranslit()

func main() {
	js.Global().Set("Transform", js.FuncOf(Transform))

	<-make(chan bool)
}

func Transform(this js.Value, args []js.Value) any {
	if len(args) > 0 {
		message := args[0].String()
		return t.Transform(message)
	}

	return ""
}
