package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func printIntr(inf interface{}) string {
	var out string

	switch v := inf.(type) {
	case string:
		out = v

	case []interface{}:
		var infs []string
		for _, n := range v {
			infs = append(infs, fmt.Sprintf("%v", n))
		}
		out = strings.Join(infs, ", ")
	}

	return out
}

func GetCollaboratorView(collaborator string) table.Model {
	userstylesRoot := GetSchema()

	rows := []table.Row{}

	columns := []table.Column{
		{Title: "Name", Width: 15},
		{Title: "Category", Width: 20},
		{Title: "App Link", Width: 50},
		{Title: "Current Maintainers", Width: 50},
	}

	for i := range userstylesRoot.Userstyles {
		// this asignment happens to save me typing the same words over and over
		us := userstylesRoot.Userstyles[i]
		cm := us.Readme.CurrentMaintainers

		var maintStr string
		for j := range cm {
			if j == 0 {
				maintStr = InferCollaboratorName(cm[j])
			} else {
				maintStr = maintStr + ", " + InferCollaboratorName(cm[j])
			}
		}

		for j := range cm {
			if cm[j].Url == collaborator {
				rows = append(rows, table.Row{
					printIntr(us.Name),
					strings.Join(us.Category, ", "),
					strings.Join(us.Readme.AppLink, ", "),
					maintStr,
				})
			}
		}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(20),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.BorderStyle(lipgloss.NormalBorder()).Foreground(accent).BorderBottom(true)
	s.Selected = s.Selected.Background(accent).Foreground(bg)

	t.SetStyles(s)

	return t
}
