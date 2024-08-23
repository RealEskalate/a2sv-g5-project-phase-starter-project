package mocks

import (
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/stretchr/testify/mock"
)

type MockBlogAssistantUsecase struct {
	mock.Mock
}

func (m *MockBlogAssistantUsecase) GenerateBlog(keywords []string, tone, audience string) (map[string]interface{}, domain.Error) {
	ret := m.Called(keywords, tone, audience)

	var r0 map[string]interface{}
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(map[string]interface{})
	}

	var r1 domain.Error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(domain.Error)
	}

	return r0, r1
}

func (m *MockBlogAssistantUsecase) EnhanceBlog(content, command string) (map[string]interface{}, domain.Error) {
	ret := m.Called(content, command)

	var r0 map[string]interface{}
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(map[string]interface{})
	}

	var r1 domain.Error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(domain.Error)
	}

	return r0, r1
}

func (m *MockBlogAssistantUsecase) SuggestBlog(industry string) (map[string]interface{}, domain.Error) {
	ret := m.Called(industry)

	var r0 map[string]interface{}
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(map[string]interface{})
	}

	var r1 domain.Error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(domain.Error)
	}

	return r0, r1
}