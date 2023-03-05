package cosine_similarity_test

import (
	"testing"

	"github.com/ugurkorkmaz/multiversal/cosine_similarity"
)

func TestCosineSimilarity(t *testing.T) {
	vectorOne := map[string]float64{
		"apple":  3,
		"banana": 2,
		"orange": 0,
		"pear":   1,
	}
	vectorSecond := map[string]float64{
		"apple":  1,
		"banana": 4,
		"orange": 3,
		"pear":   0,
	}
	expectedSimilarity := 0.5765566601970551

	similarity := cosine_similarity.CosineSimilarity(vectorOne, vectorSecond)
	if similarity != expectedSimilarity {
		t.Errorf("Expected similarity: %v, got: %v", expectedSimilarity, similarity)
	}
}
