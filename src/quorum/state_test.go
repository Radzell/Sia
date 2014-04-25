package quorum

import (
	"common"
	"common/crypto"
	"testing"
)

// Create a state, check the defaults
func TestCreateState(t *testing.T) {
	// does a state create without errors?
	s, err := CreateState(common.NewZeroNetwork())
	if err != nil {
		t.Fatal(err)
	}

	// check that previousEntropyStage1 is initialized correctly
	var emptyEntropy common.Entropy
	emptyHash, err := crypto.CalculateTruncatedHash(emptyEntropy[:])
	if err != nil {
		t.Fatal(err)
	}
	for i := range s.previousEntropyStage1 {
		if s.previousEntropyStage1[i] != emptyHash {
			t.Error("previousEntropyStage1 initialized incorrectly at index ", i)
		}
	}

	// sanity check the default values
	if s.participantIndex != 255 {
		t.Error("s.participantIndex initialized to ", s.participantIndex)
	}
	if s.currentStep != 1 {
		t.Error("s.currentStep should be initialized to 1!")
	}
	if s.wallets == nil {
		t.Error("s.wallets was not initialized")
	}
}

//
func TestJoinQuorum(t *testing.T) {
	_, err := CreateState(common.NewZeroNetwork())
	if err != nil {
		t.Fatal(err)
	}
}

// test HandleMessage and SetAddress

// check general case, check corner cases, and then do some fuzzing
func TestRandInt(t *testing.T) {
	s, err := CreateState(common.NewZeroNetwork())
	if err != nil {
		t.Fatal(err)
	}

	// check that it works in the vanilla case
	previousEntropy := s.currentEntropy
	randInt, err := s.randInt(0, 5)
	if err != nil {
		t.Fatal(err)
	}
	if randInt < 0 || randInt >= 5 {
		t.Fatal("randInt returned but is not between the bounds")
	}

	// check that s.CurrentEntropy flipped to next value
	if previousEntropy == s.currentEntropy {
		t.Fatal("When calling randInt, s.CurrentEntropy was not changed")
	}

	// check the zero value
	randInt, err = s.randInt(0, 0)
	if err == nil {
		t.Fatal("Randint(0,0) should return a bounds error")
	}

	// fuzzing, skip for short tests
	if testing.Short() {
		t.Skip()
	}

	low := 0
	high := common.QuorumSize
	for i := 0; i < 100000; i++ {
		randInt, err = s.randInt(low, high)
		if err != nil {
			t.Fatal("randInt fuzzing error: ", err)
		}

		if randInt < low || randInt >= high {
			t.Fatal("randInt fuzzing: ", randInt, " produced, expected number between ", low, " and ", high)
		}
	}
}
