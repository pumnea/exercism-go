package dndcharacter

import (
	"math"
	"math/rand"
	"sort"
)

// Character represents a D&D character with ability scores and hit points
type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2))
}

// Ability uses randomness to generate the score for an Ability
// by rolling 4d6 and returns the sum of the three highest dice
func Ability() int {
	result := make([]int, 4)
	for i := range result {
		result[i] = rand.Intn(6) + 1
	}
	// Sort result to easily find the smallest value
	sort.Ints(result)
	// Sum the three highest dice (exclude the lowest at index 0)
	return result[1] + result[2] + result[3]
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	c := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}

	// Hit points are 10 + Constitution modifier
	c.Hitpoints = 10 + Modifier(c.Constitution)

	return c
}
