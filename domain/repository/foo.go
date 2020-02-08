package repository

import (
	"context"
	entities "github.com/hobord/golang-poc-rest/domain/entity"
)

// FooRepository interface definition
// make Mock with: mockery -name=FooRepository
type FooRepository interface {
	// Get return entity by id
	GetByID(ctx context.Context, id string) (*entities.Foo, error)

	// GetAll return all FooEntities
	GetAll(ctx context.Context) ([]*entities.Foo, error)

	// Save is save to persistent the Foo
	Save(ctx context.Context, entity *entities.Foo) error

	// Delete Foo from persistent store
	Delete(ctx context.Context, id string) error
}
