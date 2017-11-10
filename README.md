
## K-nearest neighbors(KNN) algorithm

  the k-nearest neighbors algorithm (k-NN) is a non-parametric method used for classification and regression.[1] In both cases, the input consists of the k closest training examples in the feature space. The output depends on whether k-NN is used for classification 

## Usage
```go
    
    train := [][]float64{}
    labels := []string{}

    k := 3
    dm := DMT_EulerMethod
    w := []float64{}

    var test [][]float64
    test = [][]float64{}

    knn := NewKNNClassifier(k, dm, w)
    res := knn.Classify(train, test, labels)
    log.Println(res)
```

## References:

- https://en.wikipedia.org/wiki/K-nearest_neighbors_algorithm#Dimension_reduction
- http://blog.csdn.net/lsldd/article/details/41357931
- https://www.cnblogs.com/ybjourney/p/4702562.html
