package table

// GetColumnSorting returns the current sorting rules for the table as a list of
// SortColumns, which are applied from first to last.  This means that data will
// be grouped by the later elements in the list.  The returned list is a copy
// and modifications will have no effect.
func (m *Model) GetColumnSorting() []SortColumn { _ = "STUB: not implemented"; return nil }

// GetCanFilter returns true if the table enables filtering at all.  This does
// not say whether a filter is currently active, only that the feature is enabled.
func (m *Model) GetCanFilter() bool {
	_ = "STUB: not implemented"

	// GetIsFilterActive returns true if the table is currently being filtered.  This
	// does not say whether the table CAN be filtered, only whether or not a filter
	// is actually currently being applied.
	return false
}

func (m *Model) GetIsFilterActive() bool { _ = "STUB: not implemented"; return false }

// GetIsFilterInputFocused returns true if the table's built-in filter input is
// currently focused.
func (m *Model) GetIsFilterInputFocused() bool { _ = "STUB: not implemented"; return false }

// GetCurrentFilter returns the current filter text being applied, or an empty
// string if none is applied.
func (m *Model) GetCurrentFilter() string { _ = "STUB: not implemented"; return "" }

// GetVisibleRows returns sorted and filtered rows.
func (m *Model) GetVisibleRows() []Row { _ = "STUB: not implemented"; return nil }

// GetHighlightedRowIndex returns the index of the Row that's currently highlighted
// by the user.
func (m *Model) GetHighlightedRowIndex() int { _ = "STUB: not implemented"; return 0 }

// GetFocused returns whether or not the table is focused and is receiving inputs.
func (m *Model) GetFocused() bool {
	_ = "STUB: not implemented"

	// GetHorizontalScrollColumnOffset returns how many columns to the right the table
	// has been scrolled.  0 means the table is all the way to the left, which is
	// the starting default.
	return false
}

func (m *Model) GetHorizontalScrollColumnOffset() int { _ = "STUB: not implemented"; return 0 }

// GetHeaderVisibility returns true if the header has been set to visible (default)
// or false if the header has been set to hidden.
func (m *Model) GetHeaderVisibility() bool { _ = "STUB: not implemented"; return false }

// GetFooterVisibility returns true if the footer has been set to
// visible (default) or false if the footer has been set to hidden.
// Note that even if the footer is visible it will only be rendered if
// it has contents.
func (m *Model) GetFooterVisibility() bool { _ = "STUB: not implemented"; return false }

// GetPaginationWrapping returns true if pagination wrapping is enabled, or false
// if disabled.  If disabled, navigating through pages will stop at the first
// and last pages.
func (m *Model) GetPaginationWrapping() bool { _ = "STUB: not implemented"; return false }
