package helper

import (
	"fmt"
	"tng/common/utils/cfgutil"
)

const (
	PreUserProfile = "user_profile"
	PreUserSession = "user_session"
)

var (
	appName = cfgutil.Load("AppName")
)

func RedisKeyUserProfile(userID, appID string) string {
	k := fmt.Sprintf("%s-%s-%s-%s", appName, appID, PreUserProfile, userID)
	return k
}

func RedisKeyUserSession(appID, platform, userID string) string {
	k := fmt.Sprintf("%s-%s-%s-%s-%s", appName, appID, PreUserSession, platform, userID)
	return k
}
