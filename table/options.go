package table

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

// RowStyleFuncInput is the input to the style function that can
// be applied to each row.  This is useful for things like zebra
// striping or other data-based styles.
//
// Note that we use a struct here to allow for future expansion
// while keeping backwards compatibility.
type RowStyleFuncInput struct {
	// Index is the index of the row, starting at 0.
	Index int

	// Row is the full row data.
	Row Row

	// IsHighlighted is true if the row is currently highlighted.
	IsHighlighted bool
}

// WithRowStyleFunc sets a function that can be used to apply a style to each row
// based on the row data.  This is useful for things like zebra striping or other
// data-based styles.  It can be safely set to nil to remove it later.
// This style is applied after the base style and before individual row styles.
// This will override any HighlightStyle settings.
func (m Model) WithRowStyleFunc(f func(RowStyleFuncInput) lipgloss.Style) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// WithHighlightedRow sets the highlighted row to the given index.
func (m Model) WithHighlightedRow(index int) Model { _ = "STUB: not implemented"; return *new(Model) }

// HeaderStyle sets the style to apply to the header text, such as color or bold.
func (m Model) HeaderStyle(style lipgloss.Style) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// WithRows sets the rows to show as data in the table.
func (m Model) WithRows(rows []Row) Model { _ = "STUB: not implemented"; return *new(Model) }

// MaxPages is 1-index, currentPage is 0 index

// WithKeyMap sets the key map to use for controls when focused.
func (m Model) WithKeyMap(keyMap KeyMap) Model { _ = "STUB: not implemented"; return *new(Model) }

// KeyMap returns a copy of the current key map in use.
func (m Model) KeyMap() KeyMap {
	_ = "STUB: not implemented"

	// SelectableRows sets whether or not rows are selectable.  If set, adds a column
	// in the front that acts as a checkbox and responds to controls if Focused.
	return *new(KeyMap)
}

func (m Model) SelectableRows(selectable bool) Model { _ = "STUB: not implemented"; return *new(Model) }

// HighlightedRow returns the full Row that's currently highlighted by the user.
func (m Model) HighlightedRow() Row { _ = "STUB: not implemented"; return *new(Row) }

// TODO: Better way to do this without pointers/nil?  Or should it be nil?

// SelectedRows returns all rows that have been set as selected by the user.
func (m Model) SelectedRows() []Row { _ = "STUB: not implemented"; return nil }

// HighlightStyle sets a custom style to use when the row is being highlighted
// by the cursor.  This should not be used with WithRowStyleFunc.  Instead, use
// the IsHighlighted field in the style function.
func (m Model) HighlightStyle(style lipgloss.Style) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// Focused allows the table to show highlighted rows and take in controls of
// up/down/space/etc to let the user navigate the table and interact with it.
func (m Model) Focused(focused bool) Model { _ = "STUB: not implemented"; return *new(Model) }

// Filtered allows the table to show rows that match the filter.
func (m Model) Filtered(filtered bool) Model { _ = "STUB: not implemented"; return *new(Model) }

// StartFilterTyping focuses the text input to allow user typing to filter.
func (m Model) StartFilterTyping() Model { _ = "STUB: not implemented"; return *new(Model) }

// WithStaticFooter adds a footer that only displays the given text.
func (m Model) WithStaticFooter(footer string) Model { _ = "STUB: not implemented"; return *new(Model) }

// WithPageSize enables pagination using the given page size.  This can be called
// again at any point to resize the height of the table.
func (m Model) WithPageSize(pageSize int) Model { _ = "STUB: not implemented"; return *new(Model) }

// WithNoPagination disables pagination in the table.
func (m Model) WithNoPagination() Model { _ = "STUB: not implemented"; return *new(Model) }

// WithPaginationWrapping sets whether to wrap around from the beginning to the
// end when navigating through pages.  Defaults to true.
func (m Model) WithPaginationWrapping(wrapping bool) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// WithSelectedText describes what text to show when selectable rows are enabled.
// The selectable column header will use the selected text string.
func (m Model) WithSelectedText(unselected, selected string) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// WithBaseStyle applies a base style as the default for everything in the table.
// This is useful for border colors, default alignment, default color, etc.
func (m Model) WithBaseStyle(style lipgloss.Style) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// WithTargetWidth sets the total target width of the table, including borders.
// This only takes effect when using flex columns.  When using flex columns,
// columns will stretch to fill out to the total width given here.
func (m Model) WithTargetWidth(totalWidth int) Model { _ = "STUB: not implemented"; return *new(Model) }

