package ui

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type textmodel struct {
	textInput textinput.Model
	err       error
}

type model struct {
	choices []string
	cursor  int
}

type (
	errMsg error
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m textmodel) Init() tea.Cmd {
	return nil
}

func initialModel() model {
	return model{
		choices: []string{"Select Port", "Connect", "Host"},
		cursor:  0,
	}
}

func textModel() textmodel {
	ti := textinput.NewModel()
	ti.Focus()
	ti.CharLimit = 20
	ti.Width = 20

	return textmodel{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			if m.cursor == 0 {
				return textModel(), nil
			} else if m.cursor == 1 {
				return textModel(), nil
			} else {
				return messageScreenModel(), nil
			}

		}
	}
	return m, nil
}

func (m textmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:

			return messageScreenModel(), nil
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, tea.Quit
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd

}

func (m model) View() string {
	s := "Wisp\n\n"

	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}

func (m textmodel) View() string {
	return fmt.Sprintf(
		"Enter connection key\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func WispInit() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
