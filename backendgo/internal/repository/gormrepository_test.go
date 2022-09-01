package repository

import (
	"testing"

	"github.com/google/uuid"
	"github.com/m2tx/neowaytc/backendgo/core/domain"
	"github.com/m2tx/neowaytc/backendgo/core/ports"
	"github.com/stretchr/testify/assert"
)

var (
	rep = *NewIdentificationNumberRepositoryTest([]domain.IdentificationNumber{
		{uuid.MustParse("789c728f-8fa2-494b-8db1-18808a5c61d8"), "046.847.189-80", false},
		{uuid.MustParse("8ccf972c-6f24-4df3-ac65-b94853c10744"), "585.629.410-69", false},
		{uuid.MustParse("35240f60-6a08-4774-becd-826bae221876"), "335.796.160-13", true},
		{uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), "423.590.810-39", true},
	})
)

func TestGetAllIdentificationNumberRepository(t *testing.T) {
	ins := rep.GetAll()
	assert.GreaterOrEqual(t, len(ins), 3)
}

func TestGetIdentificationNumberRepository(t *testing.T) {
	type test struct {
		Name  string
		ID    uuid.UUID
		Error error
	}
	numbers := []test{
		{"Get1", uuid.MustParse("789c728f-8fa2-494b-8db1-18808a5c61d8"), nil},
		{"Get2", uuid.New(), ports.ErrorNotFoundIdentificationNumber},
	}
	for _, nb := range numbers {
		t.Run(nb.Name, func(t *testing.T) {
			_, err := rep.Get(nb.ID)
			assert.Equal(t, nb.Error, err)
		})
	}
}

func TestSaveIdentificationNumberRepository(t *testing.T) {
	type test struct {
		Name                 string
		IdentificationNumber domain.IdentificationNumber
		Error                error
	}
	numbers := []test{
		{"SaveExits1", domain.IdentificationNumber{uuid.MustParse("789c728f-8fa2-494b-8db1-18808a5c61d8"), "046.847.189-80", true}, nil},
		{"SaveNew1", domain.IdentificationNumber{uuid.New(), "133.427.340-51", false}, nil},
		{"SaveExits2", domain.IdentificationNumber{uuid.MustParse("35240f60-6a08-4774-becd-826bae221876"), "335.796.160-13", false}, nil},
		{"SaveNew2", domain.IdentificationNumber{uuid.New(), "767.148.610-87", false}, nil},
	}
	for _, nb := range numbers {
		t.Run(nb.Name, func(t *testing.T) {
			err := rep.Save(nb.IdentificationNumber)
			assert.Equal(t, nb.Error, err)
		})
	}
}

func TestDeleteIdentificationNumberRepository(t *testing.T) {
	type test struct {
		Name                 string
		IdentificationNumber domain.IdentificationNumber
		Error                error
	}
	numbers := []test{
		{"DeleteExits1", domain.IdentificationNumber{uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), "423.590.810-39", true}, nil},
		{"DeleteUnexist1", domain.IdentificationNumber{uuid.New(), "133.427.340-51", false}, ports.ErrorNotFoundIdentificationNumber},
	}
	for _, nb := range numbers {
		t.Run(nb.Name, func(t *testing.T) {
			err := rep.Delete(nb.IdentificationNumber)
			assert.Equal(t, nb.Error, err)
		})
	}
}

func TestExistsByNumberIdentificationNumberRepository(t *testing.T) {
	type test struct {
		Name     string
		Number   string
		Expected bool
	}
	numbers := []test{
		{"Exits1", "046.847.189-80", true},
		{"Exits2", "099.640.840-13", false},
	}
	for _, nb := range numbers {
		t.Run(nb.Name, func(t *testing.T) {
			atual := rep.ExistsByNumber(nb.Number)
			assert.Equal(t, nb.Expected, atual)
		})
	}
}

func TestQueryIdentificationNumberRepository(t *testing.T) {
	params := map[string]any{
		"blocked": false,
	}
	sort, err := domain.ParseSort("id,asc")
	pageable := domain.Pageable{
		Page:     0,
		PageSize: 5,
		Sort:     sort,
	}
	page, err := rep.Query(params, pageable)
	assert.Equal(t, nil, err)
	assert.Equal(t, pageable.PageSize, page.Size)
	assert.GreaterOrEqual(t, len(page.Content), 0)
}
