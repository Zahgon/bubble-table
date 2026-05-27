package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
)

const (
	columnKeyName = "name"
	columnKeyType = "type"
	columnKeyWins = "wins"
)

type Model struct {
	simpleTable table.Model

	columnSortKey string
	sortDirection string
}

func NewModel() Model { _ = "STUB: not implemented"; return *new(Model) }

func (m Model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// Within the same type, order each by wins

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

func main() {
	p := tea.NewProgram(NewModel())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
