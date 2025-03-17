package catppuccin

import (
	"testing"

	"fyne.io/fyne/v2/theme"
	"github.com/stretchr/testify/assert"
)

func TestColorDefaults(t *testing.T) {
	ctp := New()

	color := ctp.Color(theme.ColorNamePrimary, theme.VariantLight)
	assert.Equal(t, color, Latte.Blue())

	color = ctp.Color(theme.ColorNamePrimary, theme.VariantDark)
	assert.Equal(t, color, Mocha.Blue())
}

func TestColorVariants(t *testing.T) {
	ctp := New()

	// test all with a system default of dark
	ctp.SetFlavor(Latte)
	color := ctp.Color(theme.ColorNamePrimary, theme.VariantDark)
	assert.Equal(t, color, Latte.Blue())

	ctp.SetFlavor(Frappe)
	color = ctp.Color(theme.ColorNamePrimary, theme.VariantDark)
	assert.Equal(t, color, Frappe.Blue())

	ctp.SetFlavor(Macchiato)
	color = ctp.Color(theme.ColorNamePrimary, theme.VariantDark)
	assert.Equal(t, color, Macchiato.Blue())

	ctp.SetFlavor(Mocha)
	color = ctp.Color(theme.ColorNamePrimary, theme.VariantDark)
	assert.Equal(t, color, Mocha.Blue())

	// test all with a system default of light
	ctp.SetFlavor(Latte)
	color = ctp.Color(theme.ColorNamePrimary, theme.VariantLight)
	assert.Equal(t, color, Latte.Blue())

	ctp.SetFlavor(Frappe)
	color = ctp.Color(theme.ColorNamePrimary, theme.VariantLight)
	assert.Equal(t, color, Frappe.Blue())

	ctp.SetFlavor(Macchiato)
	color = ctp.Color(theme.ColorNamePrimary, theme.VariantLight)
	assert.Equal(t, color, Macchiato.Blue())

	ctp.SetFlavor(Mocha)
	color = ctp.Color(theme.ColorNamePrimary, theme.VariantLight)
	assert.Equal(t, color, Mocha.Blue())
}

func TestColorAccents(t *testing.T) {
	ctp := NewWithAccent(Mauve)

	color := ctp.Color(theme.ColorNamePrimary, theme.VariantLight)
	assert.Equal(t, color, Latte.Mauve())

	color = ctp.Color(theme.ColorNamePrimary, theme.VariantDark)
	assert.Equal(t, color, Mocha.Mauve())

	ctp.SetAccent(Flamingo)

	color = ctp.Color(theme.ColorNamePrimary, theme.VariantLight)
	assert.Equal(t, color, Latte.Flamingo())

	color = ctp.Color(theme.ColorNamePrimary, theme.VariantDark)
	assert.Equal(t, color, Mocha.Flamingo())
}
