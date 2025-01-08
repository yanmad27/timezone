package book

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type BookSuite struct {
	suite.Suite
	bookService *Service
}

func TestBookSuite(t *testing.T) {
	bookService, err := ServiceMod.SafeResolve()
	require.NoError(t, err)
	suite.Run(t, &BookSuite{
		bookService: bookService,
	})

}

func (s *BookSuite) TestGetBooks() {
	t := s.T()
	query := &GetBooksQuery{
		Page:    1,
		PerPage: 10,
		From:    "2024-01-01T00:00:00+07:00",
		To:      "2024-01-10T23:59:59+07:00",
	}
	ctx := context.Background()
	res, err := s.bookService.GetBooks(ctx, query)
	require.NoError(t, err)
	require.Len(t, res.Books, 9)
}
