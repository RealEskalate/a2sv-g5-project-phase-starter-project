package infrastructure_mocks

import(
		"github.com/stretchr/testify/mock"
		// "aait-backend-group4/Infrastructure"
)
type MockOtpService struct {
	mock.Mock
}

func (m *MockOtpService) SendPasswordResetEmail(email, subject, key string) error {
	args := m.Called(email, subject, key)
	return args.Error(0)
}
