// This file contains a full demo of most available features, for both testing
// and for reference
package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
)

const (
	columnKeyID          = "id"
	columnKeyName        = "name"
	columnKeyDescription = "description"
	columnKeyCount       = "count"
)

var (
	customBorder = table.Border{
		Top:    "─",
		Left:   "│",
		Right:  "│",
		Bottom: "─",

		TopRight:    "╮",
		TopLeft:     "╭",
		BottomRight: "╯",
		BottomLeft:  "╰",

		TopJunction:    "╥",
		LeftJunction:   "├",
		RightJunction:  "┤",
		BottomJunction: "╨",
		InnerJunction:  "╫",

		InnerDivider: "║",
	}
)

type Model struct {
	tableModel table.Model
}

func NewModel() Model { _ = "STUB: not implemented"; return *new(Model) }

// Missing name

// Apply a style to this cell

// Start with the default key map and change it slightly, just for demoing

// Throw features in... the point is not to look good, it's just reference!

func (m Model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *Model) updateFooter() { _ = "STUB: not implemented"; return }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// We control the footer text, so make sure to update it

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

// Slightly dangerous type assumption but fine for demo

func main() {
	p := tea.NewProgram(NewModel())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
