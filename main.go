package main

import (
	"fmt"
	"os"
	"pueblomo/kanbancli/global"
	"pueblomo/kanbancli/kanban"
	"pueblomo/kanbancli/storage"
	taskoverview "pueblomo/kanbancli/taskOverview"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	kanban    tea.Model
	tOverview tea.Model
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.kanban.View(),
		m.tOverview.View())
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmdKanban tea.Cmd
	var cmdOverview tea.Cmd
	m.kanban, cmdKanban = m.kanban.Update(msg)
	m.tOverview, cmdOverview = m.tOverview.Update(msg)
	if cmdKanban != nil {
		return m, cmdKanban
	}
	if cmdOverview != nil {
		return m, cmdOverview
	}
	return m, nil
}

func main() {
	storage.CheckFileExists()
	m := &model{kanban: kanban.New(), tOverview: taskoverview.New()}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	lists := m.kanban.(*kanban.Model).GetLists()
	defer storage.WriteToFile(lists[global.Todo].Items(), lists[global.InProgress].Items(), lists[global.Done].Items())
}
