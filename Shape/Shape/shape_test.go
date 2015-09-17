package main

import "testing"

func TestPerimeter(t *testing.T) {
     var cp float64
     var cr float64
     cp = perimeter(c)
     if cp != 31.4 {
           t.Error("Expected 31.4, got ", cp)
     }
     cr = perimeter(r)
     if cr != 18 {
          t.Error("Expected 18, got ", cr)
     }
}
