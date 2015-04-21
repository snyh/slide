package main

/*
#cgo pkg-config: gtk+-3.0
#include <gtk/gtk.h>
void write_c_function_in_go_source(GtkWidget* w)
{
	GtkWidget* color = gtk_color_chooser_widget_new();
	gtk_window_set_position(GTK_WINDOW(w), GTK_WIN_POS_CENTER);
	gtk_container_add(GTK_CONTAINER(w), color);
}
*/
import "C"

func main() {
	C.gtk_init(nil, nil)
	w := C.gtk_window_new(C.GTK_WINDOW_TOPLEVEL)
	C.write_c_function_in_go_source(w)
	C.gtk_widget_show_all(w)
	C.gtk_main()
}
