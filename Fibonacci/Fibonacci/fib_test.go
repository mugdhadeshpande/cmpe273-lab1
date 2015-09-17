package main

import "fmt"
import "testing"

func TestFib(t *testing.T) {
     var v uint
     v = fib(6)
     if v != 13 {
           t.Error("Expected 1.5, got ", v)
     }
}
