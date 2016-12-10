package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	var i uint32 = 4
	var d float32 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewReader(os.Stdin)

	var a uint32
	var b float32
	var text string

	fmt.Scanf("%d", &a)
	fmt.Scanf("%f", &b)
	text, _ = scanner.ReadString('\n')




	// Read and save an integer, double, and String to your variables.


	fmt.Println(i + a)
	fmt.Printf("%.1f \n", d + b)
	fmt.Println(s + text)

}
