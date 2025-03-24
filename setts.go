package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionModel struct {
	list     list.Model
	help     string
	selected string
	quitting bool
}

var sessionNames = []string{"tri", "one", "dual", "haiku", "motto", "symb", "port", "asphrase", "allit", "acronym"}

type settskeyMap struct {
	Flush    key.Binding
	FlushAll key.Binding
	Back     key.Binding
	Quit     key.Binding
}

var settskeys = settskeyMap{
	Flush:    key.NewBinding(key.WithKeys("f"), key.WithHelp("f", "flush selected session")),
	FlushAll: key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "flush all sessions")),
	Back:     key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "back")),
	Quit:     key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
}

type sessionItem struct {
	name string
}

func (i sessionItem) Title() string       { return i.name }
func (i sessionItem) Description() string { return "Active session." }
func (i sessionItem) FilterValue() string { return i.name }

func newSessionModel() sessionModel {
	items := fetchFabricSessions()

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = titleStyle.Render("Session Flusher")
	l.SetShowHelp(false)

	return sessionModel{
		list: l,
		help: "Use 'f' to flush, 'a' to flush all, 'esc' to return.",
	}
}

func fetchFabricSessions() []list.Item {
	cmd := exec.Command("zsh", "-c", "fabric -X")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error fetching sessions:", err)
		return nil
	}

	lines := strings.Split(string(output), "\n")
	var items []list.Item
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if contains(sessionNames, trimmed) {
			items = append(items, sessionItem{name: trimmed})
		}
	}
	return items
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func flushSession(name string) {
	cmd := exec.Command("zsh", "-c", fmt.Sprintf("fabric -W %s", name))
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error flushing session %s: %v\n", name, err)
	} else {
		fmt.Printf("Flushed %s\n", name)
	}
}

func flushAllSessions() {
	for _, session := range sessionNames {
		flushSession(session)
	}
}

func (m sessionModel) Init() tea.Cmd {
	return nil
}

func (m sessionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, settskeys.Quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, settskeys.Back):
			return newMainMenu(), nil
		case key.Matches(msg, settskeys.Flush):
			if selectedItem, ok := m.list.SelectedItem().(sessionItem); ok {
				flushSession(selectedItem.name)
				return newSessionModel(), nil
			}
		case key.Matches(msg, settskeys.FlushAll):
			flushAllSessions()
			return newSessionModel(), nil
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m sessionModel) View() string {
	if m.quitting {
		return OrnamentStyle.Render("Returning...")
	}

	mainView := docStyle.Render(m.list.View() + "\n" + m.help)

	return mainView
}
