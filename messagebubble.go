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

var (
	DefaultColors = &Colors{
		Bubble: ColorSet{
			Mine: []color.NRGBA{{204, 255, 204, 255}, {0, 102, 0, 255},
			},
			Other: []color.NRGBA{{230, 230, 230, 255}, {38, 38, 38, 255}},
		},
		Text: ColorSet{
			Mine: []color.NRGBA{{0, 0, 0, 255}, {255, 255, 255, 255}},
			Other: []color.NRGBA{{0, 0, 0, 255}, {255, 255, 255, 255}},
		},
		Time: ColorSet{
			Mine: []color.NRGBA{{51, 51, 51, 255}, {230, 230, 230, 255}},
			Other: []color.NRGBA{{51, 51, 51, 255}, {230, 230, 230, 255}},
		},
	}
)

type ColorSet struct {
	Mine, Other []color.NRGBA
}

type Colors struct {
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

type MessageBubble struct {
	widget.BaseWidget
	rect       *canvas.Rectangle
	timeLabel  *canvas.Text
	labelTheme *CustomLabelTheme
	override   *container.ThemeOverride
	sender,
	text,
	msgTime    string
	isMine     bool

	Colors       *Colors
	CornerRadius float32
}

func NewMessageBubble(sender, text, msgTime string, isMine bool) *MessageBubble {
	b := &MessageBubble{
		sender:  sender,
		text:    text,
		msgTime: msgTime,
		isMine:  isMine,

		Colors:       DefaultColors,
		CornerRadius: 12,
	}
	b.ExtendBaseWidget(b)

	return b
}

func (b *MessageBubble) CreateRenderer() fyne.WidgetRenderer {
	b.rect = canvas.NewRectangle(color.Transparent)
	b.rect.CornerRadius = b.CornerRadius

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

	cont := container.New(&bubbleLayout{alignRight: b.isMine, maxWidth: 300, text: b.text}, b.override)

	b.Refresh()

	return widget.NewSimpleRenderer(cont)
}

func (b *MessageBubble) Refresh() {
	b.rect.FillColor = b.getColor(b.Colors.Bubble)
	b.labelTheme.labelColor = b.getColor(b.Colors.Text)
	b.timeLabel.Color = b.getColor(b.Colors.Time)

	if b.rect.CornerRadius != b.CornerRadius {
		b.rect.CornerRadius = b.CornerRadius
	}

	b.override.Refresh()
}

func (b *MessageBubble) getColor(colors ColorSet) color.NRGBA {
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