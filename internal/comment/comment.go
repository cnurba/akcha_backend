package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

type Service struct {
	Store Store
}

// NewService New service instance.
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrivint a comment")

	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, comment Comment) error {
	fmt.Println("update a comment")
	return ErrFetchingComment
}
func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}
func (s *Service) CreateComment(ctx context.Context, comment Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
