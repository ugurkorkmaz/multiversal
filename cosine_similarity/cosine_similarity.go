package cosine_similarity

import (
	"math"
)

// CosineSimilarity calculates the cosine similarity between two vectors
func CosineSimilarity(vectorOne map[string]float64, vectorSecond map[string]float64) float64 {
	// Create a set of keys by merging keys of both vectors
	vectorKey := make(map[string]bool)

	for key := range vectorOne {
		vectorKey[key] = true
	}

	for key := range vectorSecond {
		vectorKey[key] = true
	}

	// Calculate dot product and magnitudes
	dotProduct := 0.0
	magnitudeVectorOne := 0.0
	magnitudeVectorSecond := 0.0

	for key := range vectorKey {
		keyVectorOneValue := vectorOne[key]
		keyVectorSecondValue := vectorSecond[key]
		dotProduct += (keyVectorOneValue * keyVectorSecondValue)
		magnitudeVectorOne += (keyVectorOneValue * keyVectorOneValue)
		magnitudeVectorSecond += (keyVectorSecondValue * keyVectorSecondValue)
	}

	// Calculate magnitudes by taking square root of the sum of squares of values
	magnitudeVectorOne = math.Sqrt(magnitudeVectorOne)
	magnitudeVectorSecond = math.Sqrt(magnitudeVectorSecond)

	// Calculate similarity using cosine similarity formula
	similarity := dotProduct / (magnitudeVectorOne * magnitudeVectorSecond)

	return similarity
}
