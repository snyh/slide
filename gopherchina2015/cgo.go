package main

//#include <stdlib.h>
//#include <time.h>
import "C"
import "fmt"

func main() {
	s := C.time(nil)
	C.srandom(C.uint(s))
	r := C.random()
	fmt.Println(r)
}
