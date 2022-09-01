package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/m2tx/neowaytc/backendgo/core/domain"
	"github.com/m2tx/neowaytc/backendgo/core/ports"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type identificationNumberRepository struct {
	db *gorm.DB
}

func NewDb(url string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	if url != "" {
		db, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	}
	return db, err
}

func NewIdentificationNumberRepository(db *gorm.DB) ports.IdentificationNumberRepository {
	db.Table("identification_numbers").AutoMigrate(&domain.IdentificationNumber{})
	return &identificationNumberRepository{
		db: db,
	}
}

func NewIdentificationNumberRepositoryTest(data []domain.IdentificationNumber) *ports.IdentificationNumberRepository {
	db, _ := NewDb("")
	rep := NewIdentificationNumberRepository(db)
	for i := 0; i < len(data); i++ {
		rep.Save(data[i])
	}
	return &rep
}

func (rep *identificationNumberRepository) GetAll() []domain.IdentificationNumber {
	ins := []domain.IdentificationNumber{}
	rep.db.Find(&ins)
	return ins
}

func (rep *identificationNumberRepository) Get(id uuid.UUID) (domain.IdentificationNumber, error) {
	identificationNumber := &domain.IdentificationNumber{}
	rep.db.Where(&domain.IdentificationNumber{ID: id}).First(&identificationNumber)
	if identificationNumber.ID == uuid.Nil {
		return domain.IdentificationNumber{}, ports.ErrorNotFoundIdentificationNumber
	}
	return *identificationNumber, nil
}

func (rep *identificationNumberRepository) Save(identificationNumber domain.IdentificationNumber) error {
	rep.db.Save(&identificationNumber)
	return nil
}

func (rep *identificationNumberRepository) Delete(identificationNumber domain.IdentificationNumber) error {
	rowsAffected := rep.db.Delete(&domain.IdentificationNumber{ID: identificationNumber.ID}).RowsAffected
	if rowsAffected == 0 {
		return ports.ErrorNotFoundIdentificationNumber
	}
	return nil
}

func (rep *identificationNumberRepository) ExistsByNumber(number string) bool {
	identificationNumber := &domain.IdentificationNumber{}
	rep.db.Where(&domain.IdentificationNumber{Number: number}).First(&identificationNumber)
	return identificationNumber.ID != uuid.Nil
}

func (rep *identificationNumberRepository) Query(params map[string]any, pageable domain.Pageable) (domain.Page, error) {
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
