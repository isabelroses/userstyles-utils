package cmd

import (
	"github.com/charmbracelet/bubbles/table"
	// "github.com/charmbracelet/lipgloss"
)

func GetCollaboratorView(collaborator string) table.Model {
	userstylesRoot := GetSchema()

	rows := []table.Row{}

	columns := []table.Column{
		{Title: "Name", Width: 8},
		{Title: "App Link", Width: 20},
		{Title: "Current Maintainers", Width: 20},
	}

	for i := range userstylesRoot.Userstyles {
		style := userstylesRoot.Userstyles[i]
		for j := range style.Readme.CurrentMaintainers {
			if style.Readme.CurrentMaintainers[j].Url == collaborator {
				rows = append(rows, table.Row{
					style.Name,
					style.Readme.AppLink,
					style.Readme.CurrentMaintainers[j].Name,
				})
			}
		}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.Foreground(accent).BorderBottom(true)
	s.Selected = s.Selected.Background(accent).Foreground(bg)

	t.SetStyles(s)

	return t
}
