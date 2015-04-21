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

func main() {
	C.init_array_times()
	fmt.Println("This will failed", C.tms[1].tm_year)
}
