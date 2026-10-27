package cjtrace

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	incorrectColor = lipgloss.Color("203")
	incorrectStyle = lipgloss.NewStyle().Foreground(incorrectColor)
)

type Model struct {
	TextInput textinput.Model
	Debug     bool
	Log       string
}

func New() Model {
	ti := textinput.New()
	ti.Focus()
	return Model{TextInput: ti}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		k := msg.Key()

		n := len([]rune(m.TextInput.Value()))
		if len(msg.Text) == 0 {
			m.TextInput, cmd = m.TextInput.Update(msg)
		} else if c, ok := cjmap[strings.ToLower(msg.Text)]; ok && n < 5 {
			msg.Text = c
			m.TextInput, cmd = m.TextInput.Update(msg)
		}

		n = len([]rune(m.TextInput.Value()))
		if n == 5 {
			styles := m.TextInput.Styles()
			styles.Focused.Text = incorrectStyle
			styles.Cursor.Color = incorrectColor
			m.TextInput.SetStyles(styles)
		} else {
			m.TextInput.SetStyles(textinput.DefaultDarkStyles())
		}

		m.Log = fmt.Sprintf("%s pressed", k)
	}
	return m, cmd
}

func (m Model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.TextInput.View(),
		m.footer(),
	)
}

func (m Model) footer() string {
	if !m.Debug {
		return ""
	}
	return fmt.Sprintf("[LOG] %s", m.Log)
}
