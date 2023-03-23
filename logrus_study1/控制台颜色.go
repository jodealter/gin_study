package main

import "fmt"

const (
	Block  = 0
	Red    = 1
	Green  = 2
	Yellow = 3
	Blue   = 4
	Purple = 5
	Cyan   = 6
	Gray   = 7
)

func PrintMsg(isBg bool, color int, msg string) {
	if !isBg {
		fmt.Printf("\033[3%dm%s\033[0m", color, msg)
	} else {
		fmt.Printf("\033[4%dm%s\033[0m", color, msg)
	}
}

func main() {
	PrintMsg(true, 2, "hello world")
}
