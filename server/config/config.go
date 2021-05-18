package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var MySQL struct {
	User               string
	Password           string
	Database           string
	Host               string
	ConnectionTryCount int
	MigrationSource    string
}

var Server struct {
	Address string
}

// simple func for setting env variable into string variable
func set_env(env_var *string, env_string string, default_string string) {
	var ok bool
	*env_var, ok = os.LookupEnv(env_string)
	if !ok {
		*env_var = default_string
	}
}

// Init all env variables from configs.env
func init() {
	fmt.Println("Loading configs.env")
	if err := godotenv.Load("configs.env"); err != nil {
		log.Println("No .env file found: ", err)
	}

	set_env(&MySQL.Host, "MYSQL_HOST", "localhost:3306")
	set_env(&MySQL.Database, "MYSQL_DB", "")
	set_env(&MySQL.User, "MYSQL_USER", "root")
	set_env(&MySQL.Password, "MYSQL_PASSOWRD", "")
	set_env(&MySQL.MigrationSource, "MIGRATION_SOURCE", "file:///migrations/")

	set_env(&Server.Address, "SERVER_ADDRESS", "localhost:8000")

	try_count, _ := os.LookupEnv("CONNECTION_TRY_COUNT")
	ConnectionTryCount, err := strconv.ParseInt(try_count, 10, 32)
	if err != nil || ConnectionTryCount <= 0 {
		ConnectionTryCount = 30
	}
}
