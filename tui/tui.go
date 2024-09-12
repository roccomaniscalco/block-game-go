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
	board    board.Board
	boardPos board.Cell
	pieces   []piece.Piece
	pieceI   int
}

func initialModel() model {
	return model{
		board:    board.NewBoard(),
		boardPos: board.Cell{RowI: 0, ColI: 0},
		pieces:   []piece.Piece{piece.RandomPiece(), piece.RandomPiece(), piece.RandomPiece()},
		pieceI:   0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// quit
		case "ctrl+c", "q":
			return m, tea.Quit

		// piece selection
		case "1":
			m.pieceI = 0
		case "2":
			m.pieceI = 1
		case "3":
			m.pieceI = 2

		// piece movement
		case "left":
			m.boardPos.ColI--
		case "right":
			m.boardPos.ColI++
		case "up":
			m.boardPos.RowI--
		case "down":
			m.boardPos.RowI++

		// piece placement
		case "enter", " ":
			m.board.PlacePattern(m.pieces[m.pieceI].Grid, m.boardPos)
			m.board.Evaluate()
		}
	}

	return m, nil
}

func (m model) piecesUI() string {
	pieces := []string{}

	for i, choice := range m.pieces {
		piece := ""
		if i == m.pieceI {
			piece = lipgloss.NewStyle().MarginBottom(1).Foreground(lipgloss.Color("#FF00FF")).Render(choice.ToString())
		} else {
			piece = lipgloss.NewStyle().MarginBottom(1).Render(choice.ToString())
		}
		pieces = append(pieces, piece)
	}

	return lipgloss.NewStyle().MarginRight(2).Render(lipgloss.JoinVertical(lipgloss.Center, pieces...))
}

func (m model) boardUI() string {
	selectedPiece := m.pieces[m.pieceI]

	for rowI := range selectedPiece.Grid {
		for colI := range selectedPiece.Grid[rowI] {
			if selectedPiece.Grid[rowI][colI] {
				m.board.Grid[m.boardPos.RowI+rowI][m.boardPos.ColI+colI] =
					selectedPiece.Grid[rowI][colI]
			}
		}
	}

	return m.board.ToString()
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, m.piecesUI(), m.boardUI())
}

func Play() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
