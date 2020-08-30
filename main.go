package main

import (
	"fmt"
	"random_wikipedia/api"
	"random_wikipedia/background"
	"random_wikipedia/database"
	"random_wikipedia/general"
	"random_wikipedia/login"
	"random_wikipedia/notifications"
	"random_wikipedia/wikipedia"
	"time"

	"github.com/Lol3rrr/cvault"
	"github.com/Lol3rrr/sqlvault"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

func main() {
	logrus.Infof("Starting...")

	vaultSession, err := cvault.CreateSessionEnv()
	if err != nil {
		logrus.Errorf("Creating Vault Session: %v", err)
		return
	}

	dbHost := general.GetEnvString("DB_HOST", "localhost")
	dbPort := general.GetEnvString("DB_PORT", "5432")
	dbName := general.GetEnvString("DB_NAME", "devWikipedia")
	psqlInfo := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable", dbHost, dbPort, dbName)
	sqlConf := sqlvault.Config{
		DriverName:       "postgres",
		DataSourceName:   psqlInfo,
		VaultCredsPath:   "database/creds/random-wikipedia",
		NewUserThreshold: 10 * time.Second,
	}
	sqlSession, err := sqlvault.Open(sqlConf, vaultSession.GetVaultClient())
	if err != nil {
		logrus.Errorf("Could not create SQL-Session: '%s'", err)
		return
	}

	dbPrefix := general.GetEnvString("DB_PREFIX", "")
	databaseSession := database.NewSession(sqlSession, dbPrefix)

	wikiSession := wikipedia.NewSession([]wikipedia.ListInfo{
		{
			ID:    154126,
			Title: "Unusual Articles",
		},
	}, databaseSession)

	baseURL := general.GetEnvString("BASE_URL", "http://localhost:8080")
	loginSession, err := login.NewSession(baseURL, vaultSession, databaseSession)
	if err != nil {
		logrus.Errorf("Creating Login Session: %v", err)
		return
	}

	subEmail := general.GetEnvString("SUB_EMAIL", "test@example.com")
	notificationSession, err := notifications.NewSession(subEmail, vaultSession)
	if err != nil {
		logrus.Errorf("Creating Notification Session: %v", err)
		return
	}

	// Update wikipedia lists every 3 hours
	go background.UpdateWikipediaLists(wikiSession, 3*time.Hour)
	// Do all the notification sending stuff
	go background.SendNotifications(loginSession, wikiSession, notificationSession, databaseSession)

	apiSession := api.NewAPI(loginSession, notificationSession, wikiSession, databaseSession)
	apiSession.Start(8080)
}
