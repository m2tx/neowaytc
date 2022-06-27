package repository

import (
	"github.com/google/uuid"
	"github.com/m2tx/neowaytc/backendgo/core/domain"
	"github.com/m2tx/neowaytc/backendgo/core/ports"
)

type memory struct {
	data []domain.IdentificationNumber
}

func NewIdentificationNumberMemoryRepository(data []domain.IdentificationNumber) *memory {
	return &memory{
		data: data,
	}
}

func (rep *memory) GetAll() []domain.IdentificationNumber {
	return rep.data
}

func (rep *memory) Get(id uuid.UUID) (domain.IdentificationNumber, error) {
	for i := 0; i < len(rep.data); i++ {
		if rep.data[i].ID == id {
			return rep.data[i], nil
		}
	}
	return domain.IdentificationNumber{}, ports.ErrorNotFoundIdentificationNumber
}

func (rep *memory) Save(identificationNumber domain.IdentificationNumber) error {
	for i := 0; i < len(rep.data); i++ {
		if rep.data[i].ID == identificationNumber.ID {
			rep.data[i] = identificationNumber
			return nil
		}
	}
	rep.data = append(rep.data, identificationNumber)
	return nil
}

func (rep *memory) ExitsByNumber(number string) bool {
	for i := 0; i < len(rep.data); i++ {
		if rep.data[i].Number == number {
			return true
		}
	}
	return false
}

func (rep *memory) Query(params map[string]string, pageable domain.Pageable) (domain.Page, error) {
	content := []domain.IdentificationNumber{}
	for i := 0; i < len(rep.data); i++ {
		cond := true
		if cond {
			content = append(content, rep.data[i])
		}
	}
	return domain.Page{
		Content:       content,
		TotalElements: len(content),
		Size:          10,
	}, nil
}
