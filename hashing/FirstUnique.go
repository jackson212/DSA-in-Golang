func firstUniqChar(s string) int {
	// Step 1: Create a hash map to store the frequency of each character
	freq := make(map[rune]int)

	// Step 2: Count the frequency of each character
	for _, c := range s {
		freq[c]++
	}

	// Step 3: Find the first character with a frequency of 1
	for i, c := range s {
		if freq[c] == 1 {
			return i
		}
	}

	// Step 4: Return -1 if no unique character is found
	return -1
}


