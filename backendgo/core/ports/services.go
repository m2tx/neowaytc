package ports

import (
	"github.com/google/uuid"
	"github.com/m2tx/neowaytc/backendgo/core/domain"
)

type IdentificationNumberService interface {
	Get(id uuid.UUID) (domain.IdentificationNumber, error)
	New(number string) (domain.IdentificationNumber, error)
}
