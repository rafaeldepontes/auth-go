package domain

type Pagination[T any] struct {
	Data         []T  `json:"data"`
	CurrentPage  uint `json:"currentPage"`
	TotalPages   uint `json:"totalPages"`
	TotalRecords uint `json:"totalRecords"`
	Size         uint `json:"size"`
}
