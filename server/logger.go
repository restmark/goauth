package server

import (
	log "github.com/sirupsen/logrus"
)

func (a *API) SetupLogger() {
	level, err := log.ParseLevel(a.Config.GetString("log_level"))
	if err == nil {
		log.SetLevel(level)
	}
}
