package entity

import (
	"fmt"
	"github.com/google/uuid"
)

// FooEntity struct definition
type Foo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// NewFooEntity initialize MyEntity
func CreateFooEntity(title string) (*Foo, error) {
	id := uuid.New()
	if title == "" {
		return nil, fmt.Errorf("invalid title")
	}

	return &Foo{
		ID:    id.String(),
		Title: title,
	}, nil
}
