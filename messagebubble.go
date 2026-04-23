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

const (
	DefaultCornerRadius = 12
	DefaultMaxWidth     = 300
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

// CustomLabelTheme overrides the default theme to apply custom text color
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

// bubbleLayout positions the bubble and constrains its width
type bubbleLayout struct {
	maxWidth   float32
	alignRight bool
	text       string
}

func (l *bubbleLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	obj := objects[0]

	// Calculate the total width needed for the bubble content
	objSize := obj.MinSize()
	textSize := fyne.MeasureText(l.text, theme.TextSize(), fyne.TextStyle{})
	totalWidth := objSize.Width + float32(textSize.Width)

	var width float32

	// Determine the width of the bubble based on content and maximum allowed width
	if totalWidth < l.maxWidth {
		width = totalWidth
	} else {
		width = l.maxWidth
	}

	// Set width to 80% of available width if it exceeds the container size
	if width > size.Width {
		width = size.Width * 0.8
	}

	// Resize the bubble
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
	rect        *canvas.Rectangle
	senderLabel fyne.CanvasObject
	timeLabel   *canvas.Text
	labelTheme  *CustomLabelTheme
	override    *container.ThemeOverride
	sender,
	text,
	msgTime     string
	isMine      bool

	Colors       *Colors
	CornerRadius float32
	HideSender   bool
}

func NewMessageBubble(sender, text, msgTime string, isMine bool) *MessageBubble {
	b := &MessageBubble{
		sender:  sender,
		text:    text,
		msgTime: msgTime,
		isMine:  isMine,

		Colors:       DefaultColors,
		CornerRadius: DefaultCornerRadius,
	}
	b.ExtendBaseWidget(b)

	return b
}

func (b *MessageBubble) CreateRenderer() fyne.WidgetRenderer {
	b.rect = canvas.NewRectangle(color.Transparent)
	b.rect.CornerRadius = b.CornerRadius

	b.senderLabel = widget.NewLabelWithStyle(b.sender, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	messageLabel := widget.NewLabel(b.text)
	messageLabel.Wrapping = fyne.TextWrapWord
	messageLabel.Selectable = true

	b.timeLabel = canvas.NewText(b.msgTime, color.Transparent)
	b.timeLabel.Alignment = fyne.TextAlignTrailing
	b.timeLabel.TextSize = 12
	timeLabelWrapper := container.New(layout.NewPaddedLayout(), b.timeLabel)

	bubbleContent := container.NewVBox(b.senderLabel, messageLabel, timeLabelWrapper)

	content := container.NewMax(b.rect, bubbleContent)
	contentWrapper := container.New(layout.NewPaddedLayout(), content)

	b.labelTheme = &CustomLabelTheme{Theme: theme.DefaultTheme(), labelColor: color.Transparent}
	b.override = container.NewThemeOverride(contentWrapper, b.labelTheme)

	cont := container.New(
		&bubbleLayout{alignRight: b.isMine, maxWidth: DefaultMaxWidth, text: b.text},
		b.override,
	)

	b.Refresh()

	return widget.NewSimpleRenderer(cont)
}

func (b *MessageBubble) Refresh() {
	if b.rect == nil {
		return
	}

	// Update colors
	b.rect.FillColor = b.getColor(b.Colors.Bubble)
	b.labelTheme.labelColor = b.getColor(b.Colors.Text)
	b.timeLabel.Color = b.getColor(b.Colors.Time)

	// Update corner radius
	if b.rect.CornerRadius != b.CornerRadius {
		b.rect.CornerRadius = b.CornerRadius
	}

	// Update sender label visibility
	if b.HideSender {
		b.senderLabel.Hide()
	} else {
		if b.isMine {
			b.senderLabel.Hide()
		} else {
			b.senderLabel.Show()
		}
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

	// 0x7FFF is ~50% brightness
	return brightness < 0x7FFF
}