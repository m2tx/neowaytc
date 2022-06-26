package services

import (
	"github.com/google/uuid"
	"github.com/m2tx/neowaytc/backendgo/core/domain"
	"github.com/m2tx/neowaytc/backendgo/core/ports"
)

type service struct {
	identificationNumberRepository ports.IdentificationNumberRepository
}

func NewIdentificationNumberService(identificationNumberRepository ports.IdentificationNumberRepository) *service {
	return &service{
		identificationNumberRepository: identificationNumberRepository,
	}
}

func (srv *service) GetAll() []domain.IdentificationNumber {
	return srv.identificationNumberRepository.GetAll()
}

func (srv *service) Get(id uuid.UUID) (domain.IdentificationNumber, error) {
	identificationNumber, err := srv.identificationNumberRepository.Get(id)
	if err != nil {
		return domain.IdentificationNumber{}, ports.ErrorNotFoundIdentificationNumber
	}
	return identificationNumber, nil
}

func (srv *service) New(number string) (domain.IdentificationNumber, error) {
	identificationNumber, err := domain.NewIdentificationNumber(number)
	if err != nil {
		return domain.IdentificationNumber{}, err
	}
	if srv.identificationNumberRepository.ExitsByNumber(identificationNumber.Number) {
		return domain.IdentificationNumber{}, ports.ErrorExitsIdentificationNumber
	}
	err = srv.identificationNumberRepository.Save(*identificationNumber)
	if err != nil {
		return domain.IdentificationNumber{}, err
	}
	return *identificationNumber, nil
}

func (srv *service) Update(identificationNumber domain.IdentificationNumber) error {
	err := identificationNumber.Validate()
	if err != nil {
		return err
	}
	return srv.identificationNumberRepository.Save(identificationNumber)
}
