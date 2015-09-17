

package main

import "fmt"
import "time"


func Sleep(x int) {
    <-time.After(time.Second * time.Duration(x))
}

func main() {
  x:=5
  start := time.Now()
  fmt.Println("zero seconds")
  Sleep (x)
  elapsed := time.Since(start)
  fmt.Println("Print zero seconds took ", elapsed)

}
