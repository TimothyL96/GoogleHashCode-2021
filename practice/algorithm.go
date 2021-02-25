package main

import (
	"fmt"
	_ "math/rand"
	_ "sort"
	_ "time"
)

// Main algorithm
//
// To sort p.data ID ascending :
// sort.Slice(p.data, func(i, j int) bool {
// 	return p.data[i].ID < p.data[j].ID
// })
//
func (p *problem) algorithm1() {
	// algorithm
	// Get number of players
	p.nrOfPlayers = 2*p.nrTeamOfTwo + 3*p.nrTeamOfThree + 4*p.nrTeamOfFour

	if p.nrOfPizza >= p.nrOfPlayers {
		fmt.Println("We have enough pizzas to fill all team", p.ID, p.nrOfPizza, p.nrOfPlayers)
	} else {
		fmt.Printf("%s, We are lacking %d number of pizzas\n", p.ID, p.nrOfPlayers-p.nrOfPizza)
	}

	// Record unassigned pizza
	// Unassigned pizza will be used for mutation

	// 1. Create different states

	p.answers = append(p.answers, answer{
		nrOfPeopleInTeam: 2,
		data:             p.getPizzaAndAssign(p.data[:2]),
	})

	p.answers = append(p.answers, answer{
		nrOfPeopleInTeam: 3,
		data:             p.getPizzaAndAssign(p.data[2:5]),
	})

	// Test
	// for _, v := range p.data {
	// 	fmt.Println("{Pizza:}:", v.ID, "assigned:", v.assigned)
	// }

	// Genetic algorithm
	// 2. Calculate fitness

	for k, v := range p.answers {
		fmt.Println("Fitness score for team", k, "is", v.calcFitnessScore())
	}

	// 3. Cross over
	// 4. Mutate
	// 5. Restart
}

func (a *answer) calcFitnessScore() int {
	uniqueIngredients := make(map[string]interface{})

	for _, v := range a.data {
		for _, i := range v.ingredients {
			uniqueIngredients[i] = struct{}{}
		}
	}

	a.fitnessScore = len(uniqueIngredients)

	return a.fitnessScore
}

func (p *problem) getPizzaAndAssign(pData []*problemData) []*problemData {
	for _, v := range pData {
		v.assigned = true
	}

	return pData
}

// Default recursive algorithm
//
func (p *problem) recursive(data, curData []problemData, curPD problemData, maxData []answer, maxScore *int, currentScore int) []answer {
	return maxData
}

// Endless algorithm till max reached or interrupt signalled
func (p *problem) algorithmEndless() {
	// p.algorithm2()
}

// Run recursive algorithm
func (p *problem) algorithmBruteForce() {
	// p.recursive()
}

// Calculate score from input
// Access answer struct with p.answers (type is a slice of answer)
func (p *problem) calcScoreBase(answers []answer) int {
	score := 0

	for _, v := range answers {
		// Create a hash map and record the unique ingredients in a team
		uniqueIngredients := make(map[string]interface{})

		for _, v1 := range v.data {
			for _, i := range v1.ingredients {
				uniqueIngredients[i] = struct{}{}
			}
		}

		fmt.Println("Score for team with people:", v.nrOfPeopleInTeam, "is", len(uniqueIngredients))
		score += len(uniqueIngredients)
	}

	return score
}
