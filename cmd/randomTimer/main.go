package main

// A simple example that shows how to render an animated progress bar. In this
// example we bump the progress by 25% every two seconds, animating our
// progress bar to its new target state.
//
// It's also possible to render a progress bar in a more static fashion without
// transitions. For details on that approach see the progress-static example.

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/keyzox71/randomTimer/utils"
)


const (
	padding  = 2
	maxWidth = 80
)

func main() {
	if (len(os.Args) != 2) {
		fmt.Println("Specify the time in argument")
		os.Exit(1)
	}
	time, err := strconv.ParseFloat(os.Args[1], 32)
	if (err != nil){
		log.Fatal(err)
		os.Exit(1)
	}
		
	theme := utils.GetThemeStyle()
	m := model{
		progress: progress.New(progress.WithGradient(theme.Accent1.Hex, theme.Accent2.Hex)),
		theme: theme,
		help: lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Text.Hex)),
		time: time * 60, 
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}

type model struct {
	progress progress.Model
	theme utils.Theme
	help lipgloss.Style
	time float64
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case time.Time:
		if m.progress.Percent() == 1.0 {
			return m, tea.Quit
		}

		percent := 1.0 / m.time
			cmd := m.progress.IncrPercent(percent)
		return m, tea.Batch(tickCmd(), cmd)

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

func (m model) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + m.progress.View() + "\n\n" +
		pad + m.help.Render("Press any key to quit")
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return time.Time(t)
	})
}
