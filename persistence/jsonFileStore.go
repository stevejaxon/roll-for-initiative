package persistence

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/stevejaxon/roll-for-initiative/domain"
)

// ResultSet represents the retrieve state of the database
type ResultSet struct {
	Characters []domain.Character `json:"characters"`
}

// JSONCharacterStore is an implementation of the CharacterStore interface - mostly used for early prototyping and testing
type JSONCharacterStore struct {
	DBFilePath string
}

// Create adds a character to the database
// Create should not be considered thread safe or more than a testing utility
func (store *JSONCharacterStore) Create(character *domain.Character) error {
	db, isNew, err := store.openDB()
	if err != nil {
		return fmt.Errorf("error creating the character %v: %v", character, err)
	}
	defer db.Close()
	characters := new(ResultSet)
	if !isNew {
		characters, err = store.loadDataFromDB(db)
		if err != nil {
			return fmt.Errorf("error creating the character %v: %v", character, err)
		}
		exists := store.findCharacter(character, characters)
		if exists != nil {
			return fmt.Errorf("cannot create a character that already exists, use update instead")
		}
	}
	characters.Characters = append(characters.Characters, *character)
	return store.storeDataInDB(db, characters)
}

// RetrieveAllCharacters retrieves all of the data from the database
// Due to the small number of expected characters for simplicity we just retrieve them all.
func (store *JSONCharacterStore) RetrieveAllCharacters() ([]domain.Character, error) {
	db, isNew, err := store.openDB()
	if err != nil {
		return nil, fmt.Errorf("error retrieving all of the characters: %v", err)
	}
	defer db.Close()
	if isNew {
		return []domain.Character{}, nil
	}
	characters, err := store.loadDataFromDB(db)
	if err != nil {
		return nil, fmt.Errorf("error retrieving all of the characters: %v", err)
	}
	return characters.Characters, nil
}

func (store *JSONCharacterStore) openDB() (*os.File, bool, error) {
	_, err := os.Stat(store.DBFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			db, err := os.Create(store.DBFilePath)
			if err != nil {
				return nil, false, fmt.Errorf("unable to create the DB file %v", err)
			}
			return db, true, nil
		}
		return nil, false, fmt.Errorf("unable to create the DB file %v", err)
	}
	db, err := os.OpenFile(store.DBFilePath, os.O_RDWR, 0755)
	if err != nil {
		return nil, false, fmt.Errorf("unable to open the DB file %v", err)
	}
	return db, false, nil
}

// We're not expecting there to be many characters in a campaign/database so reading the whole file into memory is fine
func (store *JSONCharacterStore) loadDataFromDB(db *os.File) (*ResultSet, error) {
	byteValue, err := ioutil.ReadAll(db)
	if err != nil {
		return &ResultSet{}, fmt.Errorf("unable to read the characters from the database %v", err)
	}
	var data ResultSet
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return &ResultSet{}, fmt.Errorf("unable to read the characters from the database %v", err)
	}
	return &data, nil
}

func (store *JSONCharacterStore) storeDataInDB(db *os.File, data *ResultSet) error {
	json, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error whilst storing the data in the DB: %v", err)
	}
	_, err = db.WriteAt(json, 0)
	return err
}

func (store *JSONCharacterStore) findCharacter(targetCharacter *domain.Character, data *ResultSet) *domain.Character {
	for _, character := range data.Characters {
		if character.Name == targetCharacter.Name {
			return &character
		}
	}
	return nil
}
