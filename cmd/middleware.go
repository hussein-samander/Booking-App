package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit D Page.")
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRF tokens to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		//Doesn't using app.InProduction here means it has default values rather than the ones defined within
		//main? We can use the variable because it is defined outside of the func in main.go
		//but the values are all set within the func body and app is not a pointer to config.AppConfig
		//so those values should not be available here and the value will always be false as it is the default.
		//Can we use the global Repo var?
		//Or why not create separate functions in main that set the InProduction + create a new session
		//and have the value returned here instead?
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves the session upon every request
func SessionLoad(next http.Handler) http.Handler {
	//Same issue as with app.InProduction
	//This is a session with default values, not the ones assigned in main
	//Ver2: session is a POINTER *scs.SessionManager
	return session.LoadAndSave(next)
}
