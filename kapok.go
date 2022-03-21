package main

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

func main() {
	label := gtk.NewLabel("Kapok")
	
	label.Show()

	window := gtk.NewWindow(gtk.WindowToplevel)
	window.ConnectDestroy(gtk.MainQuit)
	window.SetTitle("Kapok")
	window.Add(label)
	window.SetDefaultSize(400, 300)
	window.Show()

	gtk.Main()
}