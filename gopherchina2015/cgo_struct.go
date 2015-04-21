package main

/*
#include <time.h>
#include <stdlib.h>
struct tm* tms[2];
void init_array_times()
{
	struct tm* local, *gmt;
	time_t now = time(0);
	tms[0] = gmtime(&now);
	tms[1] = localtime(&now);
}
*/
import "C"
import "fmt"

func main() {
	C.init_array_times()
	fmt.Println("Tody is", C.tms[1].tm_year+1900, C.tms[1].tm_mon+1, C.tms[1].tm_mday+1)
}
