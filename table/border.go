package table

import "github.com/charmbracelet/lipgloss"

// Border defines the borders in and around the table.
type Border struct {
	Top         string
	Left        string
	Right       string
	Bottom      string
	TopRight    string
	TopLeft     string
	BottomRight string
	BottomLeft  string

	TopJunction    string
	LeftJunction   string
	RightJunction  string
	BottomJunction string

	InnerJunction string

	InnerDivider string

	// Styles for 2x2 tables and larger
	styleMultiTopLeft     lipgloss.Style
	styleMultiTop         lipgloss.Style
	styleMultiTopRight    lipgloss.Style
	styleMultiRight       lipgloss.Style
	styleMultiBottomRight lipgloss.Style
	styleMultiBottom      lipgloss.Style
	styleMultiBottomLeft  lipgloss.Style
	styleMultiLeft        lipgloss.Style
	styleMultiInner       lipgloss.Style

	// Styles for a single column table
	styleSingleColumnTop    lipgloss.Style
	styleSingleColumnInner  lipgloss.Style
	styleSingleColumnBottom lipgloss.Style

	// Styles for a single row table
	styleSingleRowLeft  lipgloss.Style
	styleSingleRowInner lipgloss.Style
	styleSingleRowRight lipgloss.Style

	// Style for a table with only one cell
	styleSingleCell lipgloss.Style

	// Style for the footer
	styleFooter lipgloss.Style
}

var (
	// https://www.w3.org/TR/xml-entity-names/025.html

	borderDefault = Border{
		Top:    "━",
		Left:   "┃",
		Right:  "┃",
		Bottom: "━",

		TopRight:    "┓",
		TopLeft:     "┏",
		BottomRight: "┛",
		BottomLeft:  "┗",

		TopJunction:    "┳",
		LeftJunction:   "┣",
		RightJunction:  "┫",
		BottomJunction: "┻",
		InnerJunction:  "╋",

		InnerDivider: "┃",
	}

	borderRounded = Border{
		Top:    "─",
		Left:   "│",
		Right:  "│",
		Bottom: "─",

		TopRight:    "╮",
		TopLeft:     "╭",
		BottomRight: "╯",
		BottomLeft:  "╰",

		TopJunction:    "┬",
		LeftJunction:   "├",
		RightJunction:  "┤",
		BottomJunction: "┴",
		InnerJunction:  "┼",

		InnerDivider: "│",
	}
)

func init() {
	borderDefault.generateStyles()
	borderRounded.generateStyles()
}

func (b *Border) generateStyles() { _ = "STUB: not implemented"; return }

// The footer is a single cell with the top taken off... usually.  We can
// re-enable the top if needed this way for certain format configurations.

func (b *Border) styleLeftWithFooter(original lipgloss.Style) lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

func (b *Border) styleRightWithFooter(original lipgloss.Style) lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

func (b *Border) styleBothWithFooter(original lipgloss.Style) lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

// This function is long, but it's just repetitive...
//
//nolint:funlen
func (b *Border) generateMultiStyles() { _ = "STUB: not implemented"; return }

func (b *Border) generateSingleColumnStyles() { _ = "STUB: not implemented"; return }

func (b *Border) generateSingleRowStyles() { _ = "STUB: not implemented"; return }

func (b *Border) generateSingleCellStyle() { _ = "STUB: not implemented"; return }

// BorderDefault uses the basic square border, useful to reset the border if
// it was changed somehow.
func (m Model) BorderDefault() Model {
	_ = "STUB: not implemented"
	// Already generated styles
	return *new(Model)
}

// BorderRounded uses a thin, rounded border.
func (m Model) BorderRounded() Model {
	_ = "STUB: not implemented"
	// Already generated styles
	return *new(Model)
}

// Border uses the given border components to render the table.
func (m Model) Border(border Border) Model { _ = "STUB: not implemented"; return *new(Model) }

type borderStyleRow struct {
	left  lipgloss.Style
	inner lipgloss.Style
	right lipgloss.Style
}

func (b *borderStyleRow) inherit(s lipgloss.Style) { _ = "STUB: not implemented"; return }

// There's a lot of branches here, but splitting it up further would make it
// harder to follow. So just be careful with comments and make sure it's tested!
//
//nolint:nestif
func (m Model) styleHeaders() borderStyleRow {
	_ = "STUB: not implemented"
	return *new(borderStyleRow)
}

// Possible configurations:
// - Single cell
// - Single row
// - Single column
// - Multi

// Single column

// Single cell

// Single row

// Multi

func (m Model) styleRows() (inner borderStyleRow, last borderStyleRow) {
	_ = "STUB: not implemented"
	return *new(borderStyleRow), *new(borderStyleRow)
}
