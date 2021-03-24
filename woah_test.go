package woahGoHasAQuickCheck

import (
	"math/rand"
	"testing"
	"testing/quick"
)

// Keeping this around as a monument to how premature optimization isn't great
// I thought quickcheck would take a while, so I didn't want to call reverse twice to do the check
// But..it's really fast so let's make the code a lil clearer
func MakeTestFn_Old(reverser func([]int) []int) func([]int) bool {
	return func(r []int) bool {
		reversed := reverser(r)

		// fast-track this case, so it doesn't pollute the code below
		if len(r) == 0 {
			return len(reversed) == len(r)
		}
		// fast-track this case since it's wrong and will goof up the loop below
		if len(r) != len(reversed) {
			return false
		}

		forwardIdx, backwardsIdx := 0, len(reversed)-1

		// we only need to trawl half
		for backwardsIdx-forwardIdx != 1 && backwardsIdx-forwardIdx != 0 {
			if r[forwardIdx] != reversed[backwardsIdx] {
				return false
			}
			forwardIdx++
			backwardsIdx--
		}
		return true
	}
}

// The clea{n,r} version!
func MakeTestFn(reverser func([]int) []int) func([]int) bool {
	return func(r []int) bool {
		reversed := reverser(r)
		reReversed := reverser(reversed)
		for i := 0; i < len(r)-1; i++ {
			if r[i] != reReversed[i] {
				return false
			}
		}
		return true
	}
}

func TestRecursiveReverse(t *testing.T) {
	f := MakeTestFn(ReverseRecursive)

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestIterativeReverse(t *testing.T) {
	f := MakeTestFn(ReverseIterative)

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// let's benchmark for laughs
const benchmarkLength = 1 << 20

func makeBenchmarkReversable() []int {
	slc := make([]int, benchmarkLength)
	for i := 0; i < benchmarkLength; i++ {
		slc[i] = rand.Int()
	}
	return slc
}
func BenchmarkRecursiveReverse(b *testing.B) {
	slc := makeBenchmarkReversable()
	b.ResetTimer()
	ReverseRecursive(slc)
}

func BenchmarkIterativeReverse(b *testing.B) {
	slc := makeBenchmarkReversable()
	b.ResetTimer()
	ReverseIterative(slc)
}
