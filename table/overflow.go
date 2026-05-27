package table

import "github.com/charmbracelet/lipgloss"

const columnKeyOverflowRight = "___overflow_r___"
const columnKeyOverflowLeft = "___overflow_l__"

func genOverflowStyle(base lipgloss.Style, width int) lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

func genOverflowColumnRight(width int) Column { _ = "STUB: not implemented"; return *new(Column) }

func genOverflowColumnLeft(width int) Column { _ = "STUB: not implemented"; return *new(Column) }