// WithMinimumHeight sets the minimum total height of the table, including borders.
func (m Model) WithMinimumHeight(minimumHeight int) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// PageDown goes to the next page of a paginated table, wrapping to the first
// page if the table is already on the last page.
func (m Model) PageDown() Model {
	_ = "STUB: not implemented"
	return *

	// PageUp goes to the previous page of a paginated table, wrapping to the
	// last page if the table is already on the first page.
	new(Model)
}

func (m Model) PageUp() Model {
	_ = "STUB: not implemented"
	return *

	// PageLast goes to the last page of a paginated table.
	new(Model)
}

func (m Model) PageLast() Model {
	_ = "STUB: not implemented"
	return *

	// PageFirst goes to the first page of a paginated table.
	new(Model)
}

func (m Model) PageFirst() Model {
	_ = "STUB: not implemented"
	return *

	// WithCurrentPage sets the current page (1 as the first page) of a paginated
	// table, bounded to the total number of pages.  The current selected row will
	// be set to the top row of the page if the page changed.
	new(Model)
}

func (m Model) WithCurrentPage(currentPage int) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// WithColumns sets the visible columns for the table, so that columns can be
// added/removed/resized or headers rewritten.
func (m Model) WithColumns(columns []Column) Model {
	_ = "STUB: not implemented"
	// Deep copy to avoid edits
	return *new(Model)
}

// Re-add the selectable column

// WithFilterInput makes the table use the provided text input bubble for
// filtering rather than using the built-in default.  This allows for external
// text input controls to be used.
func (m Model) WithFilterInput(input textinput.Model) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// WithFilterInputValue sets the filter value to the given string, immediately
// applying it as if the user had typed it in.  Useful for external filter inputs
// that are not necessarily a text input.
func (m Model) WithFilterInputValue(value string) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// WithFilterFunc adds a filter function to the model. If the function returns
// true, the row will be included in the filtered results. If the function
// is nil, the function won't be used and instead the default filtering will be applied,
// if any.
func (m Model) WithFilterFunc(shouldInclude FilterFunc) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// WithFuzzyFilter enables fuzzy filtering for the table.
func (m Model) WithFuzzyFilter() Model { _ = "STUB: not implemented"; return *new(Model) }

// WithFooterVisibility sets the visibility of the footer.
func (m Model) WithFooterVisibility(visibility bool) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// WithHeaderVisibility sets the visibility of the header.
func (m Model) WithHeaderVisibility(visibility bool) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// WithMaxTotalWidth sets the maximum total width that the table should render.
// If this width is exceeded by either the target width or by the total width
// of all the columns (including borders!), anything extra will be treated as
// overflow and horizontal scrolling will be enabled to see the rest.
func (m Model) WithMaxTotalWidth(maxTotalWidth int) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// WithHorizontalFreezeColumnCount freezes the given number of columns to the
// left side.  This is useful for things like ID or Name columns that should
// always be visible even when scrolling.
func (m Model) WithHorizontalFreezeColumnCount(columnsToFreeze int) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// ScrollRight moves one column to the right.  Use with WithMaxTotalWidth.
func (m Model) ScrollRight() Model { _ = "STUB: not implemented"; return *new(Model) }

// ScrollLeft moves one column to the left.  Use with WithMaxTotalWidth.
func (m Model) ScrollLeft() Model {
	_ = "STUB: not implemented"
	return *

	// WithMissingDataIndicator sets an indicator to use when data for a column is
	// not found in a given row.  Note that this is for completely missing data,
	// an empty string or other zero value that is explicitly set is not considered
	// to be missing.
	new(Model)
}

func (m Model) WithMissingDataIndicator(str string) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// WithMissingDataIndicatorStyled sets a styled indicator to use when data for
// a column is not found in a given row.  Note that this is for completely
// missing data, an empty string or other zero value that is explicitly set is
// not considered to be missing.
func (m Model) WithMissingDataIndicatorStyled(styled StyledCell) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// WithAllRowsDeselected deselects any rows that are currently selected.
func (m Model) WithAllRowsDeselected() Model { _ = "STUB: not implemented"; return *new(Model) }

// WithMultiline sets whether or not to wrap text in cells to multiple lines.
func (m Model) WithMultiline(multiline bool) Model { _ = "STUB: not implemented"; return *new(Model) }

// WithAdditionalShortHelpKeys enables you to add more keybindings to the 'short help' view.
func (m Model) WithAdditionalShortHelpKeys(keys []key.Binding) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// WithAdditionalFullHelpKeys enables you to add more keybindings to the 'full help' view.
func (m Model) WithAdditionalFullHelpKeys(keys []key.Binding) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

// WithGlobalMetadata applies the given metadata to the table. This metadata is passed to
// some functions in FilterFuncInput and StyleFuncInput to enable more advanced decisions,
// such as setting some global theme variable to reference, etc. Has no effect otherwise.
func (m Model) WithGlobalMetadata(metadata map[string]any) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}
