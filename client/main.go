package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/lib/pq"
)

var window fyne.Window

func updateDB() {
	controlsView := container.NewAdaptiveGrid(
		2,
		widget.NewButton("Update", updateDB),
		widget.NewButton("Update2", update2),
	)

	serverMsg := sendCommandToServer("update")
	array := parseJsonToLocationInfoArray(serverMsg)
	windowContainer := container.NewGridWithRows(2, NewLocationInfoView(array), controlsView)
	window.SetContent(windowContainer)
}

func update2() {
	sendCommandToServer("update2")
}

func main() {
	a := app.New()
	window = a.NewWindow("DB viewer")
	window.Resize(fyne.NewSize(800, 600))

	updateDB()

	window.ShowAndRun()
}

