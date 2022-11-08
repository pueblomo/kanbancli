package form

import (
	"pueblomo/kanbancli/global"
	item "pueblomo/kanbancli/model"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var elementStyle = lipgloss.NewStyle().Padding(1).Foreground(lipgloss.Color("7"))

// TODO input of project name
type model struct {
	name     textinput.Model
	oldModel tea.Model
}

func (m model) create() tea.Msg {
	return item.NewProjectMsg(global.Create, m.name.Value())
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m.oldModel, tea.ClearScrollArea
		case "ctrl+x":
			return m.oldModel, m.create
		}
	}

	m.name, cmd = m.name.Update(msg)
	return m, cmd

}

func (m *model) View() string {
	return elementStyle.Render(m.name.View())
}

func New(oldModel tea.Model) *model {
	f := &model{}
	f.oldModel = oldModel
	f.name = textinput.New()
	f.name.Placeholder = "Project name"
	f.name.Focus()
	return f
}
