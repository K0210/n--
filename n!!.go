package main

import (
    "fmt"
    "math"
)
var x int
func Double_factorial(x:int) {
  var y = 1  
  if x > 0 {
    var k int
    k = x
    for k >= 1 {
      y *= k
      k -= 2
    }
  }else if x <= 0 && x %2 != 0 {
    var k int
    k= 1
    for k >= x {
      y = y / k
      k = k - 2
    }
  }else if x != 0 {
    fmt.Println("error")
    y = Number.NaN
  }
  return y
}
