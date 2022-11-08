package main

import (
	"fmt"
	"os"
	"pueblomo/kanbancli/global"
	"pueblomo/kanbancli/kanban"
	item "pueblomo/kanbancli/model"
	"pueblomo/kanbancli/project"
	"pueblomo/kanbancli/storage"
	taskoverview "pueblomo/kanbancli/taskOverview"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var p *tea.Program

type model struct {
	projects  tea.Model
	kanban    tea.Model
	tOverview tea.Model
}

func (m *model) reloadBoard() {
	lists := m.kanban.(*kanban.Model).GetLists()
	storage.WriteToFile(lists[global.Todo].Items(), lists[global.InProgress].Items(), lists[global.Done].Items(), m.projects.(*project.Model).Projects[m.projects.(*project.Model).PrevSelected])
	m.kanban.(*kanban.Model).ReloadLists(m.projects.(*project.Model).Projects[m.projects.(*project.Model).Selected])
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, m.projects.View(), lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.kanban.View(),
		m.tOverview.View()))
}

// TODO save kanban on project change
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmdKanban tea.Cmd
	var cmdOverview tea.Cmd
	var cmdProjects tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
	case item.ProjectMsg:
		if msg.Type == global.Update {
			m.reloadBoard()
		}
	}

	m.kanban, cmdKanban = m.kanban.Update(msg)
	if cmdKanban != nil {
		return m, cmdKanban
	}
	m.projects, cmdProjects = m.projects.Update(msg)
	if cmdProjects != nil {
		return m, cmdProjects
	}
	m.tOverview, cmdOverview = m.tOverview.Update(msg)
	if cmdOverview != nil {
		return m, cmdOverview
	}
	return m, nil
}

// TODO load list of projects, check if project file exists
func main() {
	storage.Init()
	storage.CheckProjectsExist()
	projectModel := project.New()
	projectModel.Projects = storage.ReadProjects()
	m := &model{projects: projectModel, kanban: kanban.New(projectModel.Projects[0]), tOverview: taskoverview.New()}

	p = tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	lists := m.kanban.(*kanban.Model).GetLists()
	defer storage.WriteToFile(lists[global.Todo].Items(), lists[global.InProgress].Items(), lists[global.Done].Items(), projectModel.Projects[projectModel.Selected])
	defer storage.WriteProjectsToFile(projectModel.Projects)
}
