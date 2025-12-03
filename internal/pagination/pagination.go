package pagination

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"hash"
	"os"

	"github.com/rafaeldepontes/auth-go/internal/domain"
)

// NewOffSetPagination accepts a generic T type, a slice of any data, a current number of pages,
// a total number of records and a size and it will return a OffSetPagination.
func NewOffSetPagination[T any](data []T, currentPage, totalRecords, size uint) *domain.OffSetPagination[T] {
	totalPages := totalRecords / size
	if totalPages <= 0 {
		totalPages = 1
	}

	return &domain.OffSetPagination[T]{
		Data:         data,
		CurrentPage:  currentPage,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
		Size:         size,
	}
}

// NewCursorPagination accepts a generic T type, a slice of any data, a size of records per page and
// the next page being a pointer to the next id in the database and it will return a CursorPagination.
func NewCursorPagination[T any](data []T, size int, nextCursor int64) *domain.CursorPagination[T] {
	return &domain.CursorPagination[T]{
		Data:       data,
		Size:       size,
		NextCursor: nextCursor,
	}
}

// NewCursorHashedPagination accepts a generic T type, a slice of any data, a size of records per page and
// the next page being a pointer to the next id in the database and it will return a CursorHashedPagination
// which is basically a hash with all the information needed in the next request for security.
func NewCursorHashedPagination[T any](data []T, size int, nextCursor int64) (*domain.CursorHashedPagination[T], error) {
	var hashedData string

	rawData := domain.CursorPagination[T]{
		Data:       data,
		Size:       size,
		NextCursor: nextCursor,
	}

	sb, err := json.Marshal(rawData)
	if err != nil {
		return nil, err
	}

	secretKey := os.Getenv("SECRET_CURSOR_KEY")
	var mac hash.Hash = hmac.New(sha256.New, []byte(secretKey))
	mac.Write(sb)
	signature := mac.Sum(nil)

	combined := append(sb, signature...)

	hashedData = base64.RawURLEncoding.EncodeToString(combined)

	return &domain.CursorHashedPagination[T]{
		NextCursor: hashedData,
	}, nil
}
