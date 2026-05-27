package table

// FilterFuncInput is the input to a FilterFunc. It's a struct so we can add more things later
// without breaking compatibility.
type FilterFuncInput struct {
	// Columns is a list of the columns of the table
	Columns []Column

	// Row is the row that's being considered for filtering
	Row Row

	// GlobalMetadata is an arbitrary set of metadata from the table set by WithGlobalMetadata
	GlobalMetadata map[string]any

	// Filter is the filter string input to consider
	Filter string
}

// FilterFunc takes a FilterFuncInput and returns true if the row should be visible,
// or false if the row should be hidden.
type FilterFunc func(FilterFuncInput) bool

func (m Model) getFilteredRows(rows []Row) []Row { _ = "STUB: not implemented"; return nil }

// filterFuncContains returns a filterFunc that performs case-insensitive
// "contains" matching over all filterable columns in a row.
func filterFuncContains(input FilterFuncInput) bool { _ = "STUB: not implemented"; return false }

// Extract internal StyledCell data

// filterFuncFuzzy returns a filterFunc that performs case-insensitive fuzzy
// matching (subsequence) over the concatenation of all filterable column values.
func filterFuncFuzzy(input FilterFuncInput) bool { _ = "STUB: not implemented"; return false }

// uses Stringer if implemented

// fuzzySubsequenceMatch returns true if all runes in needle appear in order
// within haystack (not necessarily contiguously). Case must be normalized by caller.
func fuzzySubsequenceMatch(haystack, needle string) bool { _ = "STUB: not implemented"; return false }
