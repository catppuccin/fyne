package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/catppuccin/fyne"
)

func ColorSwatch(ctp *catppuccin.Theme) fyne.CanvasObject {
	accents := []catppuccin.Color{
		catppuccin.Rosewater,
		catppuccin.Flamingo,
		catppuccin.Pink,
		catppuccin.Mauve,
		catppuccin.Red,
		catppuccin.Maroon,
		catppuccin.Peach,
		catppuccin.Yellow,
		catppuccin.Green,
		catppuccin.Teal,
		catppuccin.Sky,
		catppuccin.Sapphire,
		catppuccin.Blue,
		catppuccin.Lavender,
	}
	grid := container.NewGridWithColumns(len(accents))
	for _, c := range accents {
		clr := NewColorRectangle(c, ctp)
		grid.Add(clr)
	}
	return grid
}
