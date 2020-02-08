package usecase

import (
	"context"
	"github.com/hobord/golang-poc-rest/domain/entity"
	"github.com/hobord/golang-poc-rest/domain/repository"
)

// FooInteractorInterface is the interface for example use case
// mockery -name=FooInteractorInterface
type FooInteractorInterface interface {
	GetByID(ctx context.Context, id string) (*entity.Foo, error)
	GetAll(ctx context.Context) ([]*entity.Foo, error)
	Save(ctx context.Context, entity *entity.Foo) error
	Delete(ctx context.Context, id string) error
}

// FooInteractor provides an example use-case implementation
type FooInteractor struct {
	FooRepository repository.FooRepository
	// ...Other repositories or interactors
}

// CreateFooInteractor is create a new example "service" / "interactor"
func CreateFooInteractor(repository repository.FooRepository) *FooInteractor {
	return &FooInteractor{
		FooRepository: repository,
	}
}

// GetByID return entity by id
func (i *FooInteractor) GetByID(ctx context.Context, id string) (*entity.Foo, error) {
	return i.FooRepository.GetByID(ctx, id)
}

// GetAll return all entity
func (i *FooInteractor) GetAll(ctx context.Context) ([]*entity.Foo, error) {
	return i.FooRepository.GetAll(ctx)
}

// Save is save to persistent the entity
func (i *FooInteractor) Save(ctx context.Context, entity *entity.Foo) error {
	return i.FooRepository.Save(ctx, entity)
}

// Delete entity from persistent store
func (i *FooInteractor) Delete(ctx context.Context, id string) error {
	return i.FooRepository.Delete(ctx, id)
}
