package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIdentificationNumber(t *testing.T) {
	type test struct {
		Name   string
		Number string
		Error  error
	}
	numbers := []test{
		{"InvalidNumber", "969022520", ErrorInvalidIdentificationNumber},
		{"InvalidCPFNumber1d1", "969.022.520-19", ErrorInvalidIdentificationNumberCPF},
		{"InvalidCPFNumber1d2", "969.022.520-01", ErrorInvalidIdentificationNumberCPF},
		{"InvalidCPFNumber2d1", "748.414.450-19", ErrorInvalidIdentificationNumberCPF},
		{"InvalidCPFNumber2d2", "748.414.450-91", ErrorInvalidIdentificationNumberCPF},
		{"InvalidCPFNumber3d1", "917.130.980-34", ErrorInvalidIdentificationNumberCPF},
		{"InvalidCPFNumber3d2", "917.130.980-23", ErrorInvalidIdentificationNumberCPF},
		{"InvalidCNPJNumber1d1", "23.391.036/0001-59", ErrorInvalidIdentificationNumberCNPJ},
		{"InvalidCNPJNumber1d2", "23.391.036/0001-44", ErrorInvalidIdentificationNumberCNPJ},
		{"InvalidCNPJNumber2d1", "27.301.665/0001-11", ErrorInvalidIdentificationNumberCNPJ},
		{"InvalidCNPJNumber2d2", "27.301.665/0001-92", ErrorInvalidIdentificationNumberCNPJ},
		{"InvalidCNPJNumber3d1", "02.702.095/0001-07", ErrorInvalidIdentificationNumberCNPJ},
		{"InvalidCNPJNumber3d2", "02.702.095/0001-70", ErrorInvalidIdentificationNumberCNPJ},
		{"InvalidCPFNumberEqualsDigits", "000.000.000-00", ErrorInvalidIdentificationNumberCPF},
		{"InvalidCNPJNumberEqualsDigits", "00.000.000/0000-00", ErrorInvalidIdentificationNumberCNPJ},
		{"ValidCPFNumber1", "969.022.520-09", nil},
		{"ValidCPFNumber2", "748.414.450-99", nil},
		{"ValidCPFNumber3", "917.130.980-24", nil},
		{"ValidCNPJNumber1", "23.391.036/0001-49", nil},
		{"ValidCNPJNumber2", "27.301.665/0001-91", nil},
		{"ValidCNPJNumber3", "02.702.095/0001-10", nil},
	}
	for _, nb := range numbers {
		t.Run(nb.Name, func(t *testing.T) {
			_, err := NewIdentificationNumber(nb.Number)
			assert.Equal(t, nb.Error, err)
		})
	}
}
