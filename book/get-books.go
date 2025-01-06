package book

import (
	"context"
	"fmt"
	"time"

	"encore.app/db"
)

type GetBooksQuery struct {
	Page    int    `query:"page"`
	PerPage int    `query:"per_page"`
	From    string `query:"from"`
	To      string `query:"to"`
}

type GetBooksResponse struct {
	Books []db.Book `json:"books"`
}

//encore:api public path=/book method=GET
func (s *Service) GetBooks(ctx context.Context, query *GetBooksQuery) (*GetBooksResponse, error) {

	books := []db.Book{}
	queryBuilder := s.db.
		Select("id, created_at")

	if query.From != "" {
		from, _ := time.Parse(time.RFC3339, query.From)
		fmt.Println(query.To, from.Format(time.RFC3339))
		queryBuilder.Where("created_at>= ?", from)
	}
	if query.To != "" {
		to, _ := time.Parse(time.RFC3339, query.To)
		fmt.Println(query.To, to.Format(time.RFC3339))
		queryBuilder.Where("created_at <= ?", to)
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
