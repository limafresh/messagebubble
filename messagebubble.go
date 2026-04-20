package messagebubble

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/layout"
)

type ColorSet struct {
	Mine, Other []color.NRGBA
}

type BubbleColors struct {
	Bubble, Text, Time ColorSet
}

type CustomLabelTheme struct {
	fyne.Theme
	labelColor color.Color
}

func (t *CustomLabelTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameForeground {
		return t.labelColor
	}

	return t.Theme.Color(name, variant)
}

type bubbleLayout struct {
	maxWidth   float32
	alignRight bool
	text       string
}

func (l *bubbleLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	obj := objects[0]

	objSize := obj.MinSize()
	textSize := fyne.MeasureText(l.text, theme.TextSize(), fyne.TextStyle{})
	totalWidth := objSize.Width + float32(textSize.Width)

	var width float32

	if totalWidth < l.maxWidth {
		width = totalWidth
	} else {
		width = l.maxWidth
	}

	if width > size.Width {
		width = size.Width * 0.8
	}

	obj.Resize(fyne.NewSize(width, obj.MinSize().Height))

	if l.alignRight {
		obj.Move(fyne.NewPos(size.Width-obj.Size().Width, 0))
	} else {
		obj.Move(fyne.NewPos(0, 0))
	}
}

func (l *bubbleLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return objects[0].MinSize()
}

type messageBubble struct {
	widget.BaseWidget
	rect         *canvas.Rectangle
	timeLabel    *canvas.Text
	labelTheme   *CustomLabelTheme
	override     *container.ThemeOverride
	sender,
	text,
	msgTime      string
	isMine       bool
	bubbleColors *BubbleColors
}

func NewMessageBubble(sender, text, msgTime string, isMine bool, bubbleColors *BubbleColors) fyne.CanvasObject {
	b := &messageBubble{
		sender:       sender,
		text:         text,
		msgTime:      msgTime,
		isMine:       isMine,
		bubbleColors: bubbleColors,
	}
	b.ExtendBaseWidget(b)

	return container.New(&bubbleLayout{alignRight: isMine, maxWidth: 300, text: text}, b)
}

func (b *messageBubble) CreateRenderer() fyne.WidgetRenderer {
	b.rect = canvas.NewRectangle(color.Transparent)
	b.rect.CornerRadius = 12

	messageLabel := widget.NewLabel(b.text)
	messageLabel.Wrapping = fyne.TextWrapWord
	messageLabel.Selectable = true

	b.timeLabel = canvas.NewText(b.msgTime, color.Transparent)
	b.timeLabel.Alignment = fyne.TextAlignTrailing
	b.timeLabel.TextSize = 12
	timeLabelWrapper := container.New(layout.NewPaddedLayout(), b.timeLabel)

	var bubbleContent fyne.CanvasObject
	if b.isMine {
		bubbleContent = container.NewVBox(messageLabel, timeLabelWrapper)
	} else {
		senderLabel := widget.NewLabelWithStyle(b.sender, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
		bubbleContent = container.NewVBox(senderLabel, messageLabel, timeLabelWrapper)
	}

	content := container.NewMax(b.rect, bubbleContent)
	contentWrapper := container.New(layout.NewPaddedLayout(), content)

	b.labelTheme = &CustomLabelTheme{Theme: theme.DefaultTheme(), labelColor: color.Transparent}
	b.override = container.NewThemeOverride(contentWrapper, b.labelTheme)

	b.Refresh()

	return widget.NewSimpleRenderer(b.override)
}

func (b *messageBubble) Refresh() {
	b.rect.FillColor = b.getColor(b.bubbleColors.Bubble)
	b.labelTheme.labelColor = b.getColor(b.bubbleColors.Text)
	b.timeLabel.Color = b.getColor(b.bubbleColors.Time)

	b.override.Refresh()
}

func (b *messageBubble) getColor(colors ColorSet) color.NRGBA {
	bg := fyne.CurrentApp().Settings().Theme().Color(theme.ColorNameBackground, theme.VariantLight)
	if isDark(bg) {
		if b.isMine {
			return colors.Mine[1]
		}
		return colors.Other[1]
	}

	if b.isMine {
		return colors.Mine[0]
	}
	return colors.Other[0]
}

func isDark(c color.Color) bool {
	r, g, b, _ := c.RGBA()
	brightness := (r + g + b) / 3

	return brightness < 0x7FFF
}