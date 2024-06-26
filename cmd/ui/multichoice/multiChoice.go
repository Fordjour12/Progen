package multichoice

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#EFEFEF")).Background(lipgloss.Color("#333333")).Padding(1, 2)
var focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#333333")).Background(lipgloss.Color("#FFCC02")).Padding(1, 2)

type Selection struct {
	Choice string
}

func (s *Selection) Update(value string) {
	s.Choice = value
}

type model struct {
	cursor   int
	header   string
	choice   *Selection
	selected map[int]struct{}
	choices  []string
}

func InitialChoiceModel(header string, choices []string, selection *Selection) model {
	return model{
		header:   titleStyle.Render(header),
		choices:  choices,
		choice:   selection,
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//TODO:(Must choice only one and not more than one fixme)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--

			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if len(m.selected) == 1 {
				m.selected = make(map[int]struct{})
			}
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "y":
			if len(m.selected) == 1 {
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := m.header + "\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if i == m.cursor {
			// cursor = ">"
			cursor = focusedStyle.Render("▸")
		}
		selected := " "
		if _, ok := m.selected[i]; ok {
			selected = focusedStyle.Render("x")
		}
		s += cursor + " " + selected + " " + choice + "\n"
		// s += fmt.Sprintf("%s %s %s\n", cursor, selected, choice)
	}

	// s += "\n\nPress 'y' to confirm selection\n"
	s += fmt.Sprintf("\n\n%s\n", focusedStyle.Render("Press 'y' to confirm selection"))
	return s

}
