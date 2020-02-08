package usecase

import (
	"context"
	"github.com/hobord/golang-poc-rest/domain/entity"
	"github.com/hobord/golang-poc-rest/domain/repository/mocks"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestFooGetByID(t *testing.T) {
	mockRepository := &mocks.FooRepository{}

	fakeID := fake.Sentence()
	fakeTitle := fake.Sentence()
	returnMockEntity := &entity.Foo{
		ID:    fakeID,
		Title: fakeTitle,
	}

	mockRepository.On("GetByID", mock.Anything, mock.Anything).Return(returnMockEntity, nil)

	interactor := CreateFooInteractor(mockRepository)
	result, err := interactor.GetByID(context.TODO(), "1")
	if err != nil {
		assert.NoError(t, err)
	}
	assert.Equal(t, result.ID, fakeID, "The result ID should be:"+fakeID)
	assert.Equal(t, result.Title, fakeTitle, "The result ID should be:"+fakeTitle)
}
