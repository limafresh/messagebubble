package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/theme"

	"github.com/limafresh/messagebubble"
)

func main() {
	a := app.New()
	w := a.NewWindow("messagebubble example")

	// Switch theme buttons
	themehbox := container.NewHBox(
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
	)

	// Example for use
	dialogbox := container.NewVBox(
		messagebubble.NewMessageBubble(
			"",
			"Hi, how are you?",
			"13:36",
			true,
		),
		messagebubble.NewMessageBubble(
			"Alex",
			"Hi, I'm fine.",
			"13:39",
			false,
		),
		messagebubble.NewMessageBubble(
			"",
			"What are your plans for tomorrow? Maybe we can meet?",
			"13:41",
			true,
		),
		messagebubble.NewMessageBubble(
			"Katya",
			"I'm writing a new project.",
			"13:50",
			false,
		),
		messagebubble.NewMessageBubble(
			"Alex",
			"I'm going to a cafe with my friends. Come on over, it'll be fun 😀",
			"13:58",
			false,
		),
	)

	w.Resize(fyne.NewSize(400, 650))
	w.SetContent(container.NewVBox(themehbox, dialogbox))
	w.ShowAndRun()
}