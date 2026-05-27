package table

// SortDirection indicates whether a column should sort by ascending or descending.
type SortDirection int

const (
	// SortDirectionAsc indicates the column should be in ascending order.
	SortDirectionAsc SortDirection = iota

	// SortDirectionDesc indicates the column should be in descending order.
	SortDirectionDesc
)

// SortColumn describes which column should be sorted and how.
type SortColumn struct {
	ColumnKey string
	Direction SortDirection
}

// SortByAsc sets the main sorting column to the given key, in ascending order.
// If a previous sort was used, it is replaced by the given column each time
// this function is called.  Values are sorted as numbers if possible, or just
// as simple string comparisons if not numbers.
func (m Model) SortByAsc(columnKey string) Model { _ = "STUB: not implemented"; return *new(Model) }

// SortByDesc sets the main sorting column to the given key, in descending order.
// If a previous sort was used, it is replaced by the given column each time
// this function is called.  Values are sorted as numbers if possible, or just
// as simple string comparisons if not numbers.
func (m Model) SortByDesc(columnKey string) Model { _ = "STUB: not implemented"; return *new(Model) }

// ThenSortByAsc provides a secondary sort after the first, in ascending order.
// Can be chained multiple times, applying to smaller subgroups each time.
func (m Model) ThenSortByAsc(columnKey string) Model { _ = "STUB: not implemented"; return *new(Model) }

// ThenSortByDesc provides a secondary sort after the first, in descending order.
// Can be chained multiple times, applying to smaller subgroups each time.
func (m Model) ThenSortByDesc(columnKey string) Model {
	_ = "STUB: not implemented"
	return *new(Model)
}

type sortableTable struct {
	rows     []Row
	byColumn SortColumn
}

func (s *sortableTable) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *sortableTable) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s *sortableTable) extractString(i int, column string) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *sortableTable) extractNumber(i int, column string) (float64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (s *sortableTable) Less(first, second int) bool { _ = "STUB: not implemented"; return false }

func getSortedRows(sortOrder []SortColumn, rows []Row) []Row { _ = "STUB: not implemented"; return nil }
