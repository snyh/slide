package main

import "fmt"
import "pkg.linuxdeepin.com/lib/gudev"
import "pkg.linuxdeepin.com/lib/glib-2.0"

func main() {
	c := gudev.NewClient(nil)
	c.Connect("uevent", func(client *gudev.Client, action string, dev *gudev.Device) {
		fmt.Println("A:", action, "D:", dev.GetSysfsPath(), "Brightness:", dev.GetSysfsAttr("brightness"))
	})

	all := c.QueryBySubsystem("backlight")
	for _, bl := range all {
		fmt.Printf(`name: %v
path: %v
brightness: %v
type:%v

`, bl.GetName(), bl.GetSysfsPath(), bl.GetSysfsAttr("brightness"), bl.GetSysfsAttr("type"))
	}
	glib.StartLoop()
}
