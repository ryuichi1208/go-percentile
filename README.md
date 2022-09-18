# go-percentile

Libraries that display percentiles

### Install & Usage

```
$ go get github.com/ryuichi1208/go-percentile@v0.0.4
```

``` go
package main

import (
	"fmt"
	"os"

	"github.com/ryuichi1208/go-percentile"
)

func main() {
	var t []int64
	for i := 1; i < 101; i++ {
		t = append(t, int64(i))

	}
	v, err := percentile.PercentileN(t, 90)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(v)
}
```

```
$ go run ./main.go
90
```
