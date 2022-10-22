package kanban

import (
	"pueblomo/kanbancli/form"
	"pueblomo/kanbancli/global"
	item "pueblomo/kanbancli/model"
	"pueblomo/kanbancli/storage"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	columnStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.HiddenBorder())
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
)

type Model struct {
	focused global.Status
	lists   []list.Model
}

func (m *Model) initLists() {
	m.lists = storage.ReadFile()
}

func (m Model) GetLists() []list.Model {
	return m.lists
}

func (m Model) GetSize() (height int, width int) {
	return m.lists[global.Todo].Height(), m.lists[global.Todo].Width()
}

func New() *Model {
	m := &Model{focused: global.Todo, lists: []list.Model{}}
	m.initLists()
	return m
}

func (m *Model) next() tea.Msg {
	if m.focused == global.Done {
		m.focused = global.Todo
	} else {
		m.focused++
	}
	return nil
}

func (m *Model) prev() tea.Msg {
	if m.focused == global.Todo {
		m.focused = global.Done
	} else {
		m.focused--
	}
	return nil
}

func (m *Model) moveToNext() tea.Msg {
	selectedItem := m.lists[m.focused].SelectedItem()
	if selectedItem != nil {
		selectedTask := selectedItem.(item.Task)
		m.lists[selectedTask.Status].RemoveItem(m.lists[m.focused].Index())
		selectedTask.Next()
		m.lists[selectedTask.Status].InsertItem(len(m.lists[selectedTask.Status].Items()), list.Item(selectedTask))
	}

	return nil
}

func (m *Model) showTask() tea.Msg {
	return item.NewTaskMsg(false, m.lists[m.focused].SelectedItem().(item.Task))
}

func (m *Model) deleteTask() tea.Msg {
	selectedItem := m.lists[m.focused].SelectedItem()
	if selectedItem != nil {
		selectedTask := selectedItem.(item.Task)
		m.lists[selectedTask.Status].RemoveItem(m.lists[m.focused].Index())
	}
	return nil
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "left":
			return m, m.prev
		case "right":
			return m, m.next
		case "enter":
			return m, m.moveToNext
		case "s":
			return m, m.showTask
		case "r":
			return m, m.deleteTask
		case "n":
			return form.New(m, m.lists[global.Todo].Height(), m.lists[global.Todo].Width()).Update(nil)
		}
	case tea.WindowSizeMsg:
		height := msg.Height - global.Divisor
		width := msg.Width/global.Divisor - global.Divisor
		columnStyle.Width(width)
		focusedStyle.Width(width)
		columnStyle.Height(height)
		focusedStyle.Height(height)
		m.lists[global.Todo].SetSize(width, height)
		m.lists[global.InProgress].SetSize(width, height)
		m.lists[global.Done].SetSize(width, height)
		return m, nil
	case item.TaskMsg:
		if msg.Create {
			return m, m.lists[global.Todo].InsertItem(len(m.lists[global.Todo].Items()), msg.Task)
		}
	}

	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	if len(m.lists) <= 0 {
		return "loading..."
	}
	todoView := m.lists[global.Todo].View()
	inProgView := m.lists[global.InProgress].View()
	doneView := m.lists[global.Done].View()
	switch m.focused {
	case global.InProgress:
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			columnStyle.Render(todoView),
			focusedStyle.Render(inProgView),
			columnStyle.Render(doneView),
		)
	case global.Done:
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			columnStyle.Render(todoView),
			columnStyle.Render(inProgView),
			focusedStyle.Render(doneView),
		)
	default:
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			focusedStyle.Render(todoView),
			columnStyle.Render(inProgView),
			columnStyle.Render(doneView),
		)
	}
}
