package widgets

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/catppuccin/fyne"
)

type ColorRectangle struct {
	widget.BaseWidget
	lock sync.RWMutex

	ctp   *catppuccin.Theme
	color catppuccin.Color
	rect  *canvas.Rectangle
}

func NewColorRectangle(c catppuccin.Color, ctp *catppuccin.Theme) *ColorRectangle {
	rect := canvas.NewRectangle(ctp.GetNamedColor(c))
	cr := &ColorRectangle{lock: sync.RWMutex{}, ctp: ctp, color: c, rect: rect}
	cr.ExtendBaseWidget(cr)
	return cr
}

func (cr *ColorRectangle) SetColor(color catppuccin.Color) {
	cr.lock.Lock()
	cr.color = color
	cr.lock.Unlock()

	cr.Refresh()
}

func (cr *ColorRectangle) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(cr.rect)
}

func (cr *ColorRectangle) MinSize() fyne.Size {
	return fyne.NewSize(30, 10)
}

func (cr *ColorRectangle) Refresh() {
	cr.lock.Lock()
	cr.rect.FillColor = cr.ctp.GetNamedColor(cr.color)
	cr.lock.Unlock()

	cr.rect.Refresh()
}
