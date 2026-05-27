package table

// View renders the table. It does not end in a newline, so that it can be
// composed with other elements more consistently.
//
//nolint:cyclop
func (m Model) View() string {
	_ = "STUB: not implemented"
	// Safety valve for empty tables
	return ""
}

//nolint: mnd // This is just getting the first newlined substring
