package main

import (
	"fmt"

	"github.com/pyKitchen/newmath"
        "github.com/paulsmith/gogeos/geos"
)

func main() {
    line, _ := geos.FromWKT("LINESTRING (0 0, 10 10, 20 20)")
    buf, _ := line.Buffer(2.5)
    fmt.Printf("Hello, world.  Sqrt(2) = %v%v\n", newmath.Sqrt(2), buf)
}
