package main

import (
	"example/widgets"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	catppuccin "github.com/catppuccin/fyne"
)

type noPadding struct {
	original fyne.Theme
}

var _ fyne.Theme = noPadding{}

// Implements fyne.Theme
func (np noPadding) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return np.original.Color(name, variant)
}

// Implements fyne.Theme
func (np noPadding) Font(style fyne.TextStyle) fyne.Resource {
	return np.original.Font(style)
}

// Implements fyne.Theme
func (np noPadding) Icon(name fyne.ThemeIconName) fyne.Resource {
	return np.original.Icon(name)
}

// Implements fyne.Theme
func (np noPadding) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNamePadding {
		return 0
	}
	return np.original.Size(name)
}

func main() {
	a := app.New()
	ctp := catppuccin.New()
	a.Settings().SetTheme(ctp)

	win := a.NewWindow("Catppuccin Theme Example")
	win.Resize(fyne.NewSize(500, 500))
	headerText := widget.NewRichTextFromMarkdown("# Catppuccin For Fyne")
	swatch := widgets.ColorSwatch(ctp)

	header := container.NewVBox(headerText, swatch)

	DisabledButton := widget.NewButton("Disabled Color", func() {})
	DisabledButton.Disable()

	NormalButton := widget.NewButton("Normal Color", func() {})

	PrimaryButton := widget.NewButton("Primary Color", func() {})
	PrimaryButton.Importance = widget.HighImportance

	SuccessButton := widget.NewButton("Success Color", func() {})
	SuccessButton.Importance = widget.SuccessImportance

	WarningButton := widget.NewButton("Warning Color", func() {})
	WarningButton.Importance = widget.WarningImportance

	ErrorButton := widget.NewButton("Error Color", func() {})
	ErrorButton.Importance = widget.DangerImportance

	buttons := container.NewVBox(layout.NewSpacer(), DisabledButton, NormalButton, PrimaryButton, SuccessButton, WarningButton, ErrorButton, layout.NewSpacer())

	themeForm := widget.NewForm()

	selectFlavors := widget.NewSelect([]string{"System Default", "Latte", "Frappe", "Macchiato", "Mocha"}, func(s string) {
		switch s {
		case "Latte":
			ctp.SetFlavor(catppuccin.Latte)
		case "Frappe":
			ctp.SetFlavor(catppuccin.Frappe)
		case "Macchiato":
			ctp.SetFlavor(catppuccin.Macchiato)
		case "Mocha":
			ctp.SetFlavor(catppuccin.Mocha)
		default:
			ctp.ResetFlavor()
		}
		a.Settings().SetTheme(ctp)
	})
	selectFlavors.SetSelected("System Default")
	themeForm.Append("Select Flavor", selectFlavors)

	accents := []string{
		"Rosewater",
		"Flamingo",
		"Pink",
		"Mauve",
		"Red",
		"Maroon",
		"Peach",
		"Yellow",
		"Green",
		"Teal",
		"Sky",
		"Sapphire",
		"Blue",
		"Lavender",
	}
	selectAccent := widget.NewSelect(accents, func(s string) {
		switch s {
		case "Rosewater":
			ctp.SetAccent(catppuccin.Rosewater)
		case "Flamingo":
			ctp.SetAccent(catppuccin.Flamingo)
		case "Pink":
			ctp.SetAccent(catppuccin.Pink)
		case "Mauve":
			ctp.SetAccent(catppuccin.Mauve)
		case "Red":
			ctp.SetAccent(catppuccin.Red)
		case "Maroon":
			ctp.SetAccent(catppuccin.Maroon)
		case "Peach":
			ctp.SetAccent(catppuccin.Peach)
		case "Yellow":
			ctp.SetAccent(catppuccin.Yellow)
		case "Green":
			ctp.SetAccent(catppuccin.Green)
		case "Teal":
			ctp.SetAccent(catppuccin.Teal)
		case "Sky":
			ctp.SetAccent(catppuccin.Sky)
		case "Sapphire":
			ctp.SetAccent(catppuccin.Sapphire)
		case "Blue":
			ctp.SetAccent(catppuccin.Blue)
		case "Lavender":
			ctp.SetAccent(catppuccin.Lavender)
		}
		a.Settings().SetTheme(ctp)
	})
	selectAccent.SetSelected("Blue")
	themeForm.Append("Select Accent", selectAccent)

	content := container.NewBorder(header, themeForm, nil, nil, buttons)
	win.SetContent(content)
	win.ShowAndRun()
}
