package cmd

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	list        list.Model
	table       table.Model
	keys        *ListKeyMap
	CurrentView string
}

type ListKeyMap struct {
	toggleHelpMenu key.Binding
	chooseItem     key.Binding
}

var (
	fg     = lipgloss.AdaptiveColor{Light: "#4c4f69", Dark: "#cdd6f4"}
	bg     = lipgloss.AdaptiveColor{Light: "#eff1f5", Dark: "#1e1e2e"}
	accent = lipgloss.AdaptiveColor{Light: "#209fb5", Dark: "#74c7ec"}

	style = lipgloss.NewStyle().
		Margin(1, 2).
		Foreground(fg)
)

func NewListKeyMap() *ListKeyMap {
	return &ListKeyMap{
		toggleHelpMenu: key.NewBinding(
			key.WithKeys("H"),
			key.WithHelp("H", "toggle help"),
		),
		chooseItem: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose item"),
		),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := style.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch {
		case key.Matches(msg, m.keys.toggleHelpMenu):
			m.list.SetShowHelp(!m.list.ShowHelp())
			return m, nil

		case key.Matches(msg, m.keys.chooseItem):
			selected := m.list.SelectedItem().(Item)

			if m.CurrentView == "Collaborators" {
				m.CurrentView = selected.Title()
				m.table = GetCollaboratorView(selected.Description())
				return m, nil
			} else {
				return m, nil
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.CurrentView == "Collaborators" {
		return style.Render(m.list.View())
	} else {
		tableStyle := lipgloss.NewStyle().Inherit(style).BorderStyle(lipgloss.NormalBorder()).BorderForeground(fg)
		return tableStyle.Render(m.table.View())
	}
}
