```go
package main

import "fmt"

func main() {
    var x [5]int
    for i := 0; i < 5; i++ {
        x[i] = i
    }
    fmt.Println(x)
}
```