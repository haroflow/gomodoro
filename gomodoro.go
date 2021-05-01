package main

import (
	_ "embed"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

//go:embed Icon.png
var appIcon []byte
var icon = fyne.NewStaticResource("Icon", appIcon)

type state int

const (
	Paused state = iota
	CountingDown
)

type Gomodoro struct {
	app    fyne.App
	Window fyne.Window
	timer  *time.Ticker

	// State
	currentState    state
	duration        binding.Int
	durationString  binding.String
	countdown       binding.Int
	countdownString binding.String

	// Controls
	incButton   *widget.Button
	decButton   *widget.Button
	startButton *widget.Button
	stopButton  *widget.Button
	timerLabel  *widget.Label
}

func New(app fyne.App) *Gomodoro {
	g := &Gomodoro{
		app: app,
	}

	g.initState()
	g.initWindow()
	go g.startTimer()

	return g
}

func (g *Gomodoro) refreshWindow() {
	switch g.currentState {
	case Paused:
		g.incButton.Enable()
		g.decButton.Enable()

		g.startButton.Show()
		g.stopButton.Hide()

		g.timerLabel.Bind(g.durationString)

	case CountingDown:
		g.incButton.Disable()
		g.decButton.Disable()

		g.startButton.Hide()
		g.stopButton.Show()

		g.timerLabel.Bind(g.countdownString)
	}
}

func (g *Gomodoro) Show() {
	g.Window.Show()
}

func (g *Gomodoro) initState() {
	g.duration = binding.NewInt()
	g.duration.Set(25)
	g.durationString = binding.IntToStringWithFormat(g.duration, "%02d:00")

	g.countdown = binding.NewInt()
	g.countdownString = binding.NewString()
}

func (g *Gomodoro) initWindow() {
	g.app.SetIcon(icon)

	g.Window = g.app.NewWindow("gomodoro")
	g.Window.Resize(fyne.NewSize(280, 90))
	g.Window.CenterOnScreen()

	g.timerLabel = widget.NewLabel("")
	g.timerLabel.TextStyle = fyne.TextStyle{
		Monospace: true,
	}

	g.incButton = widget.NewButton("+", g.btnIncrementDuration)
	g.decButton = widget.NewButton("-", g.btnDecrementDuration)
	g.startButton = widget.NewButton("Go!", g.btnStart)
	g.stopButton = widget.NewButton("Stop", g.btnStop)

	c := container.NewGridWithColumns(
		3,
		container.NewCenter(g.timerLabel),
		container.NewGridWithRows(2,
			g.incButton,
			g.decButton,
		),
		g.startButton,
		g.stopButton,
	)

	g.Window.SetContent(c)

	g.refreshWindow()
}

func (g *Gomodoro) btnIncrementDuration() {
	n, _ := g.duration.Get()

	if n < 5 {
		g.duration.Set(n + 1)
	} else if n < 60 {
		g.duration.Set(n + 5)
	}
}

func (g *Gomodoro) btnDecrementDuration() {
	n, _ := g.duration.Get()

	if n >= 10 {
		g.duration.Set(n - 5)
	} else if n > 1 {
		g.duration.Set(n - 1)
	}
}

func (g *Gomodoro) btnStart() {
	d, _ := g.duration.Get()
	durationSeconds := d * 60
	g.countdown.Set(durationSeconds)
	g.countdownString.Set(getTimerString(durationSeconds))

	g.timer.Reset(time.Second)

	g.currentState = CountingDown
	g.refreshWindow()
}

func (g *Gomodoro) btnStop() {
	g.Window.SetTitle("gomodoro")
	g.timer.Stop()

	g.currentState = Paused
	g.refreshWindow()
}

func (g *Gomodoro) startTimer() {
	g.timer = time.NewTicker(time.Second)
	g.timer.Stop()

	for range g.timer.C {
		n, _ := g.countdown.Get()
		n--
		g.countdown.Set(n)
		g.countdownString.Set(getTimerString(n))

		g.Window.SetTitle("gomodoro " + getTimerString(n))

		if n == 0 {
			g.btnStop()

			// TODO add setting to show only a system notification
			result := <-showIntrusiveAlert(g.app)
			if result == "RESTART_TIMER" {
				g.btnStart()
			}
		}
	}
}

func getTimerString(timeLeft int) string {
	minutes := timeLeft / 60
	seconds := timeLeft % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
