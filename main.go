package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

// Running on Windows: go run .
// Building for Windows: go build -ldflags -H=windowsgui .
// Packaging: fyne package

func main() {
	myApp := app.NewWithID("haroflow-gomodoro")
	myApp.Settings().SetTheme(theme.DarkTheme())

	g := New(myApp)
	g.Show()
	g.Window.SetMaster()

	myApp.Run()
}
