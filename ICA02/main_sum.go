package main


import "fmt"
import "os"
import "strconv"
//import "./log"
import "bufio"
import "./sum"



func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	var text2 string
	fmt.Println("Skriv inn to tall for å summere, skriv q for å slutte")

	for text != "q" {
		fmt.Print("Tall nummer 1: ")
		scanner.Scan()
		text = scanner.Text()
		fmt.Print("Tall nummer 2: ")
		scanner.Scan()
		text2 = scanner.Text()
		if text != "q" {

			tall1, err := strconv.ParseInt(text,10, 64)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Tallet kan ikke vere: " + text)
			}
			tall2, err := strconv.ParseInt(text2,10, 64)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Basen kan ikke vere: " + text2)
			}


			fmt.Println(text + " + " + text2 + " =", sum.SumInt64(tall1, tall2))

		}
	}
}
