package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
  determinWinner(5, 6)
}

func determinWinner(seed1 int, seed2 int) {
  // diff := math.Abs(seed1 - seed2)
  rand.Seed(time.Now().UnixNano())
  fmt.Println(rand.Intn(seed1 + seed2)+1)
  outcome := rand.Intn(seed1 + seed2)+1
  if outcome <= seed1 {
    fmt.Println("seed 2 wins")
  } else if outcome > seed1 {
    fmt.Println("seed 1 wins")
  }
}
