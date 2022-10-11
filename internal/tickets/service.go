package tickets

import (
	domain "desafio-goweb-marioescudero/internal/domains"
)

type Service interface {
	GetTotalTickets() ([]domain.Ticket, error)
	GetTicketDestination(destination string) ([]domain.Ticket, error)
	AverageDestination(destination string) (int, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets() ([]domain.Ticket, error) {
	ps, err := s.repository.GetAll() //llamamos al repository y esperamos su reespuesya
	if err != nil {
		return nil, err //procesamos
	}

	return ps, nil // lo va a recibir el handler
}

func (s *service) GetTicketDestination(destination string) ([]domain.Ticket, error) {
	ps, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return nil, err //procesamos
	}

	return ps, nil // lo va a recibir el handler
}

func (s *service) AverageDestination(destionation string) (int, error) {

	ps, _ := s.repository.GetAverageByDestination(destionation)

	return ps, nil

}
