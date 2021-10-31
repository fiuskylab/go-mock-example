package storage

import (
	"testing"

	"github.com/fiuskylab/go-mock-example/entity"
	"github.com/fiuskylab/go-mock-example/mocks"
	"github.com/google/uuid"
)

func TestUser(t *testing.T) {
	mockStore := &mocks.Store{}

	mockService := &entity.Service{
		Store: mockStore,
	}

	req := entity.UserRequest{
		Username: uuid.NewString(),
		Password: uuid.NewString(),
	}

	m := entity.UserModel{}

	mockStore.
		On("CreateUser", req).
		Return(m, nil)

	mockService.CreateUser(req)

	mockStore.AssertExpectations(t)
}
