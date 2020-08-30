package database

import "github.com/Lol3rrr/sqlvault"

// NewSession is used to obtain a new Database-Session
func NewSession(sqlSession *sqlvault.DB, tablePrefix string) Session {
	return &session{
		SQLSession:         sqlSession,
		Prefix:             tablePrefix,
		UsersTable:         tablePrefix + "Users",
		PasswordsTable:     tablePrefix + "Passwords",
		NotificationsTable: tablePrefix + "Notifications",
		SettingsTable:      tablePrefix + "Settings",
		UserlistsTable:     tablePrefix + "Userlists",
		ListsTable:         tablePrefix + "Lists",
	}
}
