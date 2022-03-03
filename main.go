package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Team struct{
  name string
  seed int
}
func main() {
  seeds := [][]int{[]int{1, 16}, []int{8, 9}, []int{5, 12}, []int{4, 13}, []int{6, 11}, []int{3, 14}, []int{7, 10}, []int{2, 15}}

  fmt.Println("")
  fmt.Println("Playing West Bracket...")
  westWinner := runBracket(seeds)
  fmt.Printf("West Winning Seed: %v\n\n", westWinner)

  fmt.Println("Playing East Bracket...")
  eastWinner := runBracket(seeds)
  fmt.Printf("East Winning Seed: %v\n\n", eastWinner)

  fmt.Println("Playing South Bracket...")
  southWinner := runBracket(seeds)
  fmt.Printf("South Winning Seed: %v\n\n", southWinner)

  fmt.Println("Playing Midwest Bracket...")
  midwestWinner := runBracket(seeds)
  fmt.Printf("Midwest Winning Seed: %v\n\n", midwestWinner)


  finalWinner := runBracket([][]int{[]int{westWinner, eastWinner}, []int{southWinner, midwestWinner}})
  fmt.Printf("FINALS WINNER: %v\n\n", finalWinner)

}

func runBracket (matchups [][]int) int {
  if len(matchups) == 1 {
    winnerIndex := determinWinner(matchups[0][0], matchups[0][1])
    return matchups[0][winnerIndex]
  }
  nextSeeds := []int{}
  for i:=0; i<len(matchups); i++ {
    winningIndex := determinWinner(matchups[i][0], matchups[i][1])
    // fmt.Printf("%v seed beats %v seed\n", matchups[i][winningIndex], matchups[i][1-winningIndex])
    nextSeeds = append(nextSeeds, matchups[i][winningIndex])
  }
  nextMatchups := [][]int{}
  for i:=0; i<len(nextSeeds); i=i+2 {
    nextMatchups = append(nextMatchups, []int{nextSeeds[i], nextSeeds[i+1]})
  }
  return runBracket(nextMatchups)
}

func determinWinner(seed1 int, seed2 int) int {
  rand.Seed(time.Now().UnixNano())
  outcome := rand.Intn(seed1 + seed2)+1
  if outcome <= seed1 {
    return 1
  } else {
    return 0
  }
}
