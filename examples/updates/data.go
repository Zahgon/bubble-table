package main

// SomeData represent some real data of some sort, unaware of tables
type SomeData struct {
	ID     string
	Score  int
	Status string
}

// NewSomeData creates SomeData that has an ID and randomized values
func NewSomeData(id string) *SomeData { _ = "STUB: not implemented"; return nil }

// Start with some random data

// RandomizeScoreAndStatus does an in-place update to simulate some data being
// updated by some other process
func (s *SomeData) RandomizeScoreAndStatus() { _ = "STUB: not implemented"; return }
