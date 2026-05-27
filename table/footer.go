package table

func (m Model) hasFooter() bool { _ = "STUB: not implemented"; return false }

func (m Model) renderFooter(width int, includeTop bool) string {
	_ = "STUB: not implemented"
	return ""
}

// paged feature enabled

// Need to apply inline style here in case of filter input cursor, because
// the input cursor resets the style after rendering.  Note that Inline(true)
// creates a copy, so it's safe to use here without mutating the underlying
// base style.
