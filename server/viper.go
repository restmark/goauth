package server

import (
	"github.com/joho/godotenv"
	"os"
)

func (a *API) SetupViper() error {

	filename := ".env"
	switch os.Getenv("GOAUTH_ENV") {
	case "testing":
		filename = "../.env.testing"
	case "prod":
		filename = ".env.prod"
	}

	godotenv.Overload(filename)

	a.Config.SetEnvPrefix("goauth")
	a.Config.AutomaticEnv()
	a.Config.AddConfigPath(".")
	a.Config.SetConfigName("config")

	err := a.Config.ReadInConfig()
	if err != nil {
		return err
	}

	a.SetupViperDefaults()

	return nil
}

func (a *API) SetupViperDefaults() {

}
