package main

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	var (
		destination        = "Argentina"
		period             = "mañana"
		destinationAverage = "Peru"
	)

	//Leer archivo
	Tickets, errFile := tickets.ReadFile("tickets.csv")

	if errFile != nil {
		panic(errFile)
	}

	//Total por país
	total, err := tickets.GetTicketsByCountry(destination, Tickets)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nEn total se realizaron %v viajes a %v.\n", total, destination)
	}

	//Total por periodo
	//Madrugada, mañana, tarde, noche
	total2, err2 := tickets.GetCountByPeriod(period, Tickets)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("En total se realizaron %v viajes durante la %v.\n", total2, period)
	}

	//Promedio por país
	average, errAverage := tickets.GetAverageDestination(destinationAverage, Tickets)

	if errAverage != nil {
		fmt.Println(errAverage)
	} else {
		fmt.Printf("En promedio el %v%% de los viajes se realizaron a %v.", average, destinationAverage)
	}

}
