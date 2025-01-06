package book

import (
	"context"
	"fmt"
	"time"

	"encore.app/db"
)

type CreateBookRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type CreateBookResponse struct {
	Book db.Book `json:"book"`
}

//encore:api public path=/book method=POST
func (s *Service) CreateBook(ctx context.Context, req *CreateBookRequest) (*CreateBookResponse, error) {
	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		return nil, err
	}
	now := time.Now().In(loc)
	fmt.Println("now", now)
	book := db.Book{
		Title:       req.Title,
		Author:      req.Author,
		Description: req.Description,
		CreatedAt:   now,
	}
	gormRes := s.db.Create(&book)
	if gormRes.Error != nil {
		return nil, gormRes.Error
	}
	return &CreateBookResponse{
		Book: book,
	}, nil
}
