package server

import (
	"github.com/restmark/goauth/hooks"
	"github.com/sirupsen/logrus"
)

func (a *API) SetupLogger() {
	level, err := logrus.ParseLevel(a.Config.GetString("log_level"))
	if err == nil {
		logrus.SetLevel(level)
	}

	if slackHook := a.Config.GetString("error_slack_hook"); slackHook != "" {
		logrus.AddHook(&hooks.SlackrusHook{
			HookURL: slackHook,
			AcceptedLevels: []logrus.Level{
				logrus.ErrorLevel,
			},
			Channel: "#errors",
		})
	}
}
