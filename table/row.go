package table

import (
	"github.com/charmbracelet/lipgloss"
)

// RowData is a map of string column keys to arbitrary data.  Data with a key
// that matches a column key will be displayed.  Data with a key that does not
// match a column key will not be displayed, but will remain attached to the Row.
// This can be useful for attaching hidden metadata for future reference when
// retrieving rows.
type RowData map[string]any

// Row represents a row in the table with some data keyed to the table columns>
// Can have a style applied to it such as color/bold.  Create using NewRow().
type Row struct {
	Style lipgloss.Style
	Data  RowData

	selected bool

	// id is an internal unique ID to match rows after they're copied
	id uint32
}

var lastRowID uint32 = 1

// NewRow creates a new row and copies the given row data.
func NewRow(data RowData) Row { _ = "STUB: not implemented"; return *new(Row) }

// Doesn't deep copy val, but close enough for now...

// WithStyle uses the given style for the text in the row.
func (r Row) WithStyle(style lipgloss.Style) Row { _ = "STUB: not implemented"; return *new(Row) }

//nolint:cyclop,funlen // Breaking this up will be more complicated than it's worth for now
func (m Model) renderRowColumnData(row Row, column Column, rowStyle lipgloss.Style, borderStyle lipgloss.Style) string {
	_ = "STUB: not implemented"
	return ""
}

func (m Model) renderRow(rowIndex int, last bool) string { _ = "STUB: not implemented"; return "" }

func (m Model) renderBlankRow(last bool) string { _ = "STUB: not implemented"; return "" }

// This is long and could use some refactoring in the future, but not quite sure
// how to pick it apart yet.
//
//nolint:funlen, cyclop
func (m Model) renderRowData(row Row, rowStyle lipgloss.Style, last bool) string {
	_ = "STUB: not implemented"
	return ""
}

// If this is the last header, we don't need to account for the
// overflow arrow column

// Selected returns a copy of the row that's set to be selected or deselected.
// The old row is not changed in-place.
func (r Row) Selected(selected bool) Row { _ = "STUB: not implemented"; return *new(Row) }
