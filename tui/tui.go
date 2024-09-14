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
	selectedPiece := m.pieces[m.pieceI]

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
			if m.boardPos.ColI < 0 {
				m.boardPos.ColI = 9 - selectedPiece.Width()
			}
		case "right":
			m.boardPos.ColI++
			if m.boardPos.ColI+selectedPiece.Width() > 9 {
				m.boardPos.ColI = 0
			}
		case "up":
			m.boardPos.RowI--
			if m.boardPos.RowI < 0 {
				m.boardPos.RowI = 9 - selectedPiece.Height()
			}
		case "down":
			m.boardPos.RowI++
			if m.boardPos.RowI+selectedPiece.Height() > 9 {
				m.boardPos.RowI = 0
			}

		// piece placement
		case "enter", " ":
			m.board.PlacePattern(selectedPiece.Grid, m.boardPos)
			m.board.Evaluate()
		}
	}

	return m, nil
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, m.piecesUI(), m.boardUI(), m.scoreUI())
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
	str := ""

	adjSelectedPiece := [9][9]bool{}
	for rowI := range selectedPiece.Grid {
		for colI := range selectedPiece.Grid[rowI] {
			if selectedPiece.Grid[rowI][colI] {
				adjSelectedPiece[rowI+m.boardPos.RowI][colI+m.boardPos.ColI] = selectedPiece.Grid[rowI][colI]
			}
		}
	}

	for rowI := range m.board.Grid {
		for colI := range m.board.Grid[rowI] {
			cell := m.board.Grid[rowI][colI]
			selectedPieceCell := adjSelectedPiece[rowI][colI]

			isLightCell := isLight(board.Cell{RowI: rowI, ColI: colI})

			cellStr := ""
			if cell {
				cellStr += "▓▓"
			} else if isLightCell {
				cellStr += "░░"
			} else {
				cellStr += "▒▒"
			}

			if selectedPieceCell {
				cellStr = lipgloss.NewStyle().Background(lipgloss.Color("#FF00FF")).Render(cellStr)
			}

			str += cellStr
		}
		str += "\n"
	}

	return str
}

func isLight(cell board.Cell) bool {
	isInOddRow := cell.RowI/3%2 == 1
	isInOddCol := cell.ColI/3%2 == 1
	return isInOddRow != isInOddCol
}

func (m model) scoreUI() string {
	return fmt.Sprintf("Score: %d\nMultiplier: %d", m.board.Score, m.board.Multiplier)
}

func Play() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
