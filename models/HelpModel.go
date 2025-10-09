// models/HelpModel.go

package models

import (
	"gotype/conf"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type HelpModel struct {
	width  int
	height int
	config conf.Config
}

func ( m HelpModel ) Init() tea.Cmd {
	return nil
}

func NewHelpModel( c conf.Config ) HelpModel { 
	return HelpModel{
		config: c,
	}
}

func ( m HelpModel ) Update( msg tea.Msg ) ( tea.Model, tea.Cmd ) {
	return m, nil
}

func ( m HelpModel ) View() string {

	content := "This is the Help page.\n\nPress 'q' to quit or 'esc' to return to the menu."
	

	helpContent := lipgloss.NewStyle().
		Height(m.height - 1).
		Width(m.width - 2).
		Render(content)

	return helpContent
}

func ( m HelpModel ) handleResize( height, width int ) HelpModel {
	m.height = height
	m.width = width
	return m
}
