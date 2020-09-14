CREATE SCHEMA randomWikipedia
  CREATE TABLE Users (ID TEXT NOT NULL PRIMARY KEY, SessionID TEXT NOT NULL, Email TEXT NOT NULL)
  CREATE TABLE Passwords (ID TEXT NOT NULL PRIMARY KEY, Password TEXT NOT NULL, Expiration INTEGER)
  CREATE TABLE Notifications (ID TEXT NOT NULL PRIMARY KEY, Subscription TEXT NOT NULL)
  CREATE TABLE Settings (ID TEXT NOT NULL PRIMARY KEY, Notifytime INTEGER)
  CREATE TABLE Userlists (ID TEXT NOT NULL, ListID INTEGER)
  CREATE TABLE Lists (ListID INTEGER PRIMARY KEY, Title TEXT NOT NULL)
  CREATE TABLE Favorites (ID TEXT NOT NULL PRIMARY KEY, ArticleID INTEGER)
  CREATE TABLE FavArticles (ArticleID INTEGER PRIMARY KEY, Title TEXT NOT NULL, URL TEXT NOT NULL)
;