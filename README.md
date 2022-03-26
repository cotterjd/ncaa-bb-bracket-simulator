# ncaa-bb-bracket-simulator
Script built in Go that will simulate a ncaa march madness tournament. 

I've spent hours on brackets that don't do any better than brackets I spend two seconds on. It all seems somewhat random in March Madness. 
So I built this script that will simulate a March Madness tournament to help you fill out a bracket quickly with configurable odds. 

## Running Script
`sudo apt install -y golang-go`<br>
`git@github.com:cotterjd/ncaa-bb-bracket-simulator.git`<br>
`cd ncaa-bb-bracket-simulator`<br>
`./bracket.linux`<br>

## Developing
`sudo apt install -y golang-go`<br>
`git@github.com:cotterjd/ncaa-bb-bracket-simulator.git`<br>
`cd ncaa-bb-bracket-simulator`<br>
`vim main.go`<br>
`go run main.go`<br>

### command-line arguments
passing no command-line arguments will prompt questions about how you want to configure the odds. 

You can also pass the configuration in and skip the questions

./bracket.linux &lt;underdog-advantage&gt; &lt;upsett-advantage&gt; &lt;verbose-mode&gt;

For example, `./brack.linux 1 1 true` will give all underdogs a slight relative advantage and also give upsetters a slight relative advantage in the their proceeding matches and turns on verbose mode. Verbose mode will show you the numbers that outcome is based on

If you have verbose mode on then you will see lines like this<br>
`Rand(1, 170) = 32 <= 13    1 seed beats 16 seed`<br>
this means that a random number was generated between 1 and 170 and that number had to be less than or equal to 13 to have the 16 seed upset the 1 seed.
