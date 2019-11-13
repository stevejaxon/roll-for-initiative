package persistence

import (
	"fmt"
	"path/filepath"
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

	storedChars, err := db.RetrieveAllCharacters()
	if err != nil {
		t.Fatalf("expected that it would be possible to create the first character in the DB: %v", err)
	}
	if len(storedChars) != 1 {
		t.Fatalf("expected that it would be possible to create the first character in the DB: %v", err)
	}
	actual := &storedChars[0]
	if match := compareCharacters(character, actual); !match {
		t.Fatalf("expected that the retrieved character would match the created one")
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
	storedChars, err := db.RetrieveAllCharacters()
	if err != nil {
		t.Fatalf("expected that it would be possible to retrieve the characters recently created: %v", err)
	}
	if len(storedChars) != len(testCases) {
		t.Fatalf("expected that the number of retrieved would match the number of characters recently created: %v", err)
	}
	for i, actual := range storedChars {
		if match := compareCharacters(testCases[i], &actual); !match {
			t.Fatalf("expected that the retrieved character would match the created one")
		}
	}
}

func TestGetAllCharacters(t *testing.T) {
	// Setup
	dbPath := filepath.Join("..", "testdata", "characterdb.json")
	db := &JSONCharacterStore{
		DBFilePath: dbPath,
	}

	// Test
	storedChars, err := db.RetrieveAllCharacters()

	// Validate
	if err != nil {
		t.Fatalf("expected that it would be possible to retrieve the characters recently created: %v", err)
	}
	if len(storedChars) != 2 {
		t.Fatalf("expected that the number of retrieved would match the number of characters recently created: %v", err)
	}
}

func compareCharacters(expected *domain.Character, actual *domain.Character) bool {
	if expected.Name != actual.Name {
		fmt.Printf("characters names do no match - expected %s, but got %s\n", expected.Name, actual.Name)
		return false
	}
	if expected.InitiativeModifier != actual.InitiativeModifier {
		fmt.Printf("characters initiative modifiers do no match - expected %d, but got %d\n", expected.InitiativeModifier, actual.InitiativeModifier)
		return false
	}
	return true
}
