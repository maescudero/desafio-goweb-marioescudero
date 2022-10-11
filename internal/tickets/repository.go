package tickets

import (
	domain "desafio-goweb-marioescudero/internal/domains"
	"fmt"
)

/*type Ticket struct {
	Id      string
	Name    string
	Email   string
	Country string
	Time    string
	Price   float64
} */

type Repository interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	GetAverageByDestination(destination string) (int, error)
}

type repository struct {
	db []domain.Ticket
}

func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Ticket, error) {

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return r.db, nil
}

func (r *repository) GetTicketByDestination(destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}

func (r *repository) GetAverageByDestination(destination string) (int, error) {

	//var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return 0, nil
	}

	cantidadDest := 0
	for _, t := range r.db {

		if t.Country == destination {
			cantidadDest++
		}

	}
	promedio := cantidadDest

	return promedio, nil

}
