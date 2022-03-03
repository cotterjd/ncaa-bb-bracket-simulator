# ncaa-bb-bracket-simulator
Script built in Go that will simulate a ncaa march madness tournament. 

I've spent hours on brackets that don't do any better than brackets I spend two seconds on. It all seems somewhat random in March Madness. 
So I built this script that will simulate a March Madness tournament to help you fill out a bracket. 

## Getting Started
`sudo apt install -y golang-go`<br>
`git@github.com:cotterjd/ncaa-bb-bracket-simulator.git`<br>
`cd ncaa-bb-bracket-simulator`<br>
`go run main.go`<br>

### command-line arguments
coefficient - pass this argument in to give underdogs more or less of chance of winning (default is 0).<br>
An example of default behavior is if a 1 seed plays a 16 seed the 1 seed has 16 times the chance of winning. If a 5 seed played a 6 seed, then the 5 seed has a 6/11 chance of winning.<br> 
The following will give underdogs a slightly higher chance of winning. 
```
go run main.go 1
```
The following will give underdogs a slightly lower chance of winning.
```
go run main.go -1
```
This value can be any integer. Any value -10 and below will always result in a 1 seed winning the tournament. <br>
A value of 160 or above will always result in a 16 seed winning the tournament. <br>


![bracket2](https://user-images.githubusercontent.com/2576700/156664831-b0fbd444-7e3b-4e4c-9e89-0236c850c217.png)
