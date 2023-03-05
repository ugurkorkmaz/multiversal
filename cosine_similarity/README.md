Cosine Similarity
=================

This package provides a function for calculating the cosine similarity between two vectors represented as maps of `float64` values.

Installation
------------

To install this package, run the following command:

`go get github.com/ugurkorkmaz/multiversal`


Usage
-----

Import the package in your Go code:

`import "github.com/ugurkorkmaz/multiversal/cosine_similarity"`

To find the cosine similarity between two vectors, call the `Find` function with two maps of `float64` values:

```go
vectorOne := map[string]float64{
	"a": 1.0,
	"b": 2.0,
	"c": 3.0,
}

vectorSecond := map[string]float64{
	"a": 3.0,
	"b": 2.0,
	"c": 1.0,
}

similarity := cosine_similarity.Find(vectorOne, vectorSecond)
```

The `Find` function returns the cosine similarity between the two vectors as a `float64` value.

Example
-------

```go
package main

import (
	"fmt"
	"github.com/ugurkorkmaz/multiversal/cosine_similarity"
)

func main() {
	vectorOne := map[string]float64{
		"a": 1.0,
		"b": 2.0,
		"c": 3.0,
	}

	vectorSecond := map[string]float64{
		"a": 3.0,
		"b": 2.0,
		"c": 1.0,
	}

	similarity := cosine_similarity.Find(vectorOne, vectorSecond)

	fmt.Println(similarity)
	// Output: 0.7142857142857143
}
```

License
-------

This package is licensed under the MIT License. See the [LICENSE](https://github.com/ugurkorkmaz/multiversal/blob/main/LICENSE) file for details.
