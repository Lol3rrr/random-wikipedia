# Random Wikipedia
A simple backend service to load and select random wikipedia articles

## Tables
### Users
`CREATE TABLE IF NOT EXISTS Users (ID TEXT NOT NULL PRIMARY KEY, SessionID TEXT NOT NULL, Email TEXT NOT NULL);`
### Passwords
`CREATE TABLE IF NOT EXISTS Passwords (ID TEXT NOT NULL PRIMARY KEY, Password TEXT NOT NULL, Expiration INTEGER);`
### Notifications
`CREATE TABLE IF NOT EXISTS Notifications (ID TEXT NOT NULL PRIMARY KEY, Subscription TEXT NOT NULL);`
### Settings
`CREATE TABLE IF NOT EXISTS Settings (ID TEXT NOT NULL PRIMARY KEY, Notifytime TEXT NOT NULL);`
### Userlists
`CREATE TABLE IF NOT EXISTS Userlists (ID TEXT NOT NULL PRIMARY KEY, ListID INTEGER);`
### Lists
`CREATE TABLE IF NOT EXISTS Lists (ListID INTEGER PRIMARY KEY, Title TEXT NOT NULL);`

## Secrets
### /kv/data/wikipedia/email
* SMTPServer: The Server Domain
* ServerPort: The Port for the server
* Email: The Email used for sending emails
* Password: The Password for the email