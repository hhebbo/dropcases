package api

import (
	"github.com/heroku/go-getting-started/src/apps/appConfig"
	"github.com/heroku/go-getting-started/src/packages/core/config"
)

func getToken(authType string) string {
	var token string

	switch authType {
	case AuthTypeFiles:
		token = config.GetValue(appConfig.DROPCASES_API_TOKEN_FILE)
	case AuthTypeMembers:
		token = config.GetValue(appConfig.DROPCASES_API_TOKEN_MEMBER)
	}

	return token
}
