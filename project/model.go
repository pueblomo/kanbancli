package project

import (
	"strconv"

	"pueblomo/kanbancli/global"
	item "pueblomo/kanbancli/model"
	"pueblomo/kanbancli/project/form"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var elementStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
var focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#EE6FF8"))

type Model struct {
	Projects     []string
	Selected     int
	PrevSelected int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) createMsg() tea.Msg {
	return item.NewProjectMsg(global.Update, "")
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//TODO project change, new tea.Model with project input
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+p":
			return form.New(m).Update(nil)
		case "1":
			m.PrevSelected = m.Selected
			m.Selected = 0
			return m, m.createMsg
		case "2":
			m.PrevSelected = m.Selected
			m.Selected = 1
			return m, m.createMsg
		case "3":
			m.PrevSelected = m.Selected
			m.Selected = 2
			return m, m.createMsg
		case "4":
			m.PrevSelected = m.Selected
			m.Selected = 3
			return m, m.createMsg
		case "5":
			m.PrevSelected = m.Selected
			m.Selected = 4
			return m, m.createMsg
		}
	case item.ProjectMsg:
		if msg.Type == global.Create {
			m.Projects = append(m.Projects, msg.Project)
			return m, tea.ClearScrollArea
		}
	}
	return m, nil
}

func (m Model) View() string {
	var convertetProjects []string
	for index, project := range m.Projects {
		if index == m.Selected {
			convertetProjects = append(convertetProjects, focusedStyle.Render("<"+strconv.Itoa(index+1)+"> "+project+"    "))
		} else {
			convertetProjects = append(convertetProjects, elementStyle.Render("<"+strconv.Itoa(index+1)+"> "+project+"    "))
		}
	}
	return lipgloss.JoinVertical(lipgloss.Left, elementStyle.Render("Projects:"), lipgloss.JoinHorizontal(lipgloss.Left, convertetProjects...))
}

// TODO give list of projects
func New() *Model {
	return &Model{Projects: []string{"Nase", "Bein", "Kopf"}, Selected: 0}
}
