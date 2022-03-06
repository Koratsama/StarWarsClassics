package dice

import (
	"fmt"
	"math/rand"
	"time"
)

type Dice struct {
	ChanceCube1 []int
	ChanceCube2 []int
}

/*
Name: CreateSpikeDice
Purpose: The CreateSpikeDice function populates the two
chance cubes with numbers 1-6.
Parameters: none
Returns: none
*/
func (re *Dice) CreateSpikeDice() {
	re.ChanceCube1 = []int{1, 2, 3, 4, 5, 6}
	re.ChanceCube2 = []int{1, 2, 3, 4, 5, 6}
}

/*
Name: RollSpikeDice
Purpose: The RollSpikeDice function picks two random
numbers from each chance cube and returns a boolean
if they are the same or not
Parameters: none
Returns: bool indicating if the dice values match or not
*/
func (re *Dice) RollSpikeDice() bool {

	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))

	var result1 = re.ChanceCube1[rand.Intn(len(re.ChanceCube1))]
	var result2 = re.ChanceCube2[rand.Intn(len(re.ChanceCube2))]
	fmt.Printf("\nSpike dice rolled: %v & %v\n", result1, result2)

	return result1 == result2
}
