package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type messageScreen struct {
	textInput textinput.Model
	Msgs      []string
}

func messageScreenModel() messageScreen {
	ti := textinput.NewModel()
	ti.Focus()
	ti.CharLimit = 512
	ti.Width = 50

	return messageScreen{
		textInput: ti,
		Msgs:      []string{"", ""},
	}
}

func (m messageScreen) Init() tea.Cmd {
	return nil
}

func (m messageScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			m.Msgs[0] = m.Msgs[1]
			m.Msgs[1] = m.textInput.Value()
			m.textInput.SetValue("")
			return m, nil
		}

	// We handle errors just like any other message
	case errMsg:
		m.Msgs[0] = "\n"
		m.Msgs[1] = fmt.Sprintf("Error: %v", msg)
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m messageScreen) View() string {
	return m.Msgs[0] + "\n" + m.Msgs[1] + "\n" + m.textInput.View()

}
