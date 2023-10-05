package mocks

import (
	"github.com/stretchr/testify/mock"
	"time"
)

type MockEvent struct {
	mock.Mock
}

func (m *MockEvent) GetName() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *MockEvent) GetDateTime() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}

func (m *MockEvent) GetPayload() interface{} {
	args := m.Called()
	return args.Get(0).(interface{})
}

func (m *MockEvent) SetPayload(payload interface{}) {
	m.Called(payload)
}
