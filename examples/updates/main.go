package main

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

const (
	columnKeyID     = "id"
	columnKeyScore  = "score"
	columnKeyStatus = "status"
)

var (
	styleCritical = lipgloss.NewStyle().Foreground(lipgloss.Color("#f00"))
	styleStable   = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0"))
	styleGood     = lipgloss.NewStyle().Foreground(lipgloss.Color("#0f0"))
)

type Model struct {
	table table.Model

	updateDelay time.Duration

	data []*SomeData
}

func rowStyleFunc(input table.RowStyleFuncInput) lipgloss.Style {
	_ = "STUB: not implemented"
	return *new(lipgloss.Style)
}

func NewModel() Model { _ = "STUB: not implemented"; return *new(Model) }

// This data is stored somewhere else, maybe on a client or some other thing
func refreshDataCmd() tea.Msg {
	_ = "STUB: not implemented"
	// This could come from some API or something
	return *new(tea.Msg)
}

// Generate columns based on how many are critical to show some summary
func generateColumns(numCritical int) []table.Column {
	_ = "STUB: not implemented"
	// Show how many critical there are
	return nil
}

// This normally applies the critical style to everything in the column,
// but in this case we apply a row style which overrides it anyway.

func (m Model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// Reapply the new data and the new columns based on critical count

// This can be from any source, but for demo purposes let's party!

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

func generateRowsFromData(data []*SomeData) []table.Row { _ = "STUB: not implemented"; return nil }

func main() {

	p := tea.NewProgram(NewModel())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
