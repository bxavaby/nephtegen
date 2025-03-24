package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	PrimaryColor    = "#D4E1DA" // Whitened cyan
	SecondaryColor  = "#915E8C" // Pinkish-purple
	AccentColor     = "#6DD6CA" // Neon celestial blue
	OrnamentColor   = "#83CE84" // Soft neon green
	BackgroundColor = "#282A36" // Dark background
	padding         = 2
)

var (
	PrimaryStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(PrimaryColor))
	SecondaryStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(SecondaryColor))
	AccentStyle     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(AccentColor))
	OrnamentStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(OrnamentColor))
	BackgroundStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(BackgroundColor))

	docStyle   = lipgloss.NewStyle().Margin(1, 2)
	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#915E8C"))
	helpStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#915E8C")).Padding(1, 2)
)

func PrintLogo() {
	logo := AccentStyle.Render(`
╔═══════════════════════════════════════════════╗
║                                               ║
║  _  _ ____ ___  _  _ ___ ____ ____ ____ _  _  ║
║  |\ | |___ |__] |__|  |  |___ | __ |___ |\ |  ║
║  | \| |___ |    |  |  |  |___ |__] |___ | \|  ║
║                                               ║
╚═══════════════════════════════════════════════╝
    `)
	subtitle := PrimaryStyle.Render("The Nebet Setren Engine\n")

	fmt.Println(logo)
	fmt.Println(subtitle)
}

// spinner
func RunSpinner(label string, duration time.Duration) {
	s := spinner.New(spinner.WithSpinner(spinner.Meter))
	ticker := time.NewTicker(s.Spinner.FPS)
	defer ticker.Stop()

	done := make(chan struct{})

	go func() {
		frame := 0
		frames := s.Spinner.Frames
		for {
			select {
			case <-ticker.C:
				fmt.Printf("\r%s %s", OrnamentStyle.Render(frames[frame]), PrimaryStyle.Render(label))
				frame = (frame + 1) % len(frames)
			case <-done:
				return
			}
		}
	}()

	time.Sleep(duration)
	close(done)

	Wiper()
}

// progress bar
type progressModel struct {
	progress progress.Model
	percent  float64
	done     bool
}

func (m progressModel) Init() tea.Cmd {
	return tickCmd()
}

func (m progressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case tickMsg:
		m.percent += 1.0 / 5
		if m.percent >= 1.0 {
			m.done = true
			return m, tea.Quit
		}
		return m, tea.Batch(tickCmd(), m.progress.IncrPercent(1.0/5))

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}

func (m progressModel) View() string {
	if m.done {
		return ""
	}
	pad := strings.Repeat(" ", padding)
	return "\n" + pad + m.progress.View() + "\n"
}

func ProgressBar(duration int) {
	p := tea.NewProgram(progressModel{
		progress: progress.New(
			progress.WithGradient(PrimaryColor, AccentColor),
		),
		percent: 0,
	})

	if err := p.Start(); err != nil {
		fmt.Println("Error running progress bar:", err)
	}

	Wiper()
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// style formatting
func formatStyles(styles []string) string {
	return strings.Join(styles, ", ")
}

// log
func logCreation(creation string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting user home directory: %v\n", err)
		return
	}

	fileName := "ignis_ex_machina.log"

	err = os.MkdirAll(filepath.Dir(fileName), 0755)
	if err != nil {
		fmt.Printf("Error creating log directory: %v\n", err)
		return
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(fmt.Sprintf("%s\n", creation))
	if err != nil {
		fmt.Printf("Error writing to log file: %v\n", err)
		return
	}

	writer.Flush()
}

func logHaiku(result string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting user home directory: %v\n", err)
		return
	}

	fileName := "haiku.log"

	err = os.MkdirAll(filepath.Dir(fileName), 0755)
	if err != nil {
		fmt.Printf("Error creating log directory: %v\n", err)
		return
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer file.Close()

	lines := strings.Split(result, "\n")
	if len(lines) >= 3 {
		haiku := strings.Join(lines[:3], "\n") + "\n\n"

		writer := bufio.NewWriter(file)
		_, err = writer.WriteString(haiku)
		if err != nil {
			fmt.Printf("Error writing to log file: %v\n", err)
			return
		}
		writer.Flush()
	}
}

// random letter fetcher
func fetchRandomLetter() string {
	filePath := filepath.Join("letters.json")

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening letters file: %v\n", err)
		return ""
	}
	defer file.Close()

	var data struct {
		Letters []string `json:"letters"`
	}

	if err := json.NewDecoder(file).Decode(&data); err != nil {
		fmt.Printf("Error decoding letters file: %v\n", err)
		return ""
	}

	rand.Seed(time.Now().UnixNano())
	return data.Letters[rand.Intn(len(data.Letters))]
}

// screen wiper
func Wiper() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout

	cmd.Run()
}
