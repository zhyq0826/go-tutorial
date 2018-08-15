package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("CONSUL_ADDRESS", "localhost:8500")
	viper.SetDefault("DATABASE_USER", "root")
	viper.SetDefault("DATABASE_PASSWD", "")
	viper.SetDefault("LOG_SERVER_PORT", 8081)
	viper.SetDefault("HTTP_SERVER_PORT", 8000)
	viper.SetDefault("GITLAB_TOKEN", "")
	viper.SetDefault("LOG_PATH", "devops.log")

	viper.SetConfigName("config")         // name of config file (without extension)
	viper.AddConfigPath("/etc/")          // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory

	consulAddress := viper.Get("CONSUL_ADDRESS")
	fmt.Println(consulAddress)
	databaseUser := viper.Get("DATABASE_USER")
	fmt.Println(databaseUser)
	databassPasswd := viper.Get("DATABASE_PASSWD")
	fmt.Println(databassPasswd)
	logServerPort := viper.Get("LOG_SERVER_PORT")
	fmt.Println(logServerPort)
	httpServerPort := viper.Get("HTTP_SERVER_PORT")
	fmt.Println(httpServerPort)
	gitlabToken := viper.Get("GITLAB_TOKEN")
	fmt.Println(gitlabToken)
}
