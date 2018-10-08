package httphelpers

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

//get a job

const sessionName = "user-session"

func SetSessionVar(r *http.Request, w http.ResponseWriter, s string) {

	// Get a session. Get() always returns a session, even if empty.
	session, err := store.Get(r, sessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set some session values.
	session.Values[s] = "bar"
	session.Values[42] = 43
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)
}
func SetAuth(r *http.Request, w http.ResponseWriter, s string, isAuth bool) {

	// Get a session. Get() always returns a session, even if empty.
	session, err := store.Get(r, sessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set some session values.
	session.Values[s] = isAuth
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)
}

func CheckAuth(r *http.Request) bool {
	session, _ := store.Get(r, sessionName)

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return false
	}
	return true
}

func GetSessionVar(r *http.Request, w http.ResponseWriter) string {
	var words string

	return words
}

func CreateFlashSession(w http.ResponseWriter, r *http.Request) {
	// Get a session.
	session, err := store.Get(r, sessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the previous flashes, if any.
	if flashes := session.Flashes(); len(flashes) > 0 {
		// Use the flash values.
	} else {
		// Set a new flash.
		session.AddFlash("Hello, flash messages world!")
	}
	session.Save(r, w)
}
