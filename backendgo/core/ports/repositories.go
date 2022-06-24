package ports

import (
	"github.com/google/uuid"
	"github.com/m2tx/neowaytc/backendgo/core/domain"
)

type IdentificationNumberRepository interface {
	Get(id uuid.UUID) (domain.IdentificationNumber, error)
	Save(domain.IdentificationNumber) error
}
