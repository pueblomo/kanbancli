package form

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	item "pueblomo/kanbancli/model"
)

type model struct {
	titel       textinput.Model
	tag         textinput.Model
	description textarea.Model
	oldModel    tea.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

// TODO esc disable focus
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+x":
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

func (m model) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, m.titel.View(), m.tag.View(), m.description.View())
}

func (m model) createTask() tea.Msg {
	task := item.New(m.titel.Value(), m.tag.Value(), m.description.Value())
	return item.NewTaskMsg(true, task)
}

func New(oldModel tea.Model) *model {
	f := &model{}
	f.titel = textinput.New()
	f.titel.Placeholder = "Titel"
	f.titel.Focus()
	f.tag = textinput.New()
	f.tag.Placeholder = "Tag"
	f.description = textarea.New()
	f.description.Placeholder = "Description"
	f.oldModel = oldModel
	return f
}
