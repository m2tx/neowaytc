package ports

import (
	"errors"

	"github.com/google/uuid"
	"github.com/m2tx/neowaytc/backendgo/core/domain"
)

var (
	ErrorNotFoundIdentificationNumber       = errors.New("identification-number-not-found")
	ErrorNotFoundIdentificationNumberUpdate = errors.New("identification-number-not-found-to-update")
	ErrorExitsIdentificationNumber          = errors.New("identification-number-exits")
)

type IdentificationNumberRepository interface {
	Get(id uuid.UUID) (domain.IdentificationNumber, error)
	Save(identificationNumber domain.IdentificationNumber) error
	ExitsByNumber(number string) bool
}

type IdentificationNumberService interface {
	Get(id uuid.UUID) (domain.IdentificationNumber, error)
	New(number string) (domain.IdentificationNumber, error)
	Update(identificationNumber domain.IdentificationNumber) error
}