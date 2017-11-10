package knnAlg

import (
	"testing"
)
import (
	. "mervin.me/knnAlg"
)

func TestKNN(t *testing.T) {
	train := [][]float64{
		[]float64{5.3, 3.7},
		[]float64{5.1, 3.8},
		[]float64{7.2, 3},
		[]float64{5.4, 3.4},
		[]float64{5.1, 3.3},
		[]float64{5.4, 3.9},
		[]float64{7.4, 2.8},
		[]float64{6.1, 2.8},
		[]float64{7.3, 2.9},
		[]float64{6, 2.7},
		[]float64{5.8, 2.8},
		[]float64{6.3, 2.3},
		[]float64{5.1, 2.5},
		[]float64{6.3, 2.5},
		[]float64{5.5, 2.4},
	}

	labels := []string{
		"Setosa",
		"Setosa",
		"Virginica",
		"Setosa",
		"Setosa",
		"Setosa",
		"Virginica",
		"Versicolor",
		"Virginica",
		"Versicolor",
		"Virginica",
		"Versicolor",
		"Versicolor",
		"Versicolor",
		"Versicolor",
	}
	dm := DMT_EulerMethod
	w := []float64{0.5, 0.5}

	var test [][]float64

	test = [][]float64{
		[]float64{5.2, 3.1},
	}
	for i := 1; i < 12; i++ {

		knn := NewKNNClassifier(i, dm, w)
		res := knn.Classify(train, test, labels)
		if res[0] != "Setosa" {
			t.Fatalf("k = %d failed", i)
		}
	}
}
