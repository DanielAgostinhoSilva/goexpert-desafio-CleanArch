package mocks

import (
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/events"
	"github.com/stretchr/testify/mock"
)

type MockEventDispatcher struct {
	mock.Mock
}

func (m *MockEventDispatcher) Register(eventName string, handler events.EventHandler) error {
	args := m.Called(eventName, handler)
	return args.Error(0)
}

func (m *MockEventDispatcher) Dispatch(event events.Event) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockEventDispatcher) Remove(eventName string, handler events.EventHandler) error {
	args := m.Called(eventName, handler)
	return args.Error(0)
}

func (m *MockEventDispatcher) Has(eventName string, handler events.EventHandler) bool {
	args := m.Called(eventName, handler)
	return args.Get(0).(bool)
}

func (m *MockEventDispatcher) Clear() {
	m.Called()
}
