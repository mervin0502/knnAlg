package knnAlg

import (
	"github.com/golang/glog"
)

//SortedDistance
type SortedDistance struct {
	idx  []int
	dist []float64

	cur int
}

//NewSortedDistance initial the SortedDistance with the size
func NewSortedDistance(size int) *SortedDistance {
	glog.Infof("size: %d", size)
	s := &SortedDistance{
		idx:  make([]int, size),
		dist: make([]float64, size),
		cur:  0,
	}
	return s
}

//Put
func (s *SortedDistance) Put(idx int, dist float64) {
	s.idx[s.cur] = idx
	s.dist[s.cur] = dist

	s.cur++
}

//Len
func (s *SortedDistance) Len() int {
	return len(s.dist)
}

//Less return true if [i] < [j].
func (s *SortedDistance) Less(i, j int) bool {
	if s.dist[i] < s.dist[j] {
		return true
	}
	return false
}

//Swap
func (s *SortedDistance) Swap(i, j int) {
	_tempIdx := s.idx[i]
	_tempDistance := s.dist[i]

	s.idx[i] = s.idx[j]
	s.dist[i] = s.dist[j]

	s.idx[j] = _tempIdx
	s.dist[j] = _tempDistance
}

//SelectTopKIdx return the index of the train
func (s *SortedDistance) SelectTopKIdx(k int) []int {
	return s.idx[0:k]
}

//GetIdx
func (s *SortedDistance) GetIdx() []int {
	return s.idx
}
