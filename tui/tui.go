package cli

import (
	"block-game-go/board"
	"block-game-go/piece"
	"block-game-go/util"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var styles = struct {
	border lipgloss.Style
}{
	border: lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true, true, true, true),
}

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
			err := m.board.PlacePiece(m.pieces[m.pieceI], m.boardPos)
			if err != nil {
				return m, nil
			}

			m.pieces = util.Remove(m.pieces, m.pieceI)
			if len(m.pieces) == 0 {
				m.pieces = []piece.Piece{piece.RandomPiece(), piece.RandomPiece(), piece.RandomPiece()}
			}
		}
	}

	// handle out of bounds of pieces
	if m.pieceI >= len(m.pieces) {
		m.pieceI = len(m.pieces) - 1
	}

	// handle out of bounds of board
	if m.boardPos.ColI < 0 {
		m.boardPos.ColI = 9 - m.pieces[m.pieceI].Width()
	}
	if m.boardPos.ColI+m.pieces[m.pieceI].Width() > 9 {
		m.boardPos.ColI = 0
	}
	if m.boardPos.RowI < 0 {
		m.boardPos.RowI = 9 - m.pieces[m.pieceI].Height()
	}
	if m.boardPos.RowI+m.pieces[m.pieceI].Height() > 9 {
		m.boardPos.RowI = 0
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
			piece = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF")).Render(choice.ToString())
		} else {
			piece = lipgloss.NewStyle().Render(choice.ToString())
		}
		pieces = append(pieces, piece)
	}

	return styles.border.Width(11).Height(18).AlignHorizontal(lipgloss.Center).Render(lipgloss.JoinVertical(lipgloss.Center, pieces...))
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

			var cellStr string
			switch {
			case cell:
				cellStr = "▓▓"
			case isLightCell:
				cellStr = "░░"
			default:
				cellStr = "▒▒"
			}

			if selectedPieceCell {
				cellStr = lipgloss.NewStyle().Background(lipgloss.Color("#FF00FF")).Render(cellStr)
			}

			str += cellStr
		}

		if rowI < len(m.board.Grid)-1 {
			str += "\n"
		}
	}

	return styles.border.MarginRight(1).Render(str)
}

func isLight(cell board.Cell) bool {
	isInOddRow := cell.RowI/3%2 == 1
	isInOddCol := cell.ColI/3%2 == 1
	return isInOddRow != isInOddCol
}

func (m model) scoreUI() string {
	return fmt.Sprintf("Score: %d\nStreak: %d", m.board.Score, m.board.Streak)
}

func Play() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
