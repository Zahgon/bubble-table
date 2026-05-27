package table

// PageSize returns the current page size for the table, or 0 if there is no
// pagination enabled.
func (m *Model) PageSize() int {
	_ = "STUB: not implemented"

	// CurrentPage returns the current page that the table is on, starting from an
	// index of 1.
	return 0
}

func (m *Model) CurrentPage() int { _ = "STUB: not implemented"; return 0 }

// MaxPages returns the maximum number of pages that are visible.
func (m *Model) MaxPages() int { _ = "STUB: not implemented"; return 0 }

// TotalRows returns the current total row count of the table.  If the table is
// paginated, this is the total number of rows across all pages.
func (m *Model) TotalRows() int { _ = "STUB: not implemented"; return 0 }

// VisibleIndices returns the current visible rows by their 0 based index.
// Useful for custom pagination footers.
func (m *Model) VisibleIndices() (start, end int) { _ = "STUB: not implemented"; return 0, 0 }

func (m *Model) pageDown() { _ = "STUB: not implemented"; return }

func (m *Model) pageUp() { _ = "STUB: not implemented"; return }

func (m *Model) pageFirst() { _ = "STUB: not implemented"; return }

func (m *Model) pageLast() { _ = "STUB: not implemented"; return }

func (m *Model) expectedPageForRowIndex(rowIndex int) int { _ = "STUB: not implemented"; return 0 }
