package pagination

import "github.com/rafaeldepontes/auth-go/internal/domain"

func NewPagination[T any](data []T, currentPage, totalRecords, size uint) *domain.Pagination[T] {
	totalPages := totalRecords / size
	if totalPages <= 0 {
		totalPages = 1
	}

	return &domain.Pagination[T]{
		Data:         data,
		CurrentPage:  currentPage,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
		Size:         size,
	}
}
