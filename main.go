package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"gotype/models"
)

func main() {
/* 
 * Debug Logger - Included for informational purposes but not necessary for functionality. 
 * You can delete this code as it's only for debugging and can be safely removed.
 * If you decide to keep it, it logs debug output to a file (debug.log) when the DEBUG environment variable is set.

	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}
*/
	p := tea.NewProgram(
		models.NewAppModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	if _, err := p.Run(); err != nil {
		print( fmt.Errorf( "error : %v", err ) )
	}
}

