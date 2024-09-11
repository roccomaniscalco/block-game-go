package cli

import (
	"block-game-go/piece"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

type model struct {
	choices       []piece.Piece
	cursor        int
	selectedPiece int
}

func initialModel() model {
	return model{
		choices: []piece.Piece{piece.RandomPiece(), piece.RandomPiece(), piece.RandomPiece()},
		cursor:  0,
		selectedPiece: -1,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left", "j":
			if m.cursor > 0 {
				m.cursor--
			}
		case "right", "k":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.selectedPiece = m.cursor
		}
	}
	return m, nil
}

func (m model) View() string {
	pieces := []string{}

	for i, choice := range m.choices {
		piece := ""
		if i == m.cursor {
			piece = lipgloss.NewStyle().MarginRight(2).Foreground(lipgloss.Color("#FF00FF")).Render(choice.ToString())
		} else {
			piece = lipgloss.NewStyle().MarginRight(2).Render(choice.ToString())
		}
		if i == m.selectedPiece {
			piece = lipgloss.NewStyle().MarginRight(2).Foreground(lipgloss.Color("#FFF")).Render(choice.ToString())
		}
		pieces = append(pieces, piece)
	}

	return lipgloss.JoinHorizontal(0.5, pieces...)
}

func Play() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
