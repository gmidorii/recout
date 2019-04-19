package main

type RecoutService interface {
	Create() error
}

type recoutService struct {
}

func NewRecoutService() RecoutService {
	return &recoutService{}
}

func (r *recoutService) Create() error {
	return nil
}
