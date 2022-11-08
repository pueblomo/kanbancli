package taskoverview

import (
	"pueblomo/kanbancli/global"
	item "pueblomo/kanbancli/model"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))

	titelStyle = lipgloss.NewStyle().
			Bold(true).
			Background(lipgloss.Color("62")).
			Padding(0, 1).
			Foreground(lipgloss.Color("7"))

	tagStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("242")).
			MarginBottom(1)
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
		focusedStyle.Height(msg.Height - global.Divisor - 2)
	case item.TaskMsg:
		if msg.Type == global.Show {
			m.task = msg.Task
		}
	}

	return m, nil
}

func (m *model) View() string {
	titleRender := ""
	if len(m.task.Title()) > 0 {
		titleRender = titelStyle.Render(m.task.Title())
	}
	return focusedStyle.Render(lipgloss.JoinVertical(
		lipgloss.Top,
		titleRender,
		tagStyle.Render(m.task.Tag),
		m.task.Desc,
	))
}

func New() *model {
	return &model{}
}
