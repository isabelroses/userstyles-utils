package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func Init() {
	userstylesRoot := GetSchema()

	items := []list.Item{}

	for _, collaborators := range userstylesRoot.Collaborators {
		if collaborators.Name != "" {
			items = append(items, Item{title: collaborators.Name, desc: collaborators.Url})
		} else {
			inferedName := strings.Replace(collaborators.Url, "https://github.com/", "", -1)
			items = append(items, Item{title: inferedName, desc: collaborators.Url})
		}
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Collaborators"
	l.Styles.Title = lipgloss.NewStyle().Foreground(bg).Background(accent)

	listKeys := NewListKeyMap()

	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.toggleHelpMenu,
		}
	}

	m := Model{list: l, keys: listKeys, CurrentView: "Collaborators"}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
