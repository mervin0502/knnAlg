package knnAlg

import (
	"errors"
	"math"

	"github.com/golang/glog"
)

var (
	errNotEqualDataLength = errors.New("The data length is not equal.")
)

//
//EluerDistance
func EulerDistance(X, Y, W []float64) float64 {

	_len := len(X)
	if _len != len(Y) || _len != len(W) {
		glog.Errorln(errNotEqualDataLength)
	}

	var x, y, w float64

	var sum float64
	for i := 0; i < _len; i++ {
		x = X[i]
		y = Y[i]
		w = W[i]
		sum += w * math.Pow(x-y, 2)
	}
	return math.Sqrt(sum)
}

//HuffmanDistance
func HuffmanDistance(X, Y, W []float64) float64 {

	_len := len(X)
	if _len != len(Y) || _len != len(W) {
		glog.Errorln(errNotEqualDataLength)
	}

	var x, y, w float64

	var sum float64
	for i := 0; i < _len; i++ {
		x = X[i]
		y = Y[i]
		w = W[i]
		sum += w * math.Abs(x-y)
	}

	return math.Sqrt(sum)
}
