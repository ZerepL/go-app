package common

import (
	"encoding/json"
	"github.com/zerepl/go-app/internal/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	dbURL      string
	dbPassword string
	dbUser     string
	appPort    string
)

// SetupENV collect all env vars and setup for the app
func SetupENV() {
	var exists bool

	log.Println("Collecting env vars")
	dbURL, exists = os.LookupEnv("DB_URL")
	if !exists {
		dbURL = "0.0.0.0"
		log.Println("Undeclared DB_URL, using default...")
	}

	dbUser, exists = os.LookupEnv("DB_USER")
	if !exists {
		dbUser = "root"
		log.Println("Undeclared DB_USER, using default...")
	}

	dbPassword, exists = os.LookupEnv("DB_PASS")
	if !exists {
		dbPassword = "123456"
		log.Println("Undeclared DB_PASS, using default...")
	}

	appPort, exists = os.LookupEnv("PORT")
	if !exists {
		appPort = "8080"
		log.Println("Undeclared PORT, using default...")
	}
	log.Printf("DB INFO -> URL: %s | User: %s | Port: %s", dbURL, dbUser, appPort)
}

// GetDBVars return DB info
func GetDBVars() (string, string, string) {
	return dbURL, dbPassword, dbUser
}

// GetPort return app port
func GetPort() string {
	return ":" + appPort
}

func ParseUser(w http.ResponseWriter, r *http.Request) (user *model.User, err error) {
	rawBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	json.Unmarshal([]byte(rawBody), &user)

	w.Write([]byte("Updated\n"))

	return user, nil
}
