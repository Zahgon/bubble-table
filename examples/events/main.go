package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

type Element string

const (
	columnKeyName    = "name"
	columnKeyElement = "element"

	// This is not a visible column, but is used to attach useful reference data
	// to the row itself for easier retrieval
	columnKeyPokemonData = "pokedata"

	elementNormal   Element = "Normal"
	elementFire     Element = "Fire"
	elementElectric Element = "Electric"
	elementWater    Element = "Water"
	elementPlant    Element = "Plant"
)

var (
	styleSubtle = lipgloss.NewStyle().Foreground(lipgloss.Color("#888"))

	styleBase = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#a7a")).
			BorderForeground(lipgloss.Color("#a38")).
			Align(lipgloss.Right)

	elementColors = map[Element]string{
		elementNormal:   "#fa0",
		elementFire:     "#f64",
		elementElectric: "#ff0",
		elementWater:    "#44f",
		elementPlant:    "#8b8",
	}
)

type Pokemon struct {
	Name                     string
	Element                  Element
	ConversationCount        int
	PositiveSentimentPercent float32
	NegativeSentimentPercent float32
}

func NewPokemon(name string, element Element, conversationCount int, positiveSentimentPercent float32, negativeSentimentPercent float32) Pokemon {
	_ = "STUB: not implemented"
	return *new(Pokemon)
}

func (p Pokemon) ToRow() table.Row { _ = "STUB: not implemented"; return *new(table.Row) }

// This isn't a visible column, but we can add the data here anyway for later retrieval

type Model struct {
	pokeTable table.Model

	currentPokemonData Pokemon

	lastSelectedEvent table.UserEventRowSelectToggled
}

func NewModel() Model { _ = "STUB: not implemented"; return *new(Model) }

func (m Model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// We can pretend this is an async data retrieval, but really we already
// have the data, so just return it after some fake delay.  Also note
// that the event has some data attached to it, but we're ignoring
// that for this example as we just want the current highlighted row.

func (m Model) View() string { _ = "STUB: not implemented"; return "" }

func main() {
	p := tea.NewProgram(NewModel())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
