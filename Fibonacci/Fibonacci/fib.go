//The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n). 


package main

import "fmt"

func fib(n uint) uint {
   if n == 0 {
       return 0
   } else if n == 1 {
   return 1
 } else {
   return fib(n-1) + fib(n-2)
 }
}

func main() {
var n uint
fmt.Println("enter a number")
fmt.Scanf("%d", &n)
fmt.Println(fib(n))
}
