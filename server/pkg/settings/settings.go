package settings

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func New() (*viper.Viper, error) {

	viper.SetDefault("service_http_port", "8080")
	viper.SetDefault("metrics_port", "9090")
	viper.SetDefault("permissive_headers", true)
	viper.SetDefault("talkgroups_file", "")
	viper.SetDefault("console_logs", false)
	// viper.SetDefault("LayoutDir", "layouts")
	// viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	viper.SetConfigName("settings")           // name of config file (without extension)
	viper.SetConfigType("toml")               // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/web-scanner/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.web-scanner") // call multiple times to add many search paths
	viper.AddConfigPath(".")                  // optionally look for config in the working directory
	err := viper.ReadInConfig()               // Find and read the config file
	if err != nil {                           // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	b, err := json.MarshalIndent(viper.Sub("client").AllSettings(), "", "  ")

	if err != nil {
		return nil, err
	}

	err = os.WriteFile("config.json", b, 0644)

	if err != nil {
		return nil, err
	}

	// viper.SetConfigFile(".env")

	// Find and read the config file
	// viper.ReadInConfig()

	// if err != nil {
	// 	return nil, fmt.Errorf("could not read ENV settings: %v", err)
	// }

	return viper.GetViper(), nil
}
