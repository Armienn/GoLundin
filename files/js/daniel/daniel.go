package main

import(
	//"fmt"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Get("outputText").Set("innerHTML", "yolo")
}
