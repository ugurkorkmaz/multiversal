package cosine_similarity

import "math"

/*
The Find function calculates the cosine similarity between two vectors represented as maps of float64 values.

@param vectorOne Map of float64 values representing the first vector.

@param vectorSecond Map of float64 values representing the second vector.

@returns Cosine similarity between two vectors as a float64 value.
*/
func Find(vectorOne, vectorSecond map[string]float64) float64 {
	dotProduct, magnitudeVectorOne, magnitudeVectorSecond := 0.0, 0.0, 0.0

	// Calculate dot product and magnitudes
	for key, value := range vectorOne {
		if value2, ok := vectorSecond[key]; ok {
			dotProduct += value * value2
		}
		magnitudeVectorOne += value * value
	}

	for _, value := range vectorSecond {
		magnitudeVectorSecond += value * value
	}

	// Calculate similarity using cosine similarity formula
	return dotProduct / (math.Sqrt(magnitudeVectorOne) * math.Sqrt(magnitudeVectorSecond))
}
