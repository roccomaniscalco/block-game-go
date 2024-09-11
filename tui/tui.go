package cli

import (
	"block-game-go/board"
	"block-game-go/piece"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	pieceChoices   []piece.Piece
	activePieceI   int
	activePiecePos board.Cell
}

func initialModel() model {
	return model{
		pieceChoices:   []piece.Piece{piece.RandomPiece(), piece.RandomPiece(), piece.RandomPiece()},
		activePieceI:   0,
		activePiecePos: board.Cell{RowI: 0, ColI: 0},
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

		// piece selection
		case "1":
			m.activePieceI = 0
		case "2":
			m.activePieceI = 1
		case "3":
			m.activePieceI = 2

		// piece movement
		case "left":
			m.activePiecePos.ColI--
		case "right":
			m.activePiecePos.ColI++
		case "up":
			m.activePiecePos.RowI--
		case "down":
			m.activePiecePos.RowI++

		}
	}
	return m, nil
}

func (m model) View() string {
	pieces := []string{}

	for i, choice := range m.pieceChoices {
		piece := ""
		if i == m.activePieceI {
			piece = lipgloss.NewStyle().MarginRight(2).Foreground(lipgloss.Color("#FF00FF")).Render(choice.ToString())
		} else {
			piece = lipgloss.NewStyle().MarginRight(2).Render(choice.ToString())
		}
		pieces = append(pieces, piece)
	}

	return lipgloss.JoinHorizontal(lipgloss.Center, pieces...)
}

func Play() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
