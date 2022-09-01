package ports

import (
	"errors"

	"github.com/google/uuid"
	"github.com/m2tx/neowaytc/backendgo/core/domain"
)

var (
	ErrorNotFoundIdentificationNumber       = errors.New("Identification number not found")
	ErrorNotFoundIdentificationNumberUpdate = errors.New("Identification number not found to update")
	ErrorExitsIdentificationNumber          = errors.New("Identification number exits")
)

type IdentificationNumberRepository interface {
	GetAll() []domain.IdentificationNumber
	Get(id uuid.UUID) (domain.IdentificationNumber, error)
	Save(identificationNumber domain.IdentificationNumber) error
	Delete(identificationNumber domain.IdentificationNumber) error
	ExistsByNumber(number string) bool
	Query(params map[string]any, pageable domain.Pageable) (domain.Page, error)
}

type IdentificationNumberService interface {
	GetAll() []domain.IdentificationNumber
	Get(id uuid.UUID) (domain.IdentificationNumber, error)
	New(number string) (domain.IdentificationNumber, error)
	Update(identificationNumber domain.IdentificationNumber) error
	Delete(identificationNumber domain.IdentificationNumber) error
	Query(params map[string]any, pageable domain.Pageable) (domain.Page, error)
}
