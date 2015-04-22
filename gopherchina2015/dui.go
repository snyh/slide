package main

import "dui"
import "runtime"
import "fmt"

func main() {
	dui.Init()

	f := dui.NewFrame(350, 200)

	txt := f.NewElement("div")
	f.Add(txt)

	quit := f.NewElement("input")
	quit.Set("type", "button")
	quit.Set("value", "exit")
	quit.Set("id", "quit")
	f.Add(quit)
	f.Ele("quit").Connect("click", dui.Quit)

	hello := f.NewElement("input")
	hello.Set("type", "button")
	hello.Set("value", "hello")
	hello.Set("id", "hello")
	f.Add(hello)

	txt.Set("style", "text-shadow: 2px 2px 2px red; font-size: 18px; -webkit-transform: rotate(-20deg);")
	txt.SetContent("DUI Test")

	var flag bool
	hello.Connect("click", func(e dui.MouseEvent) {
		fmt.Printf("x:%d, y:%d\n", e.X, e.Y)
		flag = !flag
		if flag {
			txt.SetContent("one two three")
			println(txt.Content())
		} else {
			txt.SetContent("three two one")
			println(txt.Content())
		}
		runtime.GC()
	})

	dui.Run()
}
