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

[An example can be seen here.](https://github.com/limafresh/messagebubble/blob/main/example/main.go)

```go
bubble := messagebubble.NewMessageBubble(
	"Alex", // sender's name
	"Hi, how are you?", // message text
	"13:36", // sending time
	false, // mine: true, other: false
)
```

## Customization

If the widget has already been rendered, you need to call `name.Refresh()` after customization, where `name` is the name of the widget or container in which it resides.

### Colors

```go
bubble.Colors.Bubble.Mine = []color.NRGBA{
	{255, 51, 0, 255}, // for light theme
	{153, 0, 255, 255}, // for dark theme
}
```

**Available:** *Bubble.Mine*, *Bubble.Other*, *Text.Mine*, *Text.Other*, *Time.Mine*, *Time.Other*.

### Corner radius

```go
bubble.CornerRadius = 16
```

## Customization bubbles in container

For example, let's change `CornerRadius` of all bubbles in `vbox`:

```go
for _, obj := range vbox.Objects {
	if bubble, ok := obj.(*messagebubble.MessageBubble); ok {
		bubble.CornerRadius = 16
	}
}
```