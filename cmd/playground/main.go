package main

import (
	"github.com/fengdotdev/golibs-frame/sandbox/draft1/goframe"
)

func main() {

	frame := goframe.NewWithBgColor(100, 100, goframe.RGBADot8{R: 255, G: 0, B: 0, A: 255}) // Create a new frame with a red background
	path := "test.png"
	err := frame.DrawImgToPNG(path)
	if err != nil {
		panic(err)
	}
	println("Image saved to", path)
}
