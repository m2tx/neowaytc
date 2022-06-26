package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/m2tx/neowaytc/backendgo/core/domain"
	"github.com/m2tx/neowaytc/backendgo/core/ports"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pgrep struct {
	db gorm.DB
}

func NewIdentificationNumberPostgresRepository(url string) *pgrep {
	db, _ := gorm.Open(postgres.Open(url), &gorm.Config{})

	db.Table("identification_numbers").AutoMigrate(&domain.IdentificationNumber{})
	return &pgrep{
		db: *db,
	}
}

func (rep *pgrep) GetAll() []domain.IdentificationNumber {
	ins := []domain.IdentificationNumber{}
	rep.db.Find(&ins)
	return ins
}

func (rep *pgrep) Get(id uuid.UUID) (domain.IdentificationNumber, error) {
	identificationNumber := &domain.IdentificationNumber{}
	rep.db.Where(&domain.IdentificationNumber{ID: id}).First(&identificationNumber)
	if identificationNumber.ID == uuid.Nil {
		return domain.IdentificationNumber{}, ports.ErrorNotFoundIdentificationNumber
	}
	return *identificationNumber, nil
}

func (rep *pgrep) Save(identificationNumber domain.IdentificationNumber) error {
	rep.db.Save(&identificationNumber)
	return nil
}

func (rep *pgrep) ExitsByNumber(number string) bool {
	identificationNumber := &domain.IdentificationNumber{}
	rep.db.Where(&domain.IdentificationNumber{Number: number}).First(&identificationNumber)
	return identificationNumber.ID != uuid.Nil
}

func (rep *pgrep) Query(params map[string]string, pageable domain.Pageable) (domain.Page, error) {
	ins := []domain.IdentificationNumber{}
	totalElements := int64(0)
	rep.db.
		Model(domain.IdentificationNumber{}).
		Where(params).
		Count(&totalElements)
	rep.db.
		Offset(pageable.Page * pageable.PageSize).
		Limit(pageable.PageSize).
		Order(fmt.Sprintf("%s %s", pageable.Sort.Active, pageable.Sort.Direction)).
		Model(domain.IdentificationNumber{}).
		Where(params).
		Find(&ins)
	return domain.Page{
		Content:       ins,
		TotalElements: int(totalElements),
		Size:          pageable.PageSize,
	}, nil
}
