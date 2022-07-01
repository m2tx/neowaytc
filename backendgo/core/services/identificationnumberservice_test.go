package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/m2tx/neowaytc/backendgo/core/domain"
	"github.com/m2tx/neowaytc/backendgo/core/ports"
	"github.com/m2tx/neowaytc/backendgo/internal/repository"
	"github.com/stretchr/testify/assert"
)

var (
	srv = NewIdentificationNumberService(*repository.NewIdentificationNumberRepositoryTest([]domain.IdentificationNumber{
		{uuid.MustParse("789c728f-8fa2-494b-8db1-18808a5c61d8"), "046.847.189-80", false},
		{uuid.MustParse("8ccf972c-6f24-4df3-ac65-b94853c10744"), "585.629.410-69", false},
		{uuid.MustParse("35240f60-6a08-4774-becd-826bae221876"), "335.796.160-13", true},
	}))
)

func TestGetAllIdentificationNumberService(t *testing.T) {
	ins := srv.GetAll()
	assert.Equal(t, len(ins), 3)
}

func TestGetIdentificationNumberService(t *testing.T) {
	type test struct {
		Name  string
		ID    uuid.UUID
		Error error
	}
	numbers := []test{
		{"GetService1", uuid.MustParse("789c728f-8fa2-494b-8db1-18808a5c61d8"), nil},
		{"GetService2", uuid.New(), ports.ErrorNotFoundIdentificationNumber},
	}
	for _, nb := range numbers {
		t.Run(nb.Name, func(t *testing.T) {
			in, err := srv.Get(nb.ID)
			assert.Equal(t, nb.Error, err)
			if err == nil {
				assert.Equal(t, nb.ID, in.ID)
			}
		})
	}
}

func TestNewIdentificationNumberService(t *testing.T) {
	type test struct {
		Name   string
		Number string
		Error  error
	}
	numbers := []test{
		{"InvalidNumberService", "969022520", domain.ErrorInvalidIdentificationNumber},
		{"InvalidCPFNumber1d1Service", "969.022.520-19", domain.ErrorInvalidIdentificationNumberCPF},
		{"InvalidCPFNumber1d2Service", "969.022.520-01", domain.ErrorInvalidIdentificationNumberCPF},
		{"InvalidCPFNumber2d1Service", "748.414.450-19", domain.ErrorInvalidIdentificationNumberCPF},
		{"InvalidCPFNumber2d2Service", "748.414.450-91", domain.ErrorInvalidIdentificationNumberCPF},
		{"InvalidCPFNumber3d1Service", "917.130.980-34", domain.ErrorInvalidIdentificationNumberCPF},
		{"InvalidCPFNumber3d2Service", "917.130.980-23", domain.ErrorInvalidIdentificationNumberCPF},
		{"InvalidCNPJNumber1d1Service", "23.391.036/0001-59", domain.ErrorInvalidIdentificationNumberCNPJ},
		{"InvalidCNPJNumber1d2Service", "23.391.036/0001-44", domain.ErrorInvalidIdentificationNumberCNPJ},
		{"InvalidCNPJNumber2d1Service", "27.301.665/0001-11", domain.ErrorInvalidIdentificationNumberCNPJ},
		{"InvalidCNPJNumber2d2Service", "27.301.665/0001-92", domain.ErrorInvalidIdentificationNumberCNPJ},
		{"InvalidCNPJNumber3d1Service", "02.702.095/0001-07", domain.ErrorInvalidIdentificationNumberCNPJ},
		{"InvalidCNPJNumber3d2Service", "02.702.095/0001-70", domain.ErrorInvalidIdentificationNumberCNPJ},
		{"InvalidCPFNumberEqualsDigitsService", "000.000.000-00", domain.ErrorInvalidIdentificationNumberCPF},
		{"InvalidCNPJNumberEqualsDigitsService", "00.000.000/0000-00", domain.ErrorInvalidIdentificationNumberCNPJ},
		{"ValidCPFNumber1Service", "969.022.520-09", nil},
		{"ValidCPFNumber2Service", "748.414.450-99", nil},
		{"ValidCPFNumber3Service", "917.130.980-24", nil},
		{"ValidCNPJNumber1Service", "23.391.036/0001-49", nil},
		{"ValidCNPJNumber2Service", "27.301.665/0001-91", nil},
		{"ValidCNPJNumber3Service", "02.702.095/0001-10", nil},
	}
	for _, nb := range numbers {
		t.Run(nb.Name, func(t *testing.T) {
			_, err := srv.New(nb.Number)
			assert.Equal(t, nb.Error, err)
		})
	}
}

func TestUpdateIdentificationNumberService(t *testing.T) {
	type test struct {
		Name                 string
		IdentificationNumber domain.IdentificationNumber
		Error                error
	}
	numbers := []test{
		{"UpdateIdentificationNumberService1", domain.IdentificationNumber{uuid.MustParse("789c728f-8fa2-494b-8db1-18808a5c61d8"), "046.847.189-80", true}, nil},
		{"UpdateIdentificationNumberService2", domain.IdentificationNumber{uuid.MustParse("35240f60-6a08-4774-becd-826bae221876"), "335.796.160-13", false}, nil},
		{"UpdateIdentificationNumberService2Invalid", domain.IdentificationNumber{uuid.MustParse("35240f60-6a08-4774-becd-826bae221876"), "335.796.160-12", false}, domain.ErrorInvalidIdentificationNumberCPF},
	}
	for _, nb := range numbers {
		t.Run(nb.Name, func(t *testing.T) {
			err := srv.Update(nb.IdentificationNumber)
			assert.Equal(t, nb.Error, err)
		})
	}
}
