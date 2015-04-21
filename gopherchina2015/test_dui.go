package main

import "dui"
import "dui/css"
import "runtime"
import "reflect"
import "fmt"

func main() {
	f := dui.NewFrame(350, 500)

	//img := f.NewElement("img")
	//img.SetStyle(css.Border, "1px solid blue")
	//img.Set("height", "300")
	//img.Set("src", "/dev/shm/dui/Source/dui/test.gif")
	//f.Add(img)
	print("imgWidth:", img.Get("height"), "\n")

	var a dui.Frame
	t := reflect.ValueOf(a)
	fmt.Print(t)

	txt := f.NewElement("div")
	txt.SetContent("DUI Test测试")
	txt.SetStyle(css.TextShadow, "2px 2px 2px red")
	txt.SetStyle(css.FontSize, "18px")
	txt.SetStyle(css.Transform, "rotate(-20deg)")
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
