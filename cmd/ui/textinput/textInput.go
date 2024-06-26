package textinput

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#EFEFEF")).Background(lipgloss.Color("#333333")).Padding(1, 2)

type Output struct {
	Output string
}

func (o *Output) update(val string) {
	o.Output = val
}

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	err       error
	header    string
	output    *Output
}

func InitialModel(header string, output *Output) model {
	ti := textinput.New()
	ti.Placeholder = "Type here…"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
		header:    titleStyle.Render(header),
		output:    output,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			// TODO: if error come here and check lenght
			if m.textInput.Value() != "" {
				m.output.update(m.textInput.Value())
				return m, tea.Quit
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s\n",
		m.header,
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

