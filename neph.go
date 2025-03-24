package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type mainkeyMap struct {
	Select key.Binding
	Back   key.Binding
	Quit   key.Binding
}

func (k mainkeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Select, k.Back, k.Quit},
	}
}

func (k mainkeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Select, k.Back, k.Quit}
}

var mainkeys = mainkeyMap{
	Select: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select")),
	Back:   key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "back")),
	Quit:   key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
}

type menuItem struct {
	title, desc string
	action      func() tea.Msg
}

func (i menuItem) Title() string       { return i.title }
func (i menuItem) Description() string { return i.desc }
func (i menuItem) FilterValue() string { return i.title }

type model struct {
	list     list.Model
	help     help.Model
	quitting bool
	parent   tea.Model
}

func newMainMenu() model {
	items := []list.Item{
		menuItem{
			title:  "Run the Selection",
			desc:   "Execute a pre-configured selection flow",
			action: func() tea.Msg { return runSelectionFlow() },
		},
		menuItem{
			title:  "Run Randomized Selection",
			desc:   "Execute a randomized selection flow",
			action: func() tea.Msg { return runRandomizedSelection() },
		},
		menuItem{
			title:  "Architect",
			desc:   "Enter the architect mode to create new configurations",
			action: func() tea.Msg { return architectMode() },
		},
		menuItem{
			title:  "Settings",
			desc:   "Change application settings",
			action: func() tea.Msg { return settingsMenu() },
		},
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = titleStyle.Render("Ren-Khem")
	l.SetShowHelp(false)

	return model{
		list: l,
		help: help.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, mainkeys.Quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, mainkeys.Back) && m.parent != nil:
			return m.parent, nil
		case key.Matches(msg, mainkeys.Select):
			if item, ok := m.list.SelectedItem().(menuItem); ok {
				return m, tea.Batch(func() tea.Msg { return item.action() })
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return OrnamentStyle.Render("Fare thee well!")
	}

	mainView := docStyle.Render(m.list.View() + "\n" + m.help.View(mainkeys))

	return mainView
}

func runSelectionFlow() tea.Msg {
	Wiper()
	ProgressBar(2)
	RunSelectionFlow()

	if !RestartOption() {
		return nil
	}
	return nil
}

func runRandomizedSelection() tea.Msg {
	Wiper()
	ProgressBar(5)
	fmt.Println("Running the randomized selection flow...")
	return nil
}

func architectMode() tea.Msg {
	Wiper()
	ProgressBar(5)
	fmt.Println("Evoking architect...")
	return nil
}

func settingsMenu() tea.Msg {
	Wiper()
	ProgressBar(2)

	p := tea.NewProgram(newSessionModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error running Session Flusher: %v\n", err)
	}

	return nil
}

func RestartOption() bool {
	var response string
	fmt.Print(OrnamentStyle.Render("\nShall I conceive yet another Ren? (y/n): "))
	fmt.Scanln(&response)
	return response == "y" || response == "Y"
}

// func main() {
//	Wiper()
//	PrintLogo()
//	p := tea.NewProgram(newMainMenu())
//	if err := p.Start(); err != nil {
//		fmt.Printf("Error starting program: %v\n", err)
//	}
//}

func main() {
	for {
		Wiper()
		ProgressBar(3)
		PrintLogo()
		RunSelectionFlow()

		if !RestartOption() {
			Wiper()
			fmt.Println(OrnamentStyle.Render("Fare thee well!"))
			break
		}
	}
	return
}
