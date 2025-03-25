package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

var data *viper.Viper

func init() {

	data = viper.New()

	data.SetConfigType("yaml")

	if len(os.Args) > 1 {

		data.SetConfigFile(os.Args[1])
	} else {

		data.SetConfigName("config")
	}

	data.AutomaticEnv()

	data.SetEnvPrefix("gobit")

	data.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := data.ReadInConfig()
	if err != nil {

		panic("error in reading config file: " + err.Error())
	}

	data.AllSettings()
}

func ApiHost() string {

	return data.GetString("api.host")
}

func ApiPort() string {

	return data.GetString("api.port")
}

func ApiDebug() bool {

	return data.GetBool("api.debug")
}

func ApiAuthUsername() string {

	return data.GetString("api.auth.username")
}

func ApiAuthPassword() string {

	return data.GetString("api.auth.password")
}

func ApiWhitelistIp() []string {

	return data.GetStringSlice("api.whitelist_ip")
}

func SwaggerTitle() string {

	return data.GetString("swagger.title")
}

func SwaggerHost() string {

	return data.GetString("swagger.host")
}

func SwaggerVersion() string {

	return data.GetString("swagger.version")
}

func SwaggerProtocol() string {

	return data.GetString("swagger.protocol")
}

func PostgresDsn() string {

	return data.GetString("postgres.dsn")
}
