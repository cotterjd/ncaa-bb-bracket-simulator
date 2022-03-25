package main

import (
    "fmt"
    "math/rand"
    "time"
    "os"
    "strconv"
)

type Team struct{
  Seed int
  NumUpsets int
  Id int
}

// globals
var gUnderdogAdvantage int
var gUpsetterAdvantage int

func main() {
  if len(os.Args) > 1 {
    gUnderdogAdvantage, _ = strconv.Atoi(os.Args[1])
    if len(os.Args) > 2 {
      gUpsetterAdvantage, _ = strconv.Atoi(os.Args[2])
    } else {
      fmt.Println(`On a scale from 0 to 145 how much of an advantage should a team that won and upset previously in the tournament have in their proceeding matches? With 0 being no extra advantage and 145 meaning they will win no matter what. (leaving blank will result no extra advantage)`)
      fmt.Scanln(&gUpsetterAdvantage)
    }
  } else {
    fmt.Println(`Pick a number from -60 to 160, -60 being the case that underdogs NEVER win and 160 being the case that underdogs ALWAYS win (Leave blank to keep default value of 0)`)
    fmt.Scanln(&gUnderdogAdvantage)
    fmt.Println(`On a scale from 0 to 145 how much of an advantage should a team that won and upset previously in the tournament have in their proceeding matches? With 0 being no extra advantage and 145 meaning they will win no matter what. (leaving blank will result no extra advantage)`)
    fmt.Scanln(&gUpsetterAdvantage)
  }

  runTournament()
}

func runTournament () int {
  matchups := [][]Team{
           []Team{Team{1, 0, 0}, Team{16, 0, 0}},
           []Team{Team{8, 0, 0}, Team{9, 0, 0}},
           []Team{Team{5, 0, 0}, Team{12, 0, 0}},
           []Team{Team{4, 0, 0}, Team{13, 0, 0}},
           []Team{Team{6, 0, 0}, Team{11, 0, 0}},
           []Team{Team{3, 0, 0}, Team{14, 0, 0}},
           []Team{Team{7, 0, 0}, Team{10, 0, 0}},
           []Team{Team{2, 0, 0}, Team{15, 0, 0}}}

  westWinner := runRegion("WEST BRACKET", matchups)
  eastWinner := runRegion("EAST BRACKET", matchups)
  southWinner := runRegion("SOUTH BRACKET", matchups)
  midwestWinner := runRegion("MIDWEST BRACKET", matchups)

  fmt.Println("FINAL FOUR")
  finalWinner := runBracket([][]Team{[]Team{westWinner, eastWinner}, []Team{southWinner, midwestWinner}}, -1)

  if len(os.Args) > 3 {
    if os.Args[3] == `true` {
      fmt.Printf("FINALS WINNER: %v\n\n", finalWinner)
    }
  }

  return westWinner.Seed // finalWinner
}

func runRegion (title string, matchups [][]Team) Team {
  fmt.Println(title)
  regionMatchups := make([][]Team, 8)
  copy(regionMatchups, matchups)
  for i:=0; i<len(regionMatchups); i++ {
    regionMatchups[i][0].Id = rand.Int()
    regionMatchups[i][1].Id = rand.Int()
  }
  winner := runBracket(regionMatchups, 1)
  fmt.Println("")
  return winner
}

func runBracket (matchups [][]Team, round int) Team {
  if (round >= 1) {
    fmt.Printf("Round %v\n", round)
  }
  if len(matchups) == 1 {
    finalist1 := matchups[0][0]
    finalist2 := matchups[0][1]
    winner := determinWinner(finalist1, finalist2)
    logMatch(finalist1, finalist2, winner)
    return winner
  }
  nextTeams := []Team{}
  for i:=0; i<len(matchups); i++ {
    teamOne := matchups[i][0]
    teamTwo := matchups[i][1]
    winner := determinWinner(teamOne, teamTwo)
    logMatch(teamOne, teamTwo, winner)
    nextTeams = append(nextTeams, winner)
  }
  nextMatchups := [][]Team{}
  for i:=0; i<len(nextTeams); i=i+2 {
    nextMatchups = append(nextMatchups, []Team{nextTeams[i], nextTeams[i+1]})
  }
  return runBracket(nextMatchups, round+1)
}

func logMatch (t1 Team, t2 Team, w Team) {
    var winningTeam Team
    var losingTeam Team
    if w.Id == t1.Id {
      winningTeam = t1
      losingTeam = t2
    }
    if w.Id == t2.Id {
      winningTeam = t2
      losingTeam = t1
    }
    sameSeedText := ""
    if winningTeam.Seed == losingTeam.Seed {
      sameSeedText = "(top team)"
      if t2.Id == winningTeam.Id {
        sameSeedText = "(bottom team)"
      }
    }
    fmt.Printf("%v seed%s beats %v seed\n", winningTeam.Seed, sameSeedText, losingTeam.Seed)
}

// TODO: refactor this to return a Team
func determinWinner(team1 Team, team2 Team) Team {
  favorite := team1
  underdog := team2
  if (team1.Seed > team2.Seed) {
    favorite = team2
    underdog = team1
  }
  coefficient := 10 // figure to widen range and allow more granularity in shaping the odds
  rand.Seed(time.Now().UnixNano())
  outcome := rand.Intn(favorite.Seed*coefficient + underdog.Seed*coefficient)+1
  // ex: with 1 seed and 16 seed
  // random number between 1 and 1*10+16*10 = 10+160 = 170 <= 1*10 + 0
  // random number between 1 and 170 <= 10
  if len(os.Args) >= 3 {
    if os.Args[3] == `true` {
      fmt.Printf(`Rand(%v, %v) = %v <= %v   `, 1, favorite.Seed*coefficient + underdog.Seed*coefficient, outcome, favorite.Seed*coefficient+gUnderdogAdvantage+underdog.NumUpsets)
    }
  }
  if outcome <= favorite.Seed*coefficient+gUnderdogAdvantage+underdog.NumUpsets { // the larger this second value is the higher chance an underdog wins
    // underdog wins
    underdog.NumUpsets = underdog.NumUpsets + ((underdog.Seed - favorite.Seed) * gUpsetterAdvantage)
    return underdog
  } else {
    return favorite
  }
}
