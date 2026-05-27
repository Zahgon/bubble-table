package table

import "github.com/charmbracelet/bubbles/key"

// KeyMap defines the keybindings for the table when it's focused.
type KeyMap struct {
	RowDown key.Binding
	RowUp   key.Binding

	RowSelectToggle key.Binding

	PageDown  key.Binding
	PageUp    key.Binding
	PageFirst key.Binding
	PageLast  key.Binding

	// Filter allows the user to start typing and filter the rows.
	Filter key.Binding

	// FilterBlur is the key that stops the user's input from typing into the filter.
	FilterBlur key.Binding

	// FilterClear will clear the filter while it's blurred.
	FilterClear key.Binding

	// ScrollRight will move one column to the right when overflow occurs.
	ScrollRight key.Binding

	// ScrollLeft will move one column to the left when overflow occurs.
	ScrollLeft key.Binding
}

// DefaultKeyMap returns a set of sensible defaults for controlling a focused table with help text.
func DefaultKeyMap() KeyMap { _ = "STUB: not implemented"; return *new(KeyMap) }

// FullHelp returns a multi row view of all the helpkeys that are defined. Needed to fullfil the 'help.Model' interface.
// Also appends all user defined extra keys to the help.
func (m Model) FullHelp() [][]key.Binding { _ = "STUB: not implemented"; return nil }

// ShortHelp just returns a single row of help views. Needed to fullfil the 'help.Model' interface.
// Also appends all user defined extra keys to the help.
func (m Model) ShortHelp() []key.Binding { _ = "STUB: not implemented"; return nil }
