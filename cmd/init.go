package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func Init() {
	userstylesRoot := GetSchema()

	items := []list.Item{}

	for _, collaborator := range userstylesRoot.Collaborators {
		items = append(items, Item{title: InferCollaboratorName(collaborator), desc: collaborator.Url})
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.SetShowTitle(false)
	l.SetShowHelp(false)
	l.SetSize(docStyle.GetWidth(), docStyle.GetHeight())

	// setup our view
	tabs := []string{"Collaborators", "Userstyles"}
	m := MainModel{Tabs: tabs, collaboratorTab: l, userstylesTab: table.New(), activeTab: 0}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
