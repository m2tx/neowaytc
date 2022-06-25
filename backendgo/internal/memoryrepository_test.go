package internal

import (
	"testing"

	"github.com/google/uuid"
	"github.com/m2tx/neowaytc/backendgo/core/domain"
	"github.com/m2tx/neowaytc/backendgo/core/ports"
	"github.com/stretchr/testify/assert"
)

var (
	data = []domain.IdentificationNumber{
		{uuid.MustParse("789c728f-8fa2-494b-8db1-18808a5c61d8"), "046.847.189-80", false},
		{uuid.MustParse("8ccf972c-6f24-4df3-ac65-b94853c10744"), "585.629.410-69", false},
		{uuid.MustParse("35240f60-6a08-4774-becd-826bae221876"), "335.796.160-13", true},
	}
	rep = NewIdentificationNumberMemoryRepository(data)
)

func TestGetIdentificationNumberMemoryRepository(t *testing.T) {
	type test struct {
		Name  string
		ID    uuid.UUID
		Error error
	}
	numbers := []test{
		{"GetMemoryRepository1", uuid.MustParse("789c728f-8fa2-494b-8db1-18808a5c61d8"), nil},
		{"GetMemoryRepository2", uuid.New(), ports.ErrorNotFoundIdentificationNumber},
	}
	for _, nb := range numbers {
		t.Run(nb.Name, func(t *testing.T) {
			_, err := rep.Get(nb.ID)
			assert.Equal(t, nb.Error, err)
		})
	}
}

func TestSaveIdentificationNumberMemoryRepository(t *testing.T) {
	type test struct {
		Name                 string
		IdentificationNumber domain.IdentificationNumber
		Error                error
		DataLength           int
	}
	numbers := []test{
		{"SaveExitsIdentificationNumberMemoryRepository1", domain.IdentificationNumber{uuid.MustParse("789c728f-8fa2-494b-8db1-18808a5c61d8"), "046.847.189-80", true}, nil, 3},
		{"SaveNewIdentificationNumberMemoryRepository1", domain.IdentificationNumber{uuid.New(), "133.427.340-51", false}, nil, 4},
		{"SaveExitsIdentificationNumberMemoryRepository2", domain.IdentificationNumber{uuid.MustParse("35240f60-6a08-4774-becd-826bae221876"), "335.796.160-13", false}, nil, 4},
		{"SaveNewIdentificationNumberMemoryRepository2", domain.IdentificationNumber{uuid.New(), "767.148.610-87", false}, nil, 5},
	}
	for _, nb := range numbers {
		t.Run(nb.Name, func(t *testing.T) {
			err := rep.Save(nb.IdentificationNumber)
			assert.Equal(t, nb.Error, err)
			assert.Equal(t, nb.DataLength, len(rep.data))
		})
	}
}

func TestExitsByNumberIdentificationNumberMemoryRepository(t *testing.T) {
	type test struct {
		Name     string
		Number   string
		Expected bool
	}
	numbers := []test{
		{"ExitsMemoryRepository1", "046.847.189-80", true},
		{"ExitsMemoryRepository2", "099.640.840-13", false},
	}
	for _, nb := range numbers {
		t.Run(nb.Name, func(t *testing.T) {
			atual := rep.ExitsByNumber(nb.Number)
			assert.Equal(t, nb.Expected, atual)
		})
	}
}
