package persistence

import (
	"github.com/stevejaxon/roll-for-initiative/domain"
)

// CharacterStore is an interface for all D&D initiative character datastores
type CharacterStore interface {
	Create(character *domain.Character) error
	Update(character *domain.Character) error
	Delete(character *domain.Character) error
	RetrieveAllCharacters() ([]domain.Character, error)
}
