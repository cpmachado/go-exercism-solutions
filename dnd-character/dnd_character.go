package dndcharacter

import "math/rand"

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

func RollDice() int {
	return rand.Intn(6) + 1
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return score/2 - 5
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	scores := make([]int, 4)
	sum := 0
	least := 6

	for i := 0; i < 4; i++ {
		roll := RollDice()
		sum += roll
		scores[i] = roll
		if roll < least {
			least = roll
		}
	}

	return sum - least
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constitution := Ability()
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(constitution),
	}
}
