package main

import "pkg.linuxdeepin.com/lib/dbus"
import "fmt"
import "time"

type DService struct {
	Name    string
	Version string

	Enabled bool `access:"readwrite"`

	Changed func(int64) //signals

	service []string
}

func (d *DService) StopWork() {
	dbus.NotifyChange(d, "Working")
}

func (d *DService) beginWork() {
	for {
		<-time.After(time.Second * 1)
		t := time.Now()
		if d.Changed != nil {
			//used method invoke to send dbus signal
			d.Changed(t.Unix())
		}
	}
}

func (d *DService) IsSupport(name string) bool {
	for _, sname := range d.service {
		if name == sname {
			return true
		}
	}
	return false
}

//give the dbus service, path and interface information
func (*DService) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		"org.snyh.test",
		"/org/snyh/test",
		"org.snyh.test",
	}
}

func main() {
	d := &DService{"TestSerivce", "0.1", false, nil, []string{"Close", "Check"}}

	if err := dbus.InstallOnSession(d); err != nil {
		fmt.Println("Install failed:", err)
		return
	}
	if err := dbus.Wait(); err != nil {
		fmt.Println("Bus error:", err)
		return
	}
}
