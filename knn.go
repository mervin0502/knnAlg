package knnAlg

import (
	"github.com/golang/glog"
)

type DistanceMethodType uint8

const (
	//Euler distance method
	DMT_EulerMethod DistanceMethodType = iota
	// Huffman
	DMT_HuffmanMethod
)

//KNNClassifier
type KNNClassifier struct {
	//K define the number of nearest neighbors that will be searched
	K int
	// distance method
	DistanceMethod DistanceMethodType
	//Weight define the weight vector for multi-dimension data
	Weight []float64
}

//NewKNNClassifier
func NewKNNClassifier(k int, dm DistanceMethodType, w []float64) *KNNClassifier {
	knn := &KNNClassifier{
		K:              k,
		DistanceMethod: dm,
		Weight:         w,
	}
	return knn
}

//Classify, train is an nxp matrix, where n denotes the n training data, and p represents the number of attributes. The test is an mxp matrix.
func (k *KNNClassifier) Classify(train, test [][]float64, labels []string) []string {

	// computing the distance between one testing data and all the training data
	// fmt.Println(test)
	// fmt.Println(len(test))
	var sortedDistance *SortedDistance
	result := make([]string, len(test))
	for j, _test := range test {

		sortedDistance = NewSortedDistance(len(train))
		for i, _train := range train {
			var _dist float64
			switch k.DistanceMethod {
			case DMT_EulerMethod:
				_dist = EulerDistance(_test, _train, k.Weight)
				break
			case DMT_HuffmanMethod:
				_dist = HuffmanDistance(_test, _train, k.Weight)
				break
			default:
				_dist = EulerDistance(_test, _train, k.Weight)
			}
			sortedDistance.Put(i, _dist)
		}
		topKIdx := sortedDistance.SelectTopKIdx(k.K)
		//statistic the frequent of every labels
		freqLabels := make(map[string]int, k.K)
		var maxFreq int = 0
		for _, idx := range topKIdx {
			_label := labels[idx]
			v, ok := freqLabels[_label]
			if ok {
				freqLabels[_label] = v + 1
			} else {
				freqLabels[_label] = 1
			}

			if maxFreq < v+1 {
				maxFreq = v + 1
			}

		}
		//get the label with maximum frequent
		_out := make([]string, 0)
		for k, v := range freqLabels {
			if v == maxFreq {
				_out = append(_out, k)
			}
		}
		if len(_out) == 1 {
			result[j] = _out[0]
		} else {
			glog.Warningf("the %d test data has %d classification results.", j, len(_out))
			_str := ""
			for _, v := range _out {
				_str += v + "#"
			}
			result[j] = _str[0:len(_str)]
		}

	}
	return result
}
