package main

import (
    "fmt"
    "math/rand"
    "time"
    "os"
    "strconv"
)

type Team struct{
  name string
  seed int
}
func main() {
  // winners := []int{}
  // for i:=0; i<34; i++ {
  //   winner := runTournament()
  //   winners = append(winners, winner)
  // }

  // fmt.Println(winners)
  // var total float64 = 0.0
  // for i:=0; i<len(winners); i++ {
  //   total = total + float64(winners[i])
  // }
  // fmt.Println("Mean:", total / float64(len(winners)))

  // oneSeeds := 0
  // for i:=0; i<len(winners); i++ {
  //   if winners[i] == 1 {
  //     oneSeeds++
  //   }
  // }

  // fmt.Println("One seed winners: ", oneSeeds)
  runTournament()
}

func runTournament () int {
  seeds := [][]int{[]int{1, 16}, []int{8, 9}, []int{5, 12}, []int{4, 13}, []int{6, 11}, []int{3, 14}, []int{7, 10}, []int{2, 15}}

  // fmt.Println("")
  fmt.Println("WEST BRACKET")
  westWinner := runBracket(seeds, 1)
  fmt.Println("")
  // fmt.Printf("West Winning Seed: %v\n\n", westWinner)

  fmt.Println("EAST BRACKET")
  eastWinner := runBracket(seeds, 1)
  fmt.Println("")
  // fmt.Printf("East Winning Seed: %v\n\n", eastWinner)

  fmt.Println("SOUTH BRACKET")
  southWinner := runBracket(seeds, 1)
  fmt.Println("")
  // fmt.Printf("South Winning Seed: %v\n\n", southWinner)

  fmt.Println("MIDWEST BRACKET")
  midwestWinner := runBracket(seeds, 1)
  fmt.Println("")
  // fmt.Printf("Midwest Winning Seed: %v\n\n", midwestWinner)


  fmt.Println("FINAL FOUR")
  finalWinner := runBracket([][]int{[]int{westWinner, eastWinner}, []int{southWinner, midwestWinner}}, -1)
  // fmt.Printf("FINALS WINNER: %v\n\n", finalWinner)

  return finalWinner
}

func runBracket (matchups [][]int, round int) int {
  if (round >= 1) {
    fmt.Printf("Round %v\n", round)
  }
  if len(matchups) == 1 {
    finalistSeed1 := matchups[0][0]
    finalistSeed2 := matchups[0][1]
    winnerIndex := determinWinner(finalistSeed1, finalistSeed2)
    sameSeedText := ""
    if finalistSeed1 == finalistSeed2 {
      sameSeedText = "(top team)"
      if winnerIndex == 1 {
        sameSeedText = "(bottom team)"
      }
    }
    fmt.Printf("%v seed%s beats %v seed\n", matchups[0][winnerIndex], sameSeedText, matchups[0][1-winnerIndex])
    return matchups[0][winnerIndex]
  }
  nextSeeds := []int{}
  for i:=0; i<len(matchups); i++ {
    teamOneSeed := matchups[i][0]
    teamTwoSeed := matchups[i][1]
    winningIndex := determinWinner(teamOneSeed, teamTwoSeed)
    sameSeedText := ""
    if teamOneSeed == teamTwoSeed {
      sameSeedText = "(top team)"
      if winningIndex == 1 {
        sameSeedText = "(bottom team)"
      }
    }
    fmt.Printf("%v seed%s beats %v seed\n", matchups[i][winningIndex], sameSeedText, matchups[i][1-winningIndex])
    nextSeeds = append(nextSeeds, matchups[i][winningIndex])
  }
  nextMatchups := [][]int{}
  for i:=0; i<len(nextSeeds); i=i+2 {
    nextMatchups = append(nextMatchups, []int{nextSeeds[i], nextSeeds[i+1]})
  }
  return runBracket(nextMatchups, round+1)
}

func determinWinner(seed1 int, seed2 int) int {
  smallerSeed := seed1
  largerSeed := seed2
  if (seed1 > seed2) {
    smallerSeed = seed2
    largerSeed = seed1
  }
  co := 10
  randCo := 0
  if len(os.Args) > 1 {
    randCo, _ = strconv.Atoi(os.Args[1])
  }
  rand.Seed(time.Now().UnixNano())
  outcome := rand.Intn(smallerSeed*co + largerSeed*co)+1
  if outcome <= smallerSeed*co+randCo {
    // return underdog index
    if seed2 > seed1 {
      return 1
    } else if seed1 > seed2 {
      return 0
    } else {
      return 1
    }
  } else {
    // return favorite index
    if seed2 > seed1 {
      return 0
    } else if seed1 > seed2 {
      return 1
    } else {
      return 0
    }
  }
}
