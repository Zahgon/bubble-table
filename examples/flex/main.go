package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
)

const (
	columnKeyName        = "name"
	columnKeyElement     = "element"
	columnKeyDescription = "description"

	minWidth  = 30
	minHeight = 8

	// Add a fixed margin to account for description & instructions
	fixedVerticalMargin = 4
)

type Model struct {
	flexTable table.Model

	// Window dimensions
	totalWidth  int
	totalHeight int

	// Table dimensions
	horizontalMargin int
	verticalMargin   int
}

func NewModel() Model { _ = "STUB: not implemented"; return *new(Model) }

// This table uses flex columns, but it will still need a target
// width in order to know what width it should fill.  In this example
// the target width is set below in `recalculateTable`, which sets
// the table to the width of the screen to demonstrate resizing
// with flex columns.

func (m Model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m *Model) recalculateTable() { _ = "STUB: not implemented"; return }

func (m Model) calculateWidth() int { _ = "STUB: not implemented"; return 0 }

func (m Model) calculateHeight() int { _ = "STUB: not implemented"; return 0 }

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

func main() {
	p := tea.NewProgram(NewModel())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
