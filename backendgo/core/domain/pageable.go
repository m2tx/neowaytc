package domain

import (
	"errors"
	"strings"
)

var (
	ErrorParseSortReceivedBlankValued = errors.New("ParseSort received a blank valued")
)

type Page struct {
	Content       []IdentificationNumber `json:"content"`
	TotalElements int                    `json:"totalElements"`
	Size          int                    `json:"size"`
}

type Pageable struct {
	Page     int
	PageSize int
	Sort     Sort
}

type Sort struct {
	Active    string
	Direction string
}

func ParseSort(value string) (Sort, error) {
	if value == "" {
		return Sort{
			Active:    "id",
			Direction: "asc",
		}, ErrorParseSortReceivedBlankValued
	}
	arg := strings.Split(value, ",")
	sort := Sort{
		Active:    arg[0],
		Direction: arg[1],
	}
	return sort, nil
}
