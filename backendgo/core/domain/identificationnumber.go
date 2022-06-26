package domain

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/m2tx/neowaytc/backendgo/core/utils"
)

const (
	CPFPattern  string = `([\d]{3})([\d]{3})([\d]{3})([\d]{2})`
	CNPJPattern string = `([\d]{2})([\d]{3})([\d]{3})([\d]{4})([\d]{2})`
)

var (
	ErrorInvalidIdentificationNumber     = errors.New("invalid-identification-number")
	ErrorInvalidIdentificationNumberCPF  = errors.New("invalid-identification-number-cpf")
	ErrorInvalidIdentificationNumberCNPJ = errors.New("invalid-identification-number-cnpj")
)

type IdentificationNumber struct {
	ID      uuid.UUID `json:"id" gorm:"primaryKey"`
	Number  string    `json:"number"`
	Blocked bool      `json:"blocked"`
}

func NewIdentificationNumber(number string) (*IdentificationNumber, error) {
	identificationNumber := &IdentificationNumber{
		Number: number,
	}
	err := identificationNumber.Prepare()

	if err != nil {
		return nil, err
	}

	return identificationNumber, nil
}

func (identificationNumber *IdentificationNumber) Prepare() error {
	err := identificationNumber.Validate()
	if err != nil {
		return err
	}
	identificationNumber.ID = uuid.New()
	identificationNumber.Blocked = false
	return nil
}

func (identificationNumber *IdentificationNumber) String() string {
	return identificationNumber.Number
}

func (identificationNumber *IdentificationNumber) Validate() error {
	s := identificationNumber.Number
	//TODO - VALIDATE PATTERN
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, "/", "", -1)
	s = strings.Replace(s, "-", "", -1)
	length := len(s)
	if length == 11 {
		return validateCPF(s)
	} else if length == 14 {
		return validateCNPJ(s)
	}
	return ErrorInvalidIdentificationNumber
}

func validateCPF(v string) error {
	isEquals := utils.AllDigitsEquals(v)
	if isEquals {
		return ErrorInvalidIdentificationNumberCPF
	}
	d := 0
	for i := 0; i < 9; i++ {
		d += (utils.ToInt(v[i]) * (10 - i))
	}
	d = 11 - (d % 11)
	if d > 9 {
		d = 0
	}
	if utils.ToInt(v[9]) != d {
		return ErrorInvalidIdentificationNumberCPF
	}
	d *= 2
	for i := 0; i < 9; i++ {
		d += (utils.ToInt(v[i]) * (11 - i))
	}
	d = 11 - (d % 11)
	if d > 9 {
		d = 0
	}
	if utils.ToInt(v[10]) != d {
		return ErrorInvalidIdentificationNumberCPF
	}
	return nil
}

func validateCNPJ(v string) error {
	isEquals := utils.AllDigitsEquals(v)
	if isEquals {
		return ErrorInvalidIdentificationNumberCNPJ
	}
	v = utils.Reverse(v)
	d := 0
	p := 2
	for i := 2; i < 14; i++ {
		d += utils.ToInt(v[i]) * p
		p++
		if p > 9 {
			p = 2
		}
	}
	d = 11 - (d % 11)
	if d > 9 {
		d = 0
	}
	if utils.ToInt(v[1]) != d {
		return ErrorInvalidIdentificationNumberCNPJ
	}
	d *= 2
	p = 3
	for i := 2; i < 14; i++ {
		d += utils.ToInt(v[i]) * p
		p++
		if p > 9 {
			p = 2
		}
	}
	d = 11 - (d % 11)
	if d > 9 {
		d = 0
	}
	if utils.ToInt(v[0]) != d {
		return ErrorInvalidIdentificationNumberCNPJ
	}
	return nil
}
