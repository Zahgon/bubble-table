package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

const (
	columnKeyName              = "name"
	columnKeyElement           = "element"
	columnKeyConversations     = "convos"
	columnKeyPositiveSentiment = "positive"
	columnKeyNegativeSentiment = "negative"

	colorNormal   = "#fa0"
	colorElectric = "#ff0"
	colorFire     = "#f64"
	colorPlant    = "#8b8"
	colorWater    = "#44f"
)

var (
	styleSubtle = lipgloss.NewStyle().Foreground(lipgloss.Color("#888"))

	styleBase = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#a7a")).
			BorderForeground(lipgloss.Color("#a38")).
			Align(lipgloss.Right)
)

type Model struct {
	pokeTable            table.Model
	favoriteElementIndex int
}

var elementList = []string{
	"Normal",
	"Electric",
	"Fire",
	"Plant",
	"Water",
}

var colorMap = map[any]string{
	"Electric": colorElectric,
	"Fire":     colorFire,
	"Plant":    colorPlant,
	"Water":    colorWater,
}

func makeRow(name, element string, numConversations int, positiveSentiment, negativeSentiment float32) table.Row {
	_ = "STUB: not implemented"
	return *new(table.Row)
}

func genMetadata(favoriteElementIndex int) map[string]any { _ = "STUB: not implemented"; return nil }

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
