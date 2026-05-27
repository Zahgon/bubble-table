package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
)

const (
	columnKeyID = "id"

	numCols = 100
	numRows = 10
	idWidth = 5

	colWidth = 3
	maxWidth = 30
)

type Model struct {
	scrollableTable table.Model
}

func colKey(colNum int) string { _ = "STUB: not implemented"; return "" }

func genRow(id int) table.Row { _ = "STUB: not implemented"; return *new(table.Row) }

func NewModel() Model { _ = "STUB: not implemented"; return *new(Model) }

func (m Model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

func main() {
	p := tea.NewProgram(NewModel())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
