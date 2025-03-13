package catppuccin

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	ctp "github.com/catppuccin/go"
)

// Just an alias to the original catppuccin/go type
type Flavor ctp.Flavor

// Just an alias to the original catppuccin/go vars
var (
	Latte     = ctp.Latte
	Frappe    = ctp.Frappe
	Macchiato = ctp.Macchiato
	Mocha     = ctp.Mocha
)

type Color string

// Possible colors for accent color
const (
	Rosewater Color = "rosewater"
	Flamingo  Color = "flamingo"
	Pink      Color = "pink"
	Mauve     Color = "mauve"
	Red       Color = "red"
	Maroon    Color = "maroon"
	Peach     Color = "peach"
	Yellow    Color = "yellow"
	Green     Color = "green"
	Teal      Color = "teal"
	Sky       Color = "sky"
	Sapphire  Color = "sapphire"
	Blue      Color = "blue"
	Lavender  Color = "lavender"
	Base      Color = "Base"
	Text      Color = "Text"
)

func getAccent(f Flavor, c Color) ctp.Color {
	switch c {
	case Rosewater:
		return f.Rosewater()
	case Flamingo:
		return f.Flamingo()
	case Pink:
		return f.Pink()
	case Mauve:
		return f.Mauve()
	case Red:
		return f.Red()
	case Maroon:
		return f.Maroon()
	case Peach:
		return f.Peach()
	case Yellow:
		return f.Yellow()
	case Green:
		return f.Green()
	case Teal:
		return f.Teal()
	case Sky:
		return f.Sky()
	case Sapphire:
		return f.Sapphire()
	case Blue:
		return f.Blue()
	case Lavender:
		return f.Lavender()
	}
	return f.Blue()
}

func colorWithAlpha(c ctp.Color, a uint8) color.Color {
	return color.NRGBA{c.RGB[0], c.RGB[1], c.RGB[2], a}
}

// The main theme struct necessary for app.Settings().SetTheme(...)
type Theme struct {
	flavor Flavor
	accent Color
}

// Creates new Theme with the accent color set to blue
func New() *Theme {
	return &Theme{accent: Mauve}
}

// Creates new Theme with the chosen accent color
func NewWithAccent(accent Color) Theme {
	return Theme{accent: accent}
}

// Remove any forced flavor, returns to Latte for light mode systems and Mocha for dark mode
func (ctp *Theme) ResetFlavor() {
	ctp.flavor = nil
}

// Force a specific flavor of Catppuccin to the theme
func (ctp *Theme) SetFlavor(f Flavor) {
	ctp.flavor = f
}

// Chooses a different accent color to set as Primary
func (ctp *Theme) SetAccent(c Color) {
	ctp.accent = c
}

func (c *Theme) getFlavor(variant fyne.ThemeVariant) (Flavor, fyne.ThemeVariant) {
	if c.flavor == nil {
		if variant == theme.VariantLight {
			return ctp.Latte, variant
		} else {
			return ctp.Mocha, variant
		}
	} else {
		v := theme.VariantDark
		if c.flavor == Latte {
			v = theme.VariantLight
		}
		return c.flavor, v
	}
}

// Get a color from the named colors in catppuccin, useful to color in canvas elements in a consistent way
func (ctp Theme) GetNamedColor(c Color) color.Color {
	f, _ := ctp.getFlavor(fyne.CurrentApp().Settings().ThemeVariant())
	switch c {
	case Text:
		return f.Text()
	case Base:
		return f.Base()
	default:
		return getAccent(f, c)
	}
}

var _ fyne.Theme = (*Theme)(nil)

// Implements fyne.Theme
//
// Gets color from current theme with current accent and returns it for the app to handle
func (c Theme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	f, v := c.getFlavor(variant)
	accent := f.Blue()
	if c.accent != "" {
		accent = getAccent(f, c.accent)
	}
	switch name {
	case theme.ColorNameBackground:
		return f.Base()
	case theme.ColorNameButton:
		return f.Mantle()
	case theme.ColorNameDisabledButton:
		return f.Surface0()
	case theme.ColorNameDisabled:
		return f.Text()
	case theme.ColorNameError:
		return f.Red()
	case theme.ColorNameFocus:
		return colorWithAlpha(accent, 0x2a)
	case theme.ColorNameForeground:
		return f.Text()
	case theme.ColorNameForegroundOnError:
		return f.Base()
	case theme.ColorNameForegroundOnPrimary:
		return f.Base()
	case theme.ColorNameForegroundOnSuccess:
		return f.Base()
	case theme.ColorNameForegroundOnWarning:
		return f.Base()
	case theme.ColorNameHeaderBackground:
		return f.Surface0()
	case theme.ColorNameHover:
		return colorWithAlpha(f.Overlay2(), 0x0f)
	case theme.ColorNameHyperlink:
		return f.Blue()
	case theme.ColorNameInputBackground:
		return f.Mantle()
	case theme.ColorNameInputBorder:
		return f.Surface0()
	case theme.ColorNameMenuBackground:
		return f.Mantle()
	case theme.ColorNameOverlayBackground:
		return f.Overlay0()
	case theme.ColorNamePlaceHolder:
		return f.Overlay2()
	case theme.ColorNamePressed:
		return colorWithAlpha(f.Surface0(), 0x66)
	case theme.ColorNamePrimary:
		return accent
	case theme.ColorNameScrollBar:
		return colorWithAlpha(f.Mantle(), 0x99)
	case theme.ColorNameSelection:
		return colorWithAlpha(accent, 0x40)
	case theme.ColorNameSeparator:
		return f.Crust()
	case theme.ColorNameShadow:
		return theme.DefaultTheme().Color(name, v)
	case theme.ColorNameSuccess:
		return f.Green()
	case theme.ColorNameWarning:
		return f.Yellow()
	}
	return theme.DefaultTheme().Color(name, variant)
}

// Implements fyne.Theme
func (ctp Theme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

// Implements fyne.Theme
func (ctp Theme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

// Implements fyne.Theme
func (ctp Theme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
