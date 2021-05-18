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
func set_env_string(env_var *string, env_string string, default_string string) {
	var ok bool
	*env_var, ok = os.LookupEnv(env_string)
	if !ok {
		*env_var = default_string
	}
}

// simple func for setting env variable into int variable
func set_env_int(env_var *int, env_string string, default_int int) {
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

	set_env_string(&MySQL.Host, "MYSQL_HOST", "localhost:3306")
	set_env_string(&MySQL.Database, "MYSQL_DB", "")
	set_env_string(&MySQL.User, "MYSQL_USER", "root")
	set_env_string(&MySQL.Password, "MYSQL_PASSOWRD", "")
	set_env_string(&MySQL.MigrationSource, "MIGRATION_SOURCE", "file:///migrations/")
	set_env_string(&Server.Address, "SERVER_ADDRESS", "localhost:8000")

	set_env_int(&MySQL.ConnectionTryCount, "CONNECTION_TRY_COUNT", 50)
}
