package woahGoHasAQuickCheck

// So we want to make a simple deterministic function to test. The cannonical one is reverse I think.

// ReverseRecursive returns a newly allocated reversed slice computed via recurision
func ReverseRecursive(slc []int) []int {
	if len(slc) == 1 || len(slc) == 0 {
		return slc
	}
	head := slc[0]
	// Commically inefficient.
	return append(ReverseRecursive(slc[1:]), head)
}

// Let's do it iterative just for laughs

// ReverseRecursive returns a newly allocated reversed slice computed via iteration
func ReverseIterative(slc []int) []int {
	if len(slc) == 0 {
		return []int{}
	}

	// savage optimization over here
	sliceLen := len(slc)

	newSlc := make([]int, sliceLen)

	for i := 0; i < sliceLen; i++ {
		temp := slc[sliceLen-1-i]
		newSlc[i] = temp
	}
	return newSlc
}
