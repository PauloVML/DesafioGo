package tickets

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID            int
	nombre        string
	email         string
	paisDeDestino string
	horaDelVuelo  string
	precio        int
}

func ReadFile(pathname string) ([]Ticket, error) {
	var Tickets []Ticket

	file, err := os.Open(pathname)

	if err != nil {
		return nil, errors.New("Error al leer el archivo.")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	for _, line := range lines {

		var newTicket Ticket
		newTicket.ID, _ = strconv.Atoi(line[0])
		newTicket.nombre = line[1]
		newTicket.email = line[2]
		newTicket.paisDeDestino = line[3]
		newTicket.horaDelVuelo = line[4]
		newTicket.precio, _ = strconv.Atoi(line[5])

		Tickets = append(Tickets, newTicket)
	}
	return Tickets, nil
}

func GetTicketsByCountry(destination string, tickets []Ticket) (int, error) {
	var cont int = 0
	for _, ticket := range tickets {
		if ticket.paisDeDestino == destination {
			cont++
		}
	}
	if cont == 0 {
		return cont, errors.New("No se encontraron viajes que concuerden con su país")
	}
	return cont, nil
}

func GetCountByPeriod(time string, tickets []Ticket) (int, error) {
	var (
		since int
		until int
		cont  int
	)
	switch time {
	case "madrugada":
		since = 0
		until = 6
	case "mañana":
		since = 7
		until = 12
	case "tarde":
		since = 13
		until = 19
	case "noche":
		since = 20
		until = 23
	default:
		return 0, errors.New("Usted ingresó un período incorrecto")
	}

	for _, ticket := range tickets {
		horario := strings.Split(ticket.horaDelVuelo, ":")
		hora, _ := strconv.Atoi(horario[0])
		if hora > since && hora < until {
			cont++
		}
	}

	return cont, nil
}

func GetAverageDestination(destination string, tickets []Ticket) (float64, error) {
	total, err := GetTicketsByCountry(destination, tickets)

	if err != nil {
		return 0, err
	}

	return float64(total) / float64(len(tickets)) * 100, nil

}
