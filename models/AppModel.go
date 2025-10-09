// models/AppModel.go

package models

import (
	"gotype/conf"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AppModel struct {
	width		int
	height		int
	altscreenActive bool
	StatusBar	StatusBarModel
	Menu		MenuModel
	Game 		GameModel
	Help        HelpModel
	CurrentView	string
	config 		conf.Config
	timer 		timer.Model
}

func NewAppModel() AppModel {

	c := conf.NewConfig()

	return AppModel{
		config: c,
		CurrentView: "MENU",
		StatusBar: NewStatusBarModel( c ),
		Menu: NewMenuModel( c ),
		Game: NewGameModel( c ),
		Help: NewHelpModel( c ),
	}
}

func ( m AppModel ) Init() tea.Cmd {
	return nil
}

func ( m AppModel ) Update( msg tea.Msg ) ( tea.Model, tea.Cmd ) {
	var cmd tea.Cmd
    var cmds []tea.Cmd
	
	switch msg := msg.( type ) {

		case tea.KeyMsg:
			if msg.String() == "ctrl+c" {
				return m, tea.Quit
			}
			newStatusBar, statusbarCmd := m.StatusBar.Update( msg )	
			m.StatusBar = newStatusBar.( StatusBarModel )
			cmd = statusbarCmd


		case EditorFinishedMsg:
			if msg.err != nil {
				m.CurrentView = "MENU"
				m.StatusBar.CurrentView = m.CurrentView
				m.altscreenActive = false
				cmd := tea.ExitAltScreen
				return m, cmd
			}
		case StartGameMsg:
			m.timer = timer.NewWithInterval( time.Second * 30 , time.Second )
			return m, m.timer.Init()

		case StopGameMsg:
			m.timer = timer.NewWithInterval( time.Second * 0 , time.Second )

			return m, ChangeView( "MENU" )

		case tea.WindowSizeMsg:
			m = m.handleResize( msg.Height, msg.Width )
			return m, nil
		
		case ChangeViewMsg:
			m.CurrentView = msg.View
			m.StatusBar.CurrentView = msg.View
			return m, nil

		case timer.TickMsg:
			m.timer, cmd = m.timer.Update( msg )
			m.StatusBar = m.StatusBar.updateTimeout( m.timer.Timeout.String() )
			m.Game = m.Game.updateTimeout( m.timer.Timeout.String() )
			return m, cmd

		case timer.StartStopMsg:
			m.timer, cmd = m.timer.Update( msg )
			m.StatusBar = m.StatusBar.updateTimeout( m.timer.Timeout.String() )
			m.Game = m.Game.updateTimeout( m.timer.Timeout.String() )
			return m, cmd

		case timer.TimeoutMsg:
			cmd = ChangeView( "MENU" )	

	}
	switch( m.CurrentView ) {
		case "MENU":
			newMenu, menuCmd := m.Menu.Update( msg )
			m.Menu = newMenu.( MenuModel )
			cmd = menuCmd
				
		case "Game":
			newGame, gameCmd := m.Game.Update( msg )
			m.Game = newGame.( GameModel )
			cmd = gameCmd
		case "Help":
			newHelp, helpCmd := m.Help.Update( msg )
			m.Help = newHelp.( HelpModel )
			cmd = helpCmd
						
	}

	
	cmds = append(cmds, cmd)
    return m, tea.Batch(cmds...)
}

func ( m AppModel ) View() string {

	mainWin := lipgloss.NewStyle().
		Height( m.height - 1 ).
		Width( m.width - 2 )
	mainView := ""
	switch( m.CurrentView ) {
		case "MENU":
			mainView = lipgloss.JoinVertical( 0, mainWin.Render( m.Menu.View() ), m.StatusBar.View() )
		case "Help":
			mainView = lipgloss.JoinVertical(0, mainWin.Render(m.Help.View()), m.StatusBar.View())
		case "Game":
			mainView = lipgloss.JoinVertical( 0, mainWin.Render(), m.Game.View(), m.StatusBar.View() )
	}
	
	return mainView

}

func ( m AppModel ) handleResize( height, width int ) AppModel {
	
	m.height = height
	m.width = width
	
	m.StatusBar = m.StatusBar.handleResize( height, width )
	m.Menu = m.Menu.handleResize( height, width )
	m.Game = m.Game.handleResize( height, width )
	m.Help = m.Help.handleResize( height, width )
	
	return m
}
