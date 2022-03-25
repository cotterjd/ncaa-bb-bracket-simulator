package main

import (
    "fmt"
    "ncaa-bracket/src"
    "os"
    "strconv"
)

func main() {
  var underdogCoefficient int // extra advantage given to every underdog
  var isUpsetWeighted string // wether the below veriable is accounted for
  var upsetBonus int // extra advantage given to every team who has caused an upset in a previous match times the number of upsets
  if len(os.Args) > 1 {
    underdogCoefficient, _ = strconv.Atoi(os.Args[1])
    if len(os.Args) > 2 {
      isUpsetWeighted = os.Args[2]
      if len(os.Args) > 3 {
        upsetBonus, _ = strconv.Atoi(os.Args[2])
      }
    } else {
      fmt.Println(`Do you want an underdog team that causes an upset to have an advantage in there next game? ("yes" or "no". "no" will keep game odds the same regardless of previous wins`)
      fmt.Scanln(&isUpsetWeighted)
    }
  } else {
    fmt.Println(`Pick a number from -60 to 160, -60 being the case that underdogs NEVER win and 160 being the case that underdogs ALWAYS win (Leave blank to keep default value of 0)`)
    fmt.Scanln(&underdogCoefficient)
    fmt.Println(`Do you want an underdog team that causes an upset to have an advantage in there next game? ("yes" or "no". "no" will keep game odds the same regardless of previous wins`)
    fmt.Scanln(&isUpsetWeighted)
  }

  if isUpsetWeighted == `yes` {
    if len(os.Args) <= 3 {
      fmt.Println(`Pick a number between 1 and 145, 1 being the smallest advantage an upsetter has in the the proceeding matches and 145 being the value that would make an upsetter win every proceeding match`)
      fmt.Scanln(&upsetBonus)
    }
    // fmt.Println(`Do you want an underdog team that causes an upset to have an advantage in there next game? ("y" for yes."n" for no to keep all odds the same regardless of previous wins`)
    // var isUpsetWeighted string
    // fmt.Scanln(&isUpsetWeighted)
    src.RunUpsetWeightedTournament(underdogCoefficient, upsetBonus)
  } else {
    src.RunTournament(underdogCoefficient)
  }
}
