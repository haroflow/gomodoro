package main

import (
	"bytes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showIntrusiveAlert(app fyne.App) <-chan string {
	result := make(chan string, 1)

	w := app.NewWindow("Super intrusive alert")
	w.SetIcon(icon)
	w.SetFullScreen(true)

	br := bytes.NewReader(appIcon)
	img := canvas.NewImageFromReader(br, "icon")
	img.FillMode = canvas.ImageFillOriginal

	c := container.NewCenter(
		container.NewHBox(
			img,

			container.NewVBox(
				container.NewCenter(
					widget.NewLabel("Time's up!"),
				),
				container.NewHBox(
					widget.NewButton("Restart timer", func() {
						w.Close()
						result <- "RESTART_TIMER"
					}),
					widget.NewButton("Close", func() {
						w.Close()
						result <- "CLOSE"
					}),
				),
			),
		),
	)

	w.SetContent(c)

	w.Show()

	return result
}
