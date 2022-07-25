package main

import (
	"fmt"
	"log"

	"github.com/douglasandradeee/design-patterns/models"
)

func main() {
	ak47, err := models.GetGun("AK47")
	if err != nil {
		log.Fatalln("Erorr: ", err)
	}
	musket, err := models.GetGun("museket")
	if err != nil {
		log.Fatalln("Erorr: ", err)
	}

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g models.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
