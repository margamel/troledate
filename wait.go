package main

import (
	"fmt"
	"time"
)

var i int = 1
var running bool = true
var mode int
var timec int

func main() {
	mode = getInput("[1] for time waste | [2] for time calc")
	switch mode {
	case 1:
		trole()
	case 2:
		calc()
	default:
		fmt.Println(mode)
	}
}

func trole() {
	if running == true {
		for {
			timer := time.NewTimer(time.Duration(i) * time.Second)
			<-timer.C
			if i == 1 {
				fmt.Printf("Waiting for %d second \n", i)
			} else {
				fmt.Printf("Waiting for %d seconds \n", i)
			}
			i++
		}
	}
}
func calc() {
	mode = getInput("[1] for 'how long would X iterations take?' | [2] for 'how many iterations would fit inside X seconds'")
	switch mode {
	case 1:

		cint := getInput("How many iterations would you like to calculate?")
		fmt.Print("That would be this many minutes: ")
		fmt.Println(iterations(cint))
		main()
	case 2:
		cint := getInput("How many seconds would you like to calculate?")
		fmt.Print("That would be this many iterations: ")
		fmt.Println(seconds(cint))
		main()
	default:
		fmt.Println(mode)
		main()
	}

}
func getInput(msg string) int {
	input := 0
	fmt.Println(msg)
	n, err := fmt.Scanf("%d\n", &input)
	if err != nil {
		fmt.Println(n, err)
	}
	return input
}

func iterations(amnt int) float64 {
	var final float64
	for T := 0.0; T < float64(amnt+1); T++ {
		final += float64(T)
	}
	return final / 60

}
func seconds(mtime int) int {
	timeCount := 0
	numIterations := 0
	for timeCount < mtime {
		numIterations++
		timeCount += numIterations + 1
	}

	return numIterations
}
