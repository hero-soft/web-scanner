package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func New() (*viper.Viper, error) {

	// pflag.String("service_http_port", "7102", "HTTP Service Port")
	// pflag.String("baseURL", "lpadmin-dev.tjx.com/api/config", "URL at which the service can be accessed")
	// pflag.Bool("permissive_headers", true, "Enable debug mode")

	// pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	viper.BindEnv("METRICS_PORT")
	viper.BindEnv("SERVICE_HTTP_PORT")
	viper.BindEnv("SERVICE_GRPC_PORT")
	viper.BindEnv("PERMISSIVE_HEADERS")
	viper.BindEnv("BASE_URL")
	viper.BindEnv("DB_SERVICE")
	viper.BindEnv("KEYCLOAK_URL")
	viper.BindEnv("KEYCLOAK_REALM")

	viper.BindEnv("MONGO_HOST")
	viper.BindEnv("MONGO_DATABASE")
	viper.BindEnv("MONGO_USER")
	viper.BindEnv("MONGO_PASSWORD")
	viper.BindEnv("MONGO_CERT")

	viper.BindEnv("SPACES_ENDPOINT")
	viper.BindEnv("SPACES_KEY")
	viper.BindEnv("SPACES_SECRET")
	viper.BindEnv("SPACES_SPACE")

	// viper.SetConfigFile(".env")

	// Find and read the config file
	// viper.ReadInConfig()

	// if err != nil {
	// 	return nil, fmt.Errorf("could not read ENV settings: %v", err)
	// }

	return viper.GetViper(), nil
}
