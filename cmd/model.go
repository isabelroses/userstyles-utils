package cmd

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Tabs       []string
	TabContent []string
	activeTab  int
}

type KeyMap struct {
	toggleHelpMenu key.Binding
	chooseItem     key.Binding
	quit           key.Binding
}

var (
	fg     = lipgloss.AdaptiveColor{Light: "#4c4f69", Dark: "#cdd6f4"}
	bg     = lipgloss.AdaptiveColor{Light: "#eff1f5", Dark: "#1e1e2e"}
	accent = lipgloss.AdaptiveColor{Light: "#209fb5", Dark: "#74c7ec"}

	inactiveTabStyle = lipgloss.NewStyle().Padding(0, 1)
	activeTabStyle   = inactiveTabStyle.Copy().Background(accent).Foreground(bg)

	docStyle    = lipgloss.NewStyle().Padding(1, 2)
	windowStyle = lipgloss.NewStyle().Margin(1, 0)
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "right", "l", "n", "tab":
			m.activeTab = min(m.activeTab+1, len(m.Tabs)-1)
			return m, nil

		case "left", "h", "p", "shift+tab":
			m.activeTab = max(m.activeTab-1, 0)
			return m, nil
		}
	}

	return m, nil
}

func (m Model) View() string {
	doc := strings.Builder{}

	var renderedTabs []string

	for i, t := range m.Tabs {
		var style lipgloss.Style
		isActive := i == m.activeTab
		if isActive {
			style = activeTabStyle.Copy()
		} else {
			style = inactiveTabStyle.Copy()
		}
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	doc.WriteString("\n")
	doc.WriteString(windowStyle.Render(m.TabContent[m.activeTab]))
	return docStyle.Render(doc.String())
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
