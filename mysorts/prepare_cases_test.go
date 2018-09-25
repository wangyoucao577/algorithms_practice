package mysorts

import "math/rand"

func generateRandomCase(maxArrayLen int) intSlice {
	arrayLen := rand.Intn(maxArrayLen)

	result := make([]int, arrayLen, arrayLen)
	for i := 0; i < arrayLen; i++ {
		result[i] = rand.Intn(maxArrayLen)

		//TODO: generate negative numbers
	}
	return result
}

func generateBestCase(maxArrayLen int) intSlice {
	bestCase := make([]int, maxArrayLen, maxArrayLen)
	for i := 0; i < maxArrayLen; i++ {
		bestCase[i] = i // ascending sorted
	}
	return bestCase
}

func generateWorstCase(maxArrayLen int) intSlice {
	worstCase := make([]int, maxArrayLen, maxArrayLen)
	for i := 0; i < maxArrayLen; i++ {
		worstCase[i] = maxArrayLen - i // descending sorted by expect ascending
	}
	return worstCase
}
