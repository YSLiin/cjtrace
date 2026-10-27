package main

import (
	"log"

	"github.com/ysliin/cjtrace/cjtrace"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Model struct {
	typer cjtrace.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "enter":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.typer, cmd = m.typer.Update(msg)
	return m, cmd
}

func (m Model) View() tea.View {
	return tea.NewView(
		lipgloss.JoinVertical(
			lipgloss.Left,
			m.typer.View(),
			"Press ctrl+c or Enter to exit...",
		),
	)
}

func main() {
	cjtyper := cjtrace.New()
	if _, err := tea.NewProgram(Model{typer: cjtyper}).Run(); err != nil {
		log.Fatalln(err)
	}
}
