package main

import (
	"flag"
	"fmt"

	"github.com/mikeconroy/advent-of-code-25/day1"
	"github.com/mikeconroy/advent-of-code-25/day2"
	"github.com/mikeconroy/advent-of-code-25/day3"
	"github.com/mikeconroy/advent-of-code-25/day4"
	"github.com/mikeconroy/advent-of-code-25/day5"
	"github.com/mikeconroy/advent-of-code-25/day6"
	"github.com/mikeconroy/advent-of-code-25/day7"
	"github.com/mikeconroy/advent-of-code-25/day8"
	"github.com/mikeconroy/advent-of-code-25/day9"
)

func main() {
	dayToRun := flag.Int("day", 0, "Which Day to run? Defaults to 0 (all)")
	flag.Parse()

	days := []func() (string, string){
		day1.Run,
		day2.Run,
		day3.Run,
		day4.Run,
		day5.Run,
		day6.Run,
		day7.Run,
		day8.Run,
		day9.Run,
	}

	if *dayToRun == 0 {
		for day, run := range days {
			runDay(day+1, run)
		}
	} else {
		runDay(*dayToRun, days[*dayToRun-1])
	}
}

func runDay(dayNum int, run func() (string, string)) {
	p1, p2 := run()
	fmt.Printf("Day %d\n\tP1: %s\n\tP2: %s\n", dayNum, p1, p2)
}
