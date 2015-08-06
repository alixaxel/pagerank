# pagerank [![GoDoc](https://godoc.org/github.com/alixaxel/pagerank?status.svg)](https://godoc.org/github.com/alixaxel/pagerank)

Weighted PageRank implementation in Go

## Usage

```go
package main

import (
	"fmt"

	"github.com/alixaxel/pagerank"
)

func main() {
	graph := pagerank.NewGraph()

	graph.Link(1, 2, 1.0)
	graph.Link(1, 3, 2.0)
	graph.Link(2, 3, 3.0)
	graph.Link(2, 4, 4.0)
	graph.Link(3, 1, 5.0)

	graph.Rank(0.85, 0.000001, func(id int, rank float64) {
		fmt.Println("Node", id, "has a rank of", rank)
	})
}
```

## Output

```
Node 1 has a rank of 0.34983779905464363
Node 2 has a rank of 0.1688733284604475
Node 3 has a rank of 0.3295121849483849
Node 4 has a rank of 0.15177668753652385
```

## Install

	go get github.com/alixaxel/pagerank

## License

MIT
