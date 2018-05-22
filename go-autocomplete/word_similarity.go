package main

import (
	"math"
)

// similarity computes the similarity between
// 2 words. w1len and w2len are the lengths of the 2
// words for which to compute the similarity.
// ld is the levenshtein distance between the 2 words.
// It returns a value between 0 and 1 that identifies the similarity
// between the 2 words. 2 equal words will have a similarity of 1.
func ComputeSimilarity(w1Len, w2Len, ld int) float64 {
	maxLen := math.Max(float64(w1Len), float64(w2Len))

	return 1.0 - float64(ld)/float64(maxLen)
}
