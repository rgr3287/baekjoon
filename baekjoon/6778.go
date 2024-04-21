package main

import "fmt"

func main() {
	var Antena, eyes int
	fmt.Scan(&Antena, &eyes)

	if Antena >= 3 && eyes <= 4 {
		fmt.Println("TroyMartian")
	}
	if Antena <= 6 && eyes >= 2 {
		fmt.Println("VladSaturnian")
	}
	if Antena <= 2 && eyes <= 3 {
		fmt.Println("GraemeMercurian")
	}
}
