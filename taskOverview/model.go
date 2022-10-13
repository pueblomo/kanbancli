package taskoverview

import (
	"pueblomo/kanbancli/global"
	"pueblomo/kanbancli/item"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62"))
)

type model struct {
	task item.Task
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		focusedStyle.Width(msg.Width/global.Divisor - global.Divisor)
		focusedStyle.Height(msg.Height - global.Divisor)
	case item.Task:
		m.task = msg
	}

	return m, nil
}

func (m *model) View() string {
	return focusedStyle.Render(lipgloss.JoinVertical(
		lipgloss.Top,
		m.task.Title(),
		m.task.Description(),
	))
}

func New() *model {
	return &model{item.New("", "")}
}
