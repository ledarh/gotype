// models/HelpModel.go

package models

import (
	"gotype/conf"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	//"github.com/charmbracelet/bubbles/keymap"
	//"github.com/charmbracelet/bubbles/key"
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

	switch msg := msg.(type) {
	case tea.KeyMsg:
		for _, key := range keys.ExitKeys {
			if msg.String() == key {
				cmd = ChangeView("MENU")
				return m, cmd
			}
		}
	}

	return m, nil
}




func createTextKeyMap() map[string]string {
	// Create a new keymap with Vim-like keybindings and their descriptions
	kmap := map[string]string{
		"Up":        "Move up (k, up)",
		"Down":      "Move down (j, down)",
		"Left":      "Move left (h, left)",
		"Right":     "Move right (l, right)",
		// "NextWord":  "Move to next word (w)",
		// "PrevWord":  "Move to previous word (b)",
		// "Insert":    "Enter insert mode (i)",
		// "GoTop":     "Go to top (gg)",
		// "GoBottom":  "Go to bottom (G)",
		// "Retry":     "Retry (r)",
		// "Pause":     "Pause (p)",
	}

	return kmap
}

func (m HelpModel) GetKeyText(kmap map[string]string) []string {
	leftColumn := ""
	rightColumn := ""
	for action, description := range kmap {
		leftColumn += fmt.Sprintf("%-12s\n", action)
		rightColumn += fmt.Sprintf("%-30s\n", description)
	}

	// Return both columns as a slice of strings
	return []string{leftColumn, rightColumn}
}

// func GetKeyStyle(columns []string) lipgloss.Style {

// 	style := lipgloss.NewStyle().Align(lipgloss.Center).Width(40)

// 	content := lipgloss.JoinHorizontal(
// 		lipgloss.Center,
// 		style.Render(columns[0]),
// 		style.Render(columns[1]),
// 	)

// 	return style.Render(content)
// }



func ( m HelpModel ) GetFeatStyle() lipgloss.Style {
	featStyle := lipgloss.NewStyle().
        Foreground( m.config.Colors.Color7 ).
        Width(m.width).
        Align(lipgloss.Left).
        MarginLeft(1)
	
	return featStyle
}

func ( m HelpModel ) GetFeatText() string {
	featText := `
	* Game: efe
	* Levels: fflevels
	
	`
	
	return featText
}


func ( m HelpModel ) View() string {


	// Windows
	helpTitleText := lipgloss.NewStyle().
        Foreground( m.config.Colors.Color7 ).
        Width(m.width).
        Align(lipgloss.Center).
	MarginTop(1).
		Render(
`╔╗      ╔╗     
║║      ║║     
║╚═╗╔══╗║║ ╔══╗
║╔╗║║╔╗║║║ ║╔╗║
║║║║║║═╣║╚╗║╚╝║
╚╝╚╝╚══╝╚═╝║╔═╝
           ║║  
        ╚╝`)

	helpTitleObj := lipgloss.NewStyle().
		Height(m.height-19).
		Width(m.width - 2).
		Render(helpTitleText)


	exitText := lipgloss.NewStyle().
        Foreground( m.config.Colors.Color7 ).
        Width(m.width).
        Align(lipgloss.Center).
        MarginTop(m.height-20).
		Render("Press q, esc, backspace to return to menu")

	exitTextObj := lipgloss.NewStyle().
		Bold(true).
		Render(exitText)


	featStyle := m.GetFeatStyle()
	featText  := m.GetFeatText()


	kmap := createTextKeyMap()

	columns := m.GetKeyText(kmap)
	keyStyle := lipgloss.NewStyle().Align(lipgloss.Center).Width(40)
	keyObj  := lipgloss.JoinHorizontal(
		lipgloss.Center,
		keyStyle.Render(columns[0]),
		keyStyle.Render(columns[1]),
	)


	featObj := featStyle.Render(featText)

	

    content := lipgloss.JoinVertical(
        lipgloss.Center,
        helpTitleObj,
		keyObj,
		featObj,
		exitTextObj,
	)

	return content
}







func ( m HelpModel ) handleResize( height, width int ) HelpModel {
	m.height = height
	m.width = width
	return m
}



//Render("↑/k: up • ↓/j: down • enter: select • q: quit")




/* Alternative title texts */
/*

    titleText := ` 
░█▄█▒██▀░█▒░▒█▀▄
▒█▒█░█▄▄▒█▄▄░█▀▒
`

    titleText := `
╔╗      ╔╗     
║║      ║║     
║╚═╗╔══╗║║ ╔══╗
║╔╗║║╔╗║║║ ║╔╗║
║║║║║║═╣║╚╗║╚╝║
╚╝╚╝╚══╝╚═╝║╔═╝
           ║║  
           ╚╝ 
`

    titleText := `
ＨＥＬＰ
`

*/


/*
		* Menu: Navigate the game's primary options such as starting a new game, loading a saved game
		* Game: Practice typing code-like text inside the terminal
		* Levels: Select specific level to launch game at
		* Settings: Adjust keybindings and other customizations
*/

/*
func ( m HelpModel ) GetKeyText() string {
	keyText := `
	• ↑/k: up
	• ↓/j: down
	• enter: select
	• q: quit
	`
	
	return keyText
}
*/




// func ( m HelpModel ) GetKeyMap() keymap.KeyMap {
// 	kmap := keymap.New()

// 	// Keybindings for active play
// 	kmap.Add("left", "Move left (h)")
// 	kmap.Add("down", "Move down (j)")
// 	kmap.Add("up", "Move up (k)")
// 	kmap.Add("right", "Move right (l)")
// 	kmap.Add("nextWord", "Jump to next word (w)")
// 	kmap.Add("prevWord", "Jump to previous word (b)")
// 	kmap.Add("insert", "Enter insert mode (i)") // Example: type or interact
// 	kmap.Add("exitInsert", "Exit insert mode (esc)")
// 	kmap.Add("goTop", "Go to the top (gg)") // Go to the start of the text
// 	kmap.Add("goBottom", "Go to the bottom (G)") // Go to the end of the text
// 	kmap.Add("quit", "Quit the game (q)") // Quit the program
// 	kmap.Add("retry", "Retry the test (r)") // Restart or retry the test
// 	kmap.Add("pause", "Pause the test (p)") // Pause the game or typing test
// 	kmap.Add("settings", "Open settings (s)") // Open the settings menu

// 	return kmap
// }


// func createKeymap() map[string]key.Binding {

// 	kmap := make(map[string]key.Binding)

// 	kmap["Up"] = key.NewBinding(
// 		key.WithKeys("k", "up"),
// 	)
// 	kmap["Down"] = key.NewBinding(
// 		key.WithKeys("j", "down"),
// 	)
// 	kmap["Left"] = key.NewBinding(
// 		key.WithKeys("h", "left"),
// 	)
// 	kmap["Right"] = key.NewBinding(
// 		key.WithKeys("l", "right"),
// 	)

// 	kmap["NextWord"] = key.NewBinding(
// 		key.WithKeys("w"),
// 	)
// 	kmap["PrevWord"] = key.NewBinding(
// 		key.WithKeys("b"),
// 	)
// 	kmap["Insert"] = key.NewBinding(
// 		key.WithKeys("i"),
// 	)

// 	kmap["GoTop"] = key.NewBinding(
// 		key.WithKeys("g"),
// 	)
// 	kmap["GoBottom"] = key.NewBinding(
// 		key.WithKeys("G"),
// 	)

// 	// Return the keymap
// 	return kmap
// }
