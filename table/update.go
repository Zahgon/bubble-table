package table

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) moveHighlightUp() { _ = "STUB: not implemented"; return }

func (m *Model) moveHighlightDown() { _ = "STUB: not implemented"; return }

func (m *Model) toggleSelect() { _ = "STUB: not implemented"; return }

func (m Model) updateFilterTextInput(msg tea.Msg) (Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(Model), *new(tea.Cmd)
}

// This is a series of Matches tests with minimal logic
//
//nolint:cyclop
func (m *Model) handleKeypress(msg tea.KeyMsg) { _ = "STUB: not implemented"; return }

// Update responds to input from the user or other messages from Bubble Tea.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(Model), *new(tea.Cmd)
}
