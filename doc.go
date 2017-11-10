package knnAlg

// In pattern recognition, the k-Nearest Neighbors algorithm (or k-NN for short)
// is a non-parametric method used for classification and regression.
// [1] In both cases, the input consists of the k closest training examples
// in the feature space. The output depends on whether k-NN is used for classification or regression:
//
// In k-NN classification, the output is a class membership.
// An object is classified by a majority vote of its neighbors,
// with the object being assigned to the class most common among
// its k nearest neighbors (k is a positive integer, typically small).
// If k = 1, then the object is simply assigned to the class of that single nearest neighbor.
//
// In k-NN regression, the output is the property value for the object.
// This value is the average of the values of its k nearest neighbors.
// k-NN is a type of instance-based learning, or lazy learning,
// where the function is only approximated locally and all computation
// is deferred until classification. The k-NN algorithm is among the simplest of all machine learning algorithms.
/*
References:

- https://en.wikipedia.org/wiki/K-nearest_neighbors_algorithm#Dimension_reduction
- http://blog.csdn.net/lsldd/article/details/41357931
- https://www.cnblogs.com/ybjourney/p/4702562.html


*/
