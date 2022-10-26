package form

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"pueblomo/kanbancli/global"
	item "pueblomo/kanbancli/model"
)

var (
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))

	elementStyle = lipgloss.NewStyle().Padding(1).Foreground(lipgloss.Color("7"))
)

type model struct {
	titel       textinput.Model
	tag         textinput.Model
	description textarea.Model
	oldModel    tea.Model
	editMode    bool
	status      global.Status
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		focusedStyle.Width((msg.Width/global.Divisor)*3 - global.Divisor)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+x":
			if m.editMode {
				return m.oldModel, m.updateTask
			}
			return m.oldModel, m.createTask
		case "tab":
			if m.titel.Focused() {
				m.titel.Blur()
				m.tag.Focus()
				return m, textinput.Blink
			} else if m.tag.Focused() {
				m.tag.Blur()
				m.description.Focus()
				return m, textarea.Blink
			} else {
				m.description.Blur()
				m.titel.Focus()
				return m, textinput.Blink
			}
		case "esc":
			return m.oldModel, nil
		}
	case item.Task:
		m.titel.SetValue(msg.Subject)
		m.description.SetValue(msg.Desc)
		m.tag.SetValue(msg.Tag)
		m.titel.Focus()
		m.tag.Blur()
		m.editMode = true
		m.status = msg.Status
		return m, textinput.Blink
	}

	if m.titel.Focused() {
		m.titel, cmd = m.titel.Update(msg)
		return m, cmd
	} else if m.tag.Focused() {
		m.tag, cmd = m.tag.Update(msg)
		return m, cmd
	} else {
		m.description, cmd = m.description.Update(msg)
		return m, cmd
	}
}

func (m *model) View() string {
	return focusedStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			elementStyle.Render(m.titel.View()),
			elementStyle.Render(m.tag.View()),
			elementStyle.Render(m.description.View()),
		),
	)
}

func (m model) createTask() tea.Msg {
	task := item.New(m.titel.Value(), m.tag.Value(), m.description.Value())
	return item.NewTaskMsg(global.Create, task)
}

func (m model) updateTask() tea.Msg {
	task := item.New(m.titel.Value(), m.tag.Value(), m.description.Value())
	task.Status = m.status
	return item.NewTaskMsg(global.Update, task)
}

func New(oldModel tea.Model, height, width int) *model {
	f := &model{}
	f.titel = textinput.New()
	f.titel.Placeholder = "Titel"
	f.titel.Focus()
	f.tag = textinput.New()
	f.tag.Placeholder = "Tag"
	f.description = textarea.New()
	f.description.Placeholder = "Description"
	f.oldModel = oldModel
	f.setHeight(height)
	f.setWidth(width)
	f.editMode = false
	f.status = global.Todo
	return f
}

func (m *model) setHeight(height int) {
	focusedStyle.Height(height)
	m.description.SetHeight(height - 10)
}

func (m *model) setWidth(width int) {
	focusedStyle.Width(width * 3)
	m.description.SetWidth(width*3 - 10)
}
