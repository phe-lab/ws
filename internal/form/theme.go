package form

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var t = huh.ThemeBase()

func init() {
	t.Focused.Title = t.Focused.Title.Foreground(lipgloss.Color("6"))
	t.Focused.FocusedButton = t.Focused.FocusedButton.Foreground(lipgloss.Color("255")).Background(lipgloss.Color("5"))
	t.Focused.BlurredButton = t.Focused.BlurredButton.Foreground(lipgloss.Color("255")).Background(lipgloss.Color("0"))
}
