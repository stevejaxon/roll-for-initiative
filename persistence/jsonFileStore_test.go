package persistence

import (
	"testing"

	"github.com/stevejaxon/roll-for-initiative/domain"
)

func TestCreateFirstCharacter(t *testing.T) {
	// Setup
	const dbPath = "/tmp/first_character.json"
	character := &domain.Character{
		Name:               "Test1",
		InitiativeModifier: 1,
	}
	db := &JSONCharacterStore{
		DBFilePath: dbPath,
	}

	// Test
	err := db.Create(character)

	// Verify
	if err != nil {
		t.Fatalf("expected that it would be possible to create the first character in the DB: %v", err)
	}
}
