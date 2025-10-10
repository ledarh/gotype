// models/HelpModel.go

package models

import (
	"gotype/conf"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
	ExitKeys []string
}

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
	var cmd tea.Cmd

	keys := KeyMap{
		ExitKeys: []string{"q", "esc", "backspace"},

	}

	// Check for key presses
	switch msg := msg.(type) {
	case tea.KeyMsg:
		for _, key := range keys.ExitKeys {
			if msg.String() == key {
				cmd = ChangeView("MENU")
				return m, cmd
			}
		}
	}

	// If no key matches, return the model as is
	return m, nil
}

func ( m HelpModel ) View() string {

	helpTitleText := lipgloss.NewStyle().
        Foreground( m.config.Colors.Color7 ).
        Width(m.width).
        Align(lipgloss.Center).
        MarginTop(1).
		Render("This is the Help page.\n\nPress 'q' to quit or 'esc' to return to the menu.")

	helpContent := lipgloss.NewStyle().
		Height(m.height - 1).
		Width(m.width - 2).
		Render(helpTitleText)


    content := lipgloss.JoinVertical(
        lipgloss.Center,
        //title,
        "\n",
        helpContent,
	)

	return content
}

func ( m HelpModel ) handleResize( height, width int ) HelpModel {
	m.height = height
	m.width = width
	return m
}



//Render("↑/k: up • ↓/j: down • enter: select • q: quit")
