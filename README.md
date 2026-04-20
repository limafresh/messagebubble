# messagebubble

Modern message bubble widget for Go Fyne GUI toolkit. Useful for chat applications.

## Install

```
go get github.com/limafresh/messagebubble@latest
```

![screenshot](https://raw.githubusercontent.com/limafresh/messagebubble/main/screenshot.png)

## Usage

[A example can be seen here.](https://github.com/limafresh/messagebubble/blob/main/example/main.go)

First you need to create the colors of the bubble.

```go
bubbleColors := &messagebubble.BubbleColors{
	Bubble: messagebubble.ColorSet{
		Mine: []color.NRGBA{
			{204, 255, 204, 255}, // light
			{0, 102, 0, 255}, // dark
		},
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
```

And then use them in the bubble.

```go
bubble := messagebubble.NewMessageBubble(
	"Alex", // sender's name
	"Hi, how are you?", // message text
	"13:36", // sending time
	false, // mine: true, other: false
	bubbleColors, // color scheme
)
```

If you want to change the color after creation, just do:

```go
bubbleColors.Bubble.Mine = []color.NRGBA{{255, 51, 0, 255}, {153, 0, 255, 255}}
bubble.Refresh()
```