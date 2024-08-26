package main

import "time"

// User represents a client connected to the backend. Much of this information is sent to other clients.
type User struct {
	Name    string    // Name is advertised
	DialURL string    // DialURL is the URL the advertised URL for other clients to dial in WebRTC.
	last    time.Time // last is the last time the user was seen.
}

var users = make(map[string]User)

func cleanupUsers() int {
	count := 0
	for id, user := range users {
		if time.Since(user.last) > 30*time.Second {
			delete(users, id)
			count++
		}
	}
	return count
}

func touchUser(id string) {
	if _, ok := users[id]; !ok {
		users[id] = User{
			last: time.Now(),
		}
	} else {
		user := users[id]
		user.last = time.Now()
		users[id] = user
	}
}

func setUserDial(id, dialURL string) {
	if user, ok := users[id]; ok {
		user.DialURL = dialURL
		user.last = time.Now()
		users[id] = user
	}
}
