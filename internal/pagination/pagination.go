package pagination

import "github.com/rafaeldepontes/auth-go/internal/domain"

func NewPagination[T any](data any) *domain.Pagination[T] {
	return &domain.Pagination[T]{
		// Data: data,
	}
}