package main

// This implementation takes advantage of the 2-columns
// approach as shown in
// https://en.wikipedia.org/wiki/Levenshtein_distance#Iterative_with_two_matrix_rows
//
// You are encouraged to use this function when simply
// interested in the levenshtein distance between 2 words.
func LevenshteinDistance(source, destination string) int {
	vec1 := make([]int, len(destination)+1)
	vec2 := make([]int, len(destination)+1)

	w1 := []rune(source)
	w2 := []rune(destination)

	//initializing vec1
	for i := 0; i < len(vec1); i++ {
		vec1[i] = i
	}

	//initializing matrix
	for i := 0; i < len(w1); i++ {
		for j := 0; j < len(w2); j++ {
			cost := 1
			if w1[i] == w2[j] {
				cost = 0
			}

			min := minimum(vec2[j]+1,
				vec1[j+1]+1,
				vec1[j]+cost)

			vec2[j+1] = min
		}

		for j := 0; j < len(vec1); j++ {
			vec1[j] = vec2[j]
		}
	}
	return vec2[len(w2)]
}
