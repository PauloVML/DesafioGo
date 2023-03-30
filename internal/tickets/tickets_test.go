package tickets

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestGetTicketsByCountry(t *testing.T) {
	//Arrange
	tickets := []Ticket{
		{
			ID:            1,
			nombre:        "Paulo",
			email:         "paulo@gmail.com",
			paisDeDestino: "Argentina",
			horaDelVuelo:  "10:15",
			precio:        1500,
		}, {
			ID:            2,
			nombre:        "Jose",
			email:         "jose@gmail.com",
			paisDeDestino: "Argentina",
			horaDelVuelo:  "15:15",
			precio:        1500,
		}, {
			ID:            3,
			nombre:        "Maria",
			email:         "maria@gmail.com",
			paisDeDestino: "Brazil",
			horaDelVuelo:  "20:20",
			precio:        1500,
		},
	}
	expectedTotal := 2

	//Act
	total, err := GetTicketsByCountry("Argentina", tickets)

	//Assert
	assert.Equal(t, expectedTotal, total)
	assert.Nil(t, err)

}

func TestGetCountByPeriod(t *testing.T) {
	//Arrange
	tickets := []Ticket{

		{
			ID:            1,
			nombre:        "Paulo",
			email:         "paulo@gmail.com",
			paisDeDestino: "Argentina",
			horaDelVuelo:  "10:15",
			precio:        1500,
		}, {
			ID:            2,
			nombre:        "Jose",
			email:         "jose@gmail.com",
			paisDeDestino: "Argentina",
			horaDelVuelo:  "15:15",
			precio:        1500,
		}, {
			ID:            3,
			nombre:        "Maria",
			email:         "maria@gmail.com",
			paisDeDestino: "Brazil",
			horaDelVuelo:  "20:20",
			precio:        1500,
		},
	}
	expected := 1

	//Act
	count, err := GetCountByPeriod("mañana", tickets)

	//Assert
	assert.Equal(t, expected, count)
	assert.Nil(t, err)
}

func TestGetAverageDestination(t *testing.T) {
	//Arrange
	tickets := []Ticket{

		{
			ID:            1,
			nombre:        "Paulo",
			email:         "paulo@gmail.com",
			paisDeDestino: "Argentina",
			horaDelVuelo:  "10:15",
			precio:        1500,
		}, {
			ID:            2,
			nombre:        "Jose",
			email:         "jose@gmail.com",
			paisDeDestino: "Argentina",
			horaDelVuelo:  "15:15",
			precio:        1500,
		}, {
			ID:            3,
			nombre:        "Maria",
			email:         "maria@gmail.com",
			paisDeDestino: "Brazil",
			horaDelVuelo:  "20:20",
			precio:        1500,
		},
	}
	var expected float64 = 67

	//Act
	average, err := GetAverageDestination("Argentina", tickets)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, expected, math.Round(average))

}
