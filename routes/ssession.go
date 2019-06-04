package routes

import (
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"vc.cli/utils"
)

func GetSesssionStore() sessions.RedisStore {
	instance := utils.GetInstance()
	server := instance.Config.Server
	sessionStore := server.SessionStore
	store, err := sessions.NewRedisStore(sessionStore.Size, sessionStore.NetWork, sessionStore.Host+":"+strconv.Itoa(sessionStore.Port), sessionStore.Password, []byte("secret"))
	if err != nil {
		// Handle the error. Probably bail out if we can't connect.
	}
	store.Options(sessions.Options{
		Path:   "/",
		Domain: server.SESSION_COOKIE_DOMAIN,
		MaxAge: server.SESSION_TIMEOUT, //30min
		// Secure:   true,
		HttpOnly: true,
	})
	return store
}
