package utils

import (
	"errors"
	"fmt"
	"math"
)

type MetaData struct {
	CurrentPage  int
	perPage      int
	TotalPages   int
	TotalRecords int
}

func ComputeMetaData(CurrentPage, perPage, totalRecords int) (MetaData, int, error) {

	//calculating offset

	offset := (CurrentPage - 1) * perPage
	//calculating total pages

	//totalPages := int(math.Ceil(float64(totalRecords) / float64(perPage)))
	totalPages := math.Ceil(float64(totalRecords) / float64(perPage))

	fmt.Println(perPage)

	if totalPages == 0 || CurrentPage > int(totalPages) {
		return MetaData{}, -1, errors.New(" no records ")
	}
	return MetaData{
		CurrentPage:  CurrentPage,
		perPage:      perPage,
		TotalPages:   int(totalPages),
		TotalRecords: totalRecords,
	}, offset, nil
}
