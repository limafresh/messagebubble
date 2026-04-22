# messagebubble

Modern message bubble widget for Go Fyne GUI toolkit.

## Why messagebubble?

- Automatically update colors when switching between light/dark themes
- Set colors for the bubble, message, and time for light and dark themes
- Support for text selection, which is probably important for any chat
- Changing the corner radius

## Install

```
go get github.com/limafresh/messagebubble@latest
```

![demo](https://raw.githubusercontent.com/limafresh/messagebubble/main/demo.gif)

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

### Customization after creation

If the widget has already been rendered, you need to call `name.Refresh()` after customization, where `name` is the name of the widget or container in which it resides.

#### Colors

If you want to change the color after creation (for example, `Bubble.Mine`), just do:

```go
bubbleColors.Bubble.Mine = []color.NRGBA{{255, 51, 0, 255}, {153, 0, 255, 255}}
```

#### Corner radius

```go
bubble.CornerRadius = 16
```

or

```go
for _, obj := range vbox.Objects {
	if bubble, ok := obj.(*messagebubble.MessageBubble); ok {
		bubble.CornerRadius = 16
	}
}
```