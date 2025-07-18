package tui

import (
	"block-game-go/board"
	"block-game-go/piece"
	"fmt"
	"os"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var styles = struct {
	border lipgloss.Style
}{
	border: lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true),
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
		case "w", "up":
			m.boardPos.RowI--
		case "a", "left":
			m.boardPos.ColI--
		case "s", "down":
			m.boardPos.RowI++
		case "d", "right":
			m.boardPos.ColI++

		// piece placement
		case "enter", " ":
			if err := m.board.PlacePiece(m.pieces[m.pieceI], m.boardPos); err != nil {
				return m, nil
			}

			m.pieces = slices.Delete(m.pieces, m.pieceI, m.pieceI+1)
			if len(m.pieces) == 0 {
				m.pieces = []piece.Piece{piece.RandomPiece(), piece.RandomPiece(), piece.RandomPiece()}
			}
		}
	}

	m = m.ensurePieceIndexWithinBounds()
	m = m.ensureBoardPositionWithinBounds()

	if m.board.IsGameOver(m.pieces) {
		return m, tea.Quit
	}

	return m, nil
}

func (m model) ensurePieceIndexWithinBounds() model {
	if m.pieceI >= len(m.pieces) {
		m.pieceI = len(m.pieces) - 1
	}
	return m
}

func (m model) ensureBoardPositionWithinBounds() model {
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
	return m
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, m.piecesUI(),
		lipgloss.JoinVertical(lipgloss.Top, m.scoreUI(), m.boardUI()))
}

func (m model) piecesUI() string {
	pieceStrs := []string{}

	for i, piece := range m.pieces {
		pieceStr := ""
		if i == m.pieceI {
			pieceStr = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF")).Render(piece.ToString())
		} else {
			pieceStr = lipgloss.NewStyle().Render(piece.ToString())
		}
		pieceStrs = append(pieceStrs, pieceStr)
	}

	return styles.border.
		Width(11).
		Height(17).
		AlignHorizontal(lipgloss.Center).
		MarginRight(1).
		Render(lipgloss.JoinVertical(lipgloss.Center, pieceStrs...))
}

func (m model) boardUI() string {
	selectedPiece := m.pieces[m.pieceI]
	boardStr := ""

	// selected piece adjusted to its position on the board
	selectedPieceOverlay := [9][9]bool{}
	for rowI := range selectedPiece.Grid {
		for colI := range selectedPiece.Grid[rowI] {
			if selectedPiece.Grid[rowI][colI] {
				selectedPieceOverlay[rowI+m.boardPos.RowI][colI+m.boardPos.ColI] = selectedPiece.Grid[rowI][colI]
			}
		}
	}

	for rowI := range m.board.Grid {
		for colI := range m.board.Grid[rowI] {
			isCellFilled := m.board.Grid[rowI][colI]
			isCellInOddSquare := (rowI/3%2 == 1) != (colI/3%2 == 1)
			isCellSelected := selectedPieceOverlay[rowI][colI]

			var cellStr string
			switch {
			case isCellFilled:
				cellStr = "▓▓"
			case isCellInOddSquare:
				cellStr = "░░"
			default:
				cellStr = "▒▒"
			}

			if isCellSelected {
				cellStr = lipgloss.NewStyle().Background(lipgloss.Color("#FF00FF")).Render(cellStr)
			}

			boardStr += cellStr
		}

		if rowI < len(m.board.Grid)-1 {
			boardStr += "\n"
		}
	}

	return styles.border.MarginRight(1).Render(boardStr)
}

func (m model) scoreUI() string {
	return styles.border.Width(18).Render(fmt.Sprintf("Score: %d\nStreak: %d", m.board.Score, m.board.Streak))
}

func Play() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
