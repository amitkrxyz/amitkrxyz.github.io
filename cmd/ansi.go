package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func generateAnsi(c Config) string {
	linkStyle := lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("#7aa2f7"))

	baseStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#565f89"))

	heroLogo := baseStyle.
		Foreground(lipgloss.Color("#c0caf5")).
		BorderStyle(lipgloss.Border{}).
		MarginLeft(1).
		PaddingBottom(0).
		Render(c.AsciiArt)

	heroDesc := baseStyle.
		Foreground(lipgloss.Color("#bb9af7")).
		Border(lipgloss.HiddenBorder(), true, false, true, false).
		PaddingLeft(2).
		Render(fmt.Sprintf("%s\ncurl %s\nSource: %s", c.Username, linkStyle.Render(c.Url), linkStyle.Render(c.Source)))
	// "amitkrxyz\ncurl " + linkStyle.Render("https://amitkr.xyz") + "\nSource: " + linkStyle.Render("https://github.com/amitkrxyz/amitkrxyz.github.io")

	about := baseStyle.
		Width(72).
		Border(lipgloss.HiddenBorder(), true, false, true, false).
		PaddingLeft(1).
		Render(strings.TrimSpace(c.About))

	linksName, linksUrl := LinksToString(c.Links)

	social := lipgloss.JoinHorizontal(lipgloss.Top,
		baseStyle.
			Bold(true).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#7aa2f7")).
			Foreground(lipgloss.Color("#bb9af7")).
			Underline(true).
			Padding(0, 1, 0, 1).
			Render(strings.Join(linksName, "\n")),
		baseStyle.
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#7aa2f7")).
			Foreground(lipgloss.Color("#bb9af7")).
			MarginLeft(0).
			BorderTop(true).
			BorderLeft(false).
			BorderRight(true).
			BorderBottom(true).
			PaddingLeft(1).
			PaddingRight(1).
			Render(strings.Join(linksUrl, "\n")),
	)

	// section := lipgloss.JoinHorizontal(lipgloss.Top, about, social)

	ansiText := lipgloss.JoinVertical(lipgloss.Left, heroDesc, heroLogo, about, social)

	return ansiText
}
