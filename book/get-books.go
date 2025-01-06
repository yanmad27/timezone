package book

import (
	"context"
	"time"

	"encore.app/db"
)

type GetBooksQuery struct {
	Page    int   `query:"page"`
	PerPage int   `query:"per_page"`
	From    int64 `query:"from"`
	To      int64 `query:"to"`
}

type GetBooksResponse struct {
	Books []db.Book `json:"books"`
}

//encore:api public path=/book method=GET
func (s *Service) GetBooks(ctx context.Context, query *GetBooksQuery) (*GetBooksResponse, error) {

	books := []db.Book{}
	queryBuilder := s.db.
		Select("id, created_at")

	if query.From != 0 {
		queryBuilder.Where("created_at>= ?", time.Unix(query.From, 0))
	}
	if query.To != 0 {
		queryBuilder.Where("created_at <= ?", time.Unix(query.To, 0))
	}

	queryBuilder.Limit(query.PerPage).
		Offset((query.Page - 1) * query.PerPage).
		Order("created_at asc")

	gormRes := queryBuilder.Find(&books)
	if gormRes.Error != nil {
		return nil, gormRes.Error
	}
	return &GetBooksResponse{
		Books: books,
	}, nil
}
