package main

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("don't care"))

// track adds a session for the given user so we can identify who they are.
func track(w http.ResponseWriter, r *http.Request) (*sessions.Session, error) {
	session, err := store.Get(r, "session")
	if err != nil {
		return nil, err
	}

	if session.IsNew {
		session.Values["ID"] = uuid.NewString()
		if err := session.Save(r, w); err != nil {
			return nil, err
		}
	}
	return session, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		session, err := track(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Hello, " + session.Values["ID"].(string)))
	})

	http.HandleFunc("/rooms", func(w http.ResponseWriter, r *http.Request) {
		session, err := track(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		switch r.Method {
		case "POST": // POST attempts to create the room.
			httpCreateRoom(w, r, session.Values["ID"].(string))
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/rooms/{rid}", func(w http.ResponseWriter, r *http.Request) {
		session, err := track(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		switch r.Method {
		case "POST": // POST attempts to create the room.
			httpCreateRoom(w, r, session.Values["ID"].(string))
		case "DELETE": // DELETE attempts to delete the room if the user is the admin.
			httpDestroyRoom(w, r, session.Values["ID"].(string), r.PathValue("rid"))
		case "GET": // GET attempts to join the room.
			httpJoinRoom(w, r, session.Values["ID"].(string), r.PathValue("rid"))
		case "PUT": // PUT attempts to change the state of the game from the admin.
			w.Write([]byte("PUT /rooms/{rid}"))
		}
	})

	http.HandleFunc("/users/{uid}", func(w http.ResponseWriter, r *http.Request) {
		session, err := track(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		switch r.Method {
		case "POST": // POST attempts to set the session user's dial.
			httpCreateUser(w, r, session.Values["ID"].(string))
		case "GET": // GET attempts to request the given user's dial.
			httpGetUser(w, r, r.PathValue("uid"))
		}
	})

	go func() {
		var totalUsers, totalRooms int
		for {
			totalUsers += cleanupUsers()
			time.Sleep(10 * time.Second)
			totalRooms += cleanupRooms()
			time.Sleep(10 * time.Second)
			if totalUsers > 30 {
				log.Printf("cleaned up %d users", totalUsers)
				totalUsers = 0
			}
			if totalRooms > 30 {
				log.Printf("cleaned up %d rooms", totalRooms)
				totalRooms = 0
			}
		}
	}()

	http.ListenAndServe(":9000", nil)
}
