package database

import (
	"random_wikipedia/general"

	"github.com/stretchr/testify/mock"
)

// MockDatabase is a simple database struct used for testing
type MockDatabase struct {
	mock.Mock
}

// InsertPassword is needed to comply with the interface
func (m *MockDatabase) InsertPassword(ID, Password string, Expiration int64) error {
	args := m.Called(ID, Password, Expiration)
	return args.Error(0)
}

// DeletePassword is needed to comply with the interface
func (m *MockDatabase) DeletePassword(ID string) error {
	args := m.Called(ID)
	return args.Error(0)
}

// LoadPassword is needed to comply with the interface
func (m *MockDatabase) LoadPassword(ID string) (string, int64, error) {
	args := m.Called(ID)
	return args.String(0), args.Get(1).(int64), args.Error(2)
}

// UpdateSessionID is needed to comply with the interface
func (m *MockDatabase) UpdateSessionID(ID, SessionID string) error {
	args := m.Called(ID, SessionID)
	return args.Error(0)
}

// InsertSubscription is needed to comply with the interface
func (m *MockDatabase) InsertSubscription(ID, subscription string, update bool) error {
	args := m.Called(ID, subscription, update)
	return args.Error(0)
}

// InsertSettings is needed to comply with the interface
func (m *MockDatabase) InsertSettings(ID string, nSettings general.Settings, update bool) error {
	args := m.Called(ID, nSettings, update)
	return args.Error(0)
}

// InsertUser is needed to comply with the interface
func (m *MockDatabase) InsertUser(ID, Email string) error {
	args := m.Called(ID, Email)
	return args.Error(0)
}

// LoadUserID is needed to comply with the interface
func (m *MockDatabase) LoadUserID(ID string) (general.User, error) {
	args := m.Called(ID)
	return args.Get(0).(general.User), args.Error(1)
}

// LoadUserSessionID is needed to comply with the interface
func (m *MockDatabase) LoadUserSessionID(SessionID string) (general.User, error) {
	args := m.Called(SessionID)
	return args.Get(0).(general.User), args.Error(1)
}

// LoadUsersNotifyTime is needed to comply with the interface
func (m *MockDatabase) LoadUsersNotifyTime(notifyTime int) ([]general.User, error) {
	args := m.Called(notifyTime)
	return args.Get(0).([]general.User), args.Error(1)
}

// InsertUserList is needed to comply with the interface
func (m *MockDatabase) InsertUserList(ID string, listID int) error {
	args := m.Called(ID, listID)
	return args.Error(0)
}

// InsertList is needed to comply with the interface
func (m *MockDatabase) InsertList(listID int, title string) error {
	args := m.Called(listID, title)
	return args.Error(0)
}
