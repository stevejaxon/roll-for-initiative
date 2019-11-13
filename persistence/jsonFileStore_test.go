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

func TestCreateCalledMultipleTimes(t *testing.T) {
	// Setup
	const dbPath = "/tmp/multiple_characters.json"
	db := &JSONCharacterStore{
		DBFilePath: dbPath,
	}
	testCases := []*domain.Character{
		&domain.Character{
			Name:               "Test1",
			InitiativeModifier: -1,
		},
		&domain.Character{
			Name:               "Test2",
			InitiativeModifier: 10,
		},
		&domain.Character{
			Name:               "Test3",
			InitiativeModifier: -10,
		},
		&domain.Character{
			Name:               "Test4",
			InitiativeModifier: 3,
		},
	}

	// Test
	for _, testCase := range testCases {
		err := db.Create(testCase)

		// Verify
		if err != nil {
			t.Fatalf("expected that it would be possible to create characters in the DB: %v", err)
		}
	}

	// Verify
}
