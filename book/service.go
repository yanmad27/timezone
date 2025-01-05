package book

import (
	"context"
	"fmt"
	"time"

	"encore.app/db"
	"gorm.io/gorm"
)

//encore:service
type Service struct {
	db *gorm.DB
}

func initService() (*Service, error) {
	db := db.Init()
	return &Service{db: db}, nil
}

type GetBooksQuery struct {
	Page    int `query:"page"`
	PerPage int `query:"per_page"`
}

type GetBooksResponse struct {
	Books []db.Book `json:"books"`
}

//encore:api public path=/book method=GET
func (s *Service) GetBooks(ctx context.Context, query *GetBooksQuery) (*GetBooksResponse, error) {
	fmt.Println("params", query)

	books := []db.Book{}
	gormRes := s.db.
		// Limit(query.PerPage).
		// Offset(query.Page * query.PerPage).
		Order("created_at desc").
		Find(&books)
	if gormRes.Error != nil {
		return nil, gormRes.Error
	}
	return &GetBooksResponse{
		Books: books,
	}, nil
}

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
