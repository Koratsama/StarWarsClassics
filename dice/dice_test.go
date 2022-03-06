package dice_test

import (
	"fmt"
	"testing"

	"github.com/Koratsama/StarWarsClassics/dice"
)

func TestCreateSpikeDice(t *testing.T) {
	var spikeDice = dice.Dice{}
	spikeDice.CreateSpikeDice()

	for i := 0; i < len(spikeDice.ChanceCube1); i++ {
		if spikeDice.ChanceCube1[i] != i+1 {
			fmt.Println("Spike dice did not properly initialize.")
			t.Fail()
		}
		if spikeDice.ChanceCube2[i] != i+1 {
			fmt.Println("Spike dice did not properly initialize.")
			t.Fail()
		}
	}
}

func TestRollSpikeDice(t *testing.T) {
	var spikeDice = dice.Dice{}
	spikeDice.CreateSpikeDice()
	var sabaccShift = spikeDice.RollSpikeDice()

	if sabaccShift {
		fmt.Println("Sabacc Shift!!!")
	} else {
		fmt.Println("Dice rolled. no sabacc shift this round.")
	}
}
