package util

func ArraysMatchInt32(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}

	freq := make(map[int32]int)
	for _, num := range a {
		freq[num]++
	}
	for _, num := range b {
		freq[num]--
		if freq[num] < 0 {
			return false
		}
	}

	return true
}
