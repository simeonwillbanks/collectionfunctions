## Go by Example: Collection Functions Unit Tests

### Overview
`collectionfunctions` package contains functions from
[Go by Example: Collection Functions](https://gobyexample.com/collection-functions).

As an exercise, unit tests were added to the package.

### Example

```go
package main

import (
    "github.com/simeonwillbanks/collectionfunctions"
    "fmt"
    "strings"
)

func main() {
    var strs = []string{"peach", "apple", "pear", "plum"}

    fmt.Println(collectionfunctions.Index(strs, "pear"))

    fmt.Println(collectionfunctions.Include(strs, "grape"))

    fmt.Println(collectionfunctions.Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(collectionfunctions.All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(collectionfunctions.Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
    }))

    fmt.Println(collectionfunctions.Map(strs, strings.ToUpper))
}
```

**Thanks [Go by Example](https://gobyexample.com)!**
