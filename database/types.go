package database

import (
	"random_wikipedia/general"

	"github.com/Lol3rrr/sqlvault"
)

// Session is a simple abstraction for all database interactions
type Session interface {
	InsertPassword(ID, Password string, Expiration int64) error
	DeletePassword(ID string) error
	LoadPassword(ID string) (string, int64, error)

	UpdateSessionID(ID, SessionID string) error

	InsertSubscription(ID, subscription string, update bool) error
	InsertSettings(ID string, nSettings general.Settings, update bool) error

	InsertUser(ID, Email string) error
	LoadUserID(ID string) (general.User, error)
	LoadUserSessionID(SessionID string) (general.User, error)
	LoadUsersNotifyTime(notifyTime int) ([]general.User, error)
	InsertUserList(ID string, listID int) error
	RemoveUserList(ID string, listID int) error
	InsertUserFavorite(ID string, ArticleID int, ArticleTitle, ArticleURL string) error
	RemoveUserFavorite(ID string, ArticleID int) error

	InsertList(listID int, title string) error
}

type session struct {
	SQLSession         sqlvault.Session
	Prefix             string
	UsersTable         string
	PasswordsTable     string
	NotificationsTable string
	SettingsTable      string
	UserlistsTable     string
	ListsTable         string
	FavoritesTable     string
	FavArticlesTable   string
}
