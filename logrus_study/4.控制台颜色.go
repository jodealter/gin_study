package main

import "fmt"

const (
	Block  = 0
	Red    = 1
	Greed  = 2
	Yellow = 3
	Blue   = 4
	Purple = 5
	Cyan   = 6
	Gray   = 7
)

func ColorPrint(color int, msg string, IsBg bool) {
	if !IsBg {
		fmt.Printf("\033[3%dm%s\033[0m", color, msg)
		return
	}
	fmt.Printf("\033[4%dm%s\033[0m", color, msg)

}
func main() {
	fmt.Println("\033[30mhello\033[0m")
	fmt.Println("\033[31mhello\033[0m")
	fmt.Println("\033[32mhello\033[0m")
	fmt.Println("\033[33mhello\033[0m")
	fmt.Println("\033[34mhello\033[0m")
	fmt.Println("\033[35mhello\033[0m")
	fmt.Println("\033[36mhello\033[0m")
	fmt.Println("\033[37mhello\033[0m")

	fmt.Println("\033[40mhello\033[0m")
	fmt.Println("\033[41mhello\033[0m")
	fmt.Println("\033[42mhello\033[0m")
	fmt.Println("\033[43mhello\033[0m")
	fmt.Println("\033[44mhello\033[0m")
	fmt.Println("\033[45mhello\033[0m")
	fmt.Println("\033[46mhello\033[0m")
	fmt.Println("\033[47mhello\033[0m")

	ColorPrint(1, "你好啊", true)

}
