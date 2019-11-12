package domain

import (
	"math/rand"
	"time"
)

const (
	// RollMax is a constant defining the highest possible roll on a D20
	RollMax int = 20
	// RollMin is a constant defining the highest possible roll on a D20
	RollMin int = 1
)

// Character holds the initiative-related state of a D&D character
type Character struct {
	Name               string `json:"name"`
	InitiativeModifier int    `json:"modifier"`
	Initiative         int
}

// Roll generates the initiative roll for a character - based on a D20
func (c *Character) Roll() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(RollMax-RollMin) + RollMin
}
