# gomodoro
Simple pomodoro timer in golang + [fyne](https://fyne.io/).

Tested on Windows only.

![gomodoro-img](https://user-images.githubusercontent.com/4776931/116794679-3aef0200-aaa5-11eb-9e0e-18638f9a1a01.png)
> Of course... - You, seeing this ☝️

## Motivation

Looked for timers on Microsoft Store, it crashed, three times.

Found a pomodoro online, it didn't notify me.

So I made one... probably works.

## Screenshots:
![Screenshot_2](https://user-images.githubusercontent.com/4776931/116794695-65d95600-aaa5-11eb-9f2c-0e29d6d7318f.png)

![Screenshot_3](https://user-images.githubusercontent.com/4776931/116794698-696cdd00-aaa5-11eb-8179-886a066db6f2.png)

## Run

- Install go 1.16+ (requires go:embed)
- Install fyne dependencies (tested on v2.0.3): https://github.com/fyne-io/fyne
- Run fyne examples, to check it's working.
- Run the app:
```
git clone https://github.com/haroflow/gomodoro
cd gomodoro
go run .
```

## Package

Install fyne command (tested on v2.0.3):
```
go get fyne.io/fyne/v2/cmd/fyne
fyne version
```
Build package:
```
fyne package
```
Run:
```
gomodoro.exe
```

## To-do

- [ ] Play a sound when timer expires?
- [ ] The alert is pretty intrusive (which works for me). Add an option to send system notifications, in case you don't want to be disturbed by the fullscreen alert.
- [ ] Add option to go to systray when minimized. Depends on a OnMinimized event from fyne.
- [ ] When resizing, the timer stops counting down. We are doing this on another goroutine... not sure but maybe getting values from the data-binding is waiting for resize to end.
- [ ] ~Remove golang puns~
