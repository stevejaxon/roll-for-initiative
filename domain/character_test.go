package domain

import (
	"testing"
)

func TestRollInExpectedRange(t *testing.T) {
	character := new(Character)
	for i := 0; i < 1000; i++ {
		roll := character.Roll()
		if roll < RollMin || roll > RollMax {
			t.Fatalf("Expected the roll to be within the range [%d, %d], but was %d", RollMin, RollMax, roll)
		}
	}
}
