package erasure

import (
	"bytes"
	"common"
	"crypto/rand"
	"crypto/sha256"
	"testing"
)

// Basic test for reed-solomon coding, verifies that standard input
// will produce the correct results.
func TestCoding(t *testing.T) {
	k := 100
	m := common.SWARMSIZE - k
	bytesPerSlice := 1024

	// generate a random original file
	numRandomBytes := bytesPerSlice * k
	randomBytes := make([]byte, numRandomBytes)
	rand.Read(randomBytes)

	// get hash of original file
	randomBytesHash := common.Hash(sha256.New(), string(randomBytes))

	// encode original file into a data ring
	ringSlices, err := EncodeRing(k, bytesPerSlice, randomBytes)
	if err != nil {
		t.Fatal(err)
	}

	// verify that first k slices are still original data
	originalDataHash := common.Hash(sha256.New(), string(randomBytes))
	if !(bytes.Equal([]byte(originalDataHash), []byte(randomBytesHash))) {
		t.Fatal("original data was modified after caling EncodeRing!")
	}

	// reduce file to a set of k slices and print those slices out
	remainingSlices := make([]string, k)
	sliceIndicies := make([]uint8, k)
	for i := m; i < common.SWARMSIZE; i++ {
		remainingSlices[i-m] = ringSlices[i]
		sliceIndicies[i-m] = uint8(i)
	}

	// recover original data
	recoveredData, err := RebuildBlock(k, bytesPerSlice, remainingSlices, sliceIndicies)
	if err != nil {
		t.Fatal(err)
	}

	// compare to hash of data when first generated
	recoveredDataHash := common.Hash(sha256.New(), string(recoveredData))
	if !(bytes.Equal([]byte(recoveredDataHash), []byte(randomBytesHash))) {
		t.Fatal("recovered data is different from original data")
	}
}

// At some point, there should be a long test that explores all of the edge cases.

// There should be a fuzzing test that explores random inputs. In particular, I would
// like to fuzz the 'RebuildBlock' function

// There should also be a benchmarking test here.
