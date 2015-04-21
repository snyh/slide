package main

/*
#include <time.h>
#include <stdlib.h>
struct tm** tms;
void init_array_times()
{
	struct tm* local, *gmt;
	tms = malloc(sizeof(struct tm*) * 2);
	time_t now = time(0);
	tms[0] = gmtime(&now);
	tms[1] = localtime(&now);
}
*/
import "C"
import "fmt"
import "unsafe"
import "reflect"

func main() {
	C.init_array_times()
	tms := (*[4](*C.struct_tm))(unsafe.Pointer(C.tms))
	fmt.Println("type of C.tms:", reflect.TypeOf(C.tms))
	fmt.Println("type of tms:", reflect.TypeOf(tms))
	fmt.Println("This will OK", tms[0].tm_year, len(tms))
}
