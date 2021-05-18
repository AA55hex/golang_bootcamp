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
func setEnvString(env_var *string, env_string string, default_string string) {
	var ok bool
	*env_var, ok = os.LookupEnv(env_string)
	if !ok {
		*env_var = default_string
	}
}

// simple func for setting env variable into int variable
func setEnvInt(env_var *int, env_string string, default_int int) {
	str, ok := os.LookupEnv(env_string)
	buff, err := strconv.ParseInt(str, 10, 32)
	switch {
	case !ok:
		*env_var = default_int
		fmt.Println(env_var, ": set to default.")
		return
	case err != nil:
		*env_var = default_int
		fmt.Println("Bad ", env_var, ": ", buff, ". Set to default.")
	default:
		*env_var = int(buff)
		fmt.Println(env_var, ": ", buff)
	}
	return
}

// Init all env variables from configs.env
func init() {
	fmt.Println("Loading configs.env")
	if err := godotenv.Load("configs.env"); err != nil {
		log.Println("No .env file found: ", err)
	}

	setEnvString(&MySQL.Host, "MYSQL_HOST", "localhost:3306")
	setEnvString(&MySQL.Database, "MYSQL_DB", "")
	setEnvString(&MySQL.User, "MYSQL_USER", "root")
	setEnvString(&MySQL.Password, "MYSQL_PASSOWRD", "")
	setEnvString(&MySQL.MigrationSource, "MIGRATION_SOURCE", "file:///migrations/")
	setEnvString(&Server.Address, "SERVER_ADDRESS", "localhost:8000")

	setEnvInt(&MySQL.ConnectionTryCount, "CONNECTION_TRY_COUNT", 50)
}
