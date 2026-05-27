package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
)

type Model struct {
	tableDefault        table.Model
	tableWithRowIndices table.Model

	rowCount int
}

func genRows(columnCount int, rowCount int) []table.Row { _ = "STUB: not implemented"; return nil }

func genTable(columnCount int, rowCount int) table.Model {
	_ = "STUB: not implemented"
	return *new(table.Model)
}

func NewModel() Model { _ = "STUB: not implemented"; return *new(Model) }

func (m *Model) regenTableRows() { _ = "STUB: not implemented"; return }

func (m Model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// Write a custom footer

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

func main() {
	p := tea.NewProgram(NewModel())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
