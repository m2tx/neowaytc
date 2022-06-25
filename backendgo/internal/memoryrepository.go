package internal

import (
	"github.com/google/uuid"
	"github.com/m2tx/neowaytc/backendgo/core/domain"
	"github.com/m2tx/neowaytc/backendgo/core/ports"
)

type repository struct {
	data []domain.IdentificationNumber
}

func NewIdentificationNumberMemoryRepository(data []domain.IdentificationNumber) *repository {
	return &repository{
		data: data,
	}
}

func (rep *repository) Get(id uuid.UUID) (domain.IdentificationNumber, error) {
	for i := 0; i < len(rep.data); i++ {
		if rep.data[i].ID == id {
			return rep.data[i], nil
		}
	}
	return domain.IdentificationNumber{}, ports.ErrorNotFoundIdentificationNumber
}

func (rep *repository) Save(identificationNumber domain.IdentificationNumber) error {
	for i := 0; i < len(rep.data); i++ {
		if rep.data[i].ID == identificationNumber.ID {
			rep.data[i] = identificationNumber
			return nil
		}
	}
	rep.data = append(rep.data, identificationNumber)
	return nil
}

func (rep *repository) ExitsByNumber(number string) bool {
	for i := 0; i < len(rep.data); i++ {
		if rep.data[i].Number == number {
			return true
		}
	}
	return false
}
