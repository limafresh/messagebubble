package main

import (
	"image/color"

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

	bubbleColors := &messagebubble.BubbleColors{
		Bubble: messagebubble.ColorSet{
			Mine: []color.NRGBA{{204, 255, 204, 255}, {0, 102, 0, 255}},
			Other: []color.NRGBA{{230, 230, 230, 255}, {38, 38, 38, 255}},
		},
		Text: messagebubble.ColorSet{
			Mine: []color.NRGBA{{0, 0, 0, 255}, {255, 255, 255, 255}},
			Other: []color.NRGBA{{0, 0, 0, 255}, {255, 255, 255, 255}},
		},
		Time: messagebubble.ColorSet{
			Mine: []color.NRGBA{{51, 51, 51, 255}, {230, 230, 230, 255}},
			Other: []color.NRGBA{{51, 51, 51, 255}, {230, 230, 230, 255}},
		},
	}

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
			bubbleColors,
		),
		messagebubble.NewMessageBubble(
			"Alex",
			"Hi, I'm fine.",
			"13:39",
			false,
			bubbleColors,
		),
		messagebubble.NewMessageBubble(
			"",
			"What are your plans for tomorrow? Maybe we can meet?",
			"13:41",
			true,
			bubbleColors,
		),
		messagebubble.NewMessageBubble(
			"Katya",
			"I'm writing a new project.",
			"13:50",
			false,
			bubbleColors,
		),
		messagebubble.NewMessageBubble(
			"Alex",
			"I'm going to a cafe with my friends. Come on over, it'll be fun 😀",
			"13:58",
			false,
			bubbleColors,
		),
	)

	w.Resize(fyne.NewSize(400, 650))
	w.SetContent(container.NewVBox(themehbox, dialogbox))
	w.ShowAndRun()
}