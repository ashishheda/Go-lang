package main
import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	fmt.Scan(&n)

	if n % 2 == 0 {

		if n >= 2 && n <=5 {

			fmt.Println("Not Weird")

		} else if n >= 6 && n <=20 {

			fmt.Println("Weird")

		} else if n >= 20 {

			fmt.Println("Not Weird")
		}

	} else  {
		fmt.Println( "Weird")
	}
}
