package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

func main() {
	app := app.NewWithID("beyt")
	var desk desktop.App
	var ok bool
	if desk, ok = app.(desktop.App); ok {
		m := fyne.NewMenu("beyt")
		desk.SetSystemTrayMenu(m)
	}

	go func() {
		for {
			t := Beat(time.Now().Unix())
			iconImg := buildImage(t)
			desk.SetSystemTrayIcon(iconImg)
			time.Sleep(1 * time.Second)
		}
	}()

	app.Run()
}
