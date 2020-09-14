# Random Wikipedia
A simple backend service to load and select random wikipedia articles

## Tables
All Tables can also be renamed, by adding a simple Prefix to them.
The Prefix is defined using an environment varialbe and will simply be added to the front of all the names
### Users
`CREATE TABLE IF NOT EXISTS Users (ID TEXT NOT NULL PRIMARY KEY, SessionID TEXT NOT NULL, Email TEXT NOT NULL);`
### Passwords
`CREATE TABLE IF NOT EXISTS Passwords (ID TEXT NOT NULL PRIMARY KEY, Password TEXT NOT NULL, Expiration INTEGER);`
### Notifications
`CREATE TABLE IF NOT EXISTS Notifications (ID TEXT NOT NULL PRIMARY KEY, Subscription TEXT NOT NULL);`
### Settings
`CREATE TABLE IF NOT EXISTS Settings (ID TEXT NOT NULL PRIMARY KEY, Notifytime INTEGER);`
### Userlists
`CREATE TABLE IF NOT EXISTS Userlists (ID TEXT NOT NULL, ListID INTEGER);`
### Lists
`CREATE TABLE IF NOT EXISTS Lists (ListID INTEGER PRIMARY KEY, Title TEXT NOT NULL);`
### Favorites
`CREATE TABLE IF NOT EXISTS Favorites (ID TEXT NOT NULL PRIMARY KEY, ArticleID INTEGER);`
### FavArticles
`CREATE TABLE IF NOT EXISTS FavArticles (ArticleID INTEGER PRIMARY KEY, Title TEXT NOT NULL, URL TEXT NOT NULL);`

## Secrets
### /kv/data/wikipedia/email
* SMTPServer: The Server Domain
* ServerPort: The Port for the server
* Email: The Email used for sending emails
* Password: The Password for the email