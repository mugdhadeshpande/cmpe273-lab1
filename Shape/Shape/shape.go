package main

import "fmt"
import "math"

type Shape interface {
    perimeter() float64
    }

    type circle struct {
      radius float64
    }
    type rectangle struct {
      length, width float64
    }

func (c circle) perimeter() float64{
  return 2*math.Pi*c.radius
}
func (r rectangle) perimeter() float64{
  return 2*r.length + 2*r.width
}

func perimeter (s Shape) {
  fmt.Println(s.perimeter())
}

func main () {
c:= circle {radius: 5}
r:= rectangle {length: 4, width: 5}
perimeter(c)
perimeter(r)
}
