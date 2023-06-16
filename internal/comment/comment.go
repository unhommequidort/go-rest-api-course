// Package comment defines the service through
// which we interact with comments
package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	// ErrFetchingComment - an error that is returned
	// when we are unable to fetch a comment
	ErrFetchingComment = errors.New("Error fetching comment by ID")

	// ErrNotImplemented - an error that is returned
	// when a method has not been implemented
	ErrNotImplemented = errors.New("Not implemented")
)

// Comment - a representation of a the comment
// structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - defines the methods that our service
// will use to interact with our database
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - is the struct on which all our
// logic will be built on top of
type Service struct {
	store Store
}

// NewService - returns a pointer to a new
// service
func NewService(store Store) *Service {
	return &Service{
		store: store,
	}
}

// GetComment - a function that gets passed an id and
// retrieves the associated comment
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")
	cmt, err := s.store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

// UpdateComment - a function that gets passed a comment
// and updates the associated comment
func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

// DeleteComment - a function that gets passed an id and
// deletes the associated comment
func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

// CreateComment - a function that gets passed a comment
// and creates a new comment
func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
