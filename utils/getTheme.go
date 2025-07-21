package utils

import (
	"os/exec"
	"strings"

	catppuccin "github.com/catppuccin/go"
)

type Theme struct {
	Accent1 catppuccin.Color
	Accent2 catppuccin.Color
	Text    catppuccin.Color
}

func isDarkMode() bool {
	out, err := exec.Command("gsettings", "get", "org.gnome.desktop.interface", "color-scheme").Output()
	if err != nil {
		return true // fallback to dark
	}
	s := strings.TrimSpace(string(out))
	return strings.Contains(strings.ToLower(s), "dark")
}

func GetThemeStyle() Theme {
	var t Theme

	if isDarkMode() {
		t = Theme{
			Accent1: catppuccin.Frappe.Sapphire(),
			Accent2: catppuccin.Frappe.Mauve(),
			Text:    catppuccin.Frappe.Text(),
		}
	} else {
		t = Theme{
			Accent1: catppuccin.Latte.Sapphire(),
			Accent2: catppuccin.Latte.Mauve(),
			Text:    catppuccin.Latte.Text(),
		}
	}

	return t
}
