package table

import (
	"github.com/charmbracelet/lipgloss"
)

// Column is a column in the table.
type Column struct {
	title string
	key   string
	width int

	flexFactor int

	filterable bool
	style      lipgloss.Style

	fmtString string
}

// NewColumn creates a new fixed-width column with the given information.
func NewColumn(key, title string, width int) Column { _ = "STUB: not implemented"; return *new(Column) }

// NewFlexColumn creates a new flexible width column that tries to fill in the
// total table width.  If multiple flex columns exist, each will measure against
// each other depending on their flexFactor.  For example, if both have a flexFactor
// of 1, they will have equal width.  If one has a flexFactor of 1 and the other
// has a flexFactor of 3, the second will be 3 times larger than the first.  You
// must use WithTargetWidth if you have any flex columns, so that the table knows
// how much width it should fill.
func NewFlexColumn(key, title string, flexFactor int) Column {
	_ = "STUB: not implemented"
	return *new(Column)
}

// WithStyle applies a style to the column as a whole.
func (c Column) WithStyle(style lipgloss.Style) Column {
	_ = "STUB: not implemented"
	return *new(Column)
}

// WithFiltered sets whether the column should be considered for filtering (true)
// or not (false).
func (c Column) WithFiltered(filterable bool) Column {
	_ = "STUB: not implemented"
	return *new(Column)
}

// WithFormatString sets the format string used by fmt.Sprintf to display the data.
// If not set, the default is "%v" for all data types.  Intended mainly for
// numeric formatting.
//
// Since data is of the any type, make sure that all data in the column
// is of the expected type or the format may fail.  For example, hardcoding '3'
// instead of '3.0' and using '%.2f' will fail because '3' is an integer.
func (c Column) WithFormatString(fmtString string) Column {
	_ = "STUB: not implemented"
	return *new(Column)
}

func (c *Column) isFlex() bool { _ = "STUB: not implemented"; return false }

// Title returns the title of the column.
func (c Column) Title() string {
	_ = "STUB: not implemented"

	// Key returns the key of the column.
	return ""
}

func (c Column) Key() string {
	_ = "STUB: not implemented"

	// Width returns the width of the column.
	return ""
}

func (c Column) Width() int {
	_ = "STUB: not implemented"

	// FlexFactor returns the flex factor of the column.
	return 0
}

func (c Column) FlexFactor() int { _ = "STUB: not implemented"; return 0 }

// IsFlex returns whether the column is a flex column.
func (c Column) IsFlex() bool {
	_ = "STUB: not implemented"

	// Filterable returns whether the column is filterable.
	return false
}

func (c Column) Filterable() bool { _ = "STUB: not implemented"; return false }

// Style returns the style of the column.
func (c Column) Style() lipgloss.Style {
	_ = "STUB: not implemented"

	// FmtString returns the format string of the column.
	return *new(lipgloss.Style)
}

func (c Column) FmtString() string { _ = "STUB: not implemented"; return "" }
