package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type MockCacheService struct {
	mock.Mock
}

func (m *MockCacheService) Get(key string) (string, error) {
	args := m.Called(key)
	return args.String(0), args.Error(1)
}

func (m *MockCacheService) Set(key string, value interface{}, expiration time.Duration) error {
	args := m.Called(key, value, expiration)
	return args.Error(0)
}

func (m *MockCacheService) Delete(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func (m *MockCacheService) Increment(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func (m *MockCacheService) Decrement(key string) error {
	args := m.Called(key)
	return args.Error(0)
}