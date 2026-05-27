package table

func (m *Model) recalculateWidth() { _ = "STUB: not implemented"; return }

// Updates column width in-place.  This could be optimized but should be called
// very rarely so we prioritize simplicity over performance here.
func updateColumnWidths(cols []Column, totalWidth int) { _ = "STUB: not implemented"; return }

// We use the GCD here because otherwise very large values won't divide
// nicely as ints

// Take borders into account for the actual style

func (m *Model) recalculateHeight() { _ = "STUB: not implemented"; return }

// Header always has the top border

func (m *Model) calculatePadding(numRows int) int { _ = "STUB: not implemented"; return 0 }

// additional 1 for bottom border

// This is an edge case where we want to add 1 additional line of height, i.e.
// add a border without an empty row. However, this is not possible, so we need
// to add an extra row which will result in the table being 1 row taller than
// the requested minimum height.

// Table is already larger than minimum height, do nothing.
