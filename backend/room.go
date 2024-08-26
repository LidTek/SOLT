package main

import (
	"fmt"
	"time"

	"math/rand"
)

// Room represents a room that exists.
type Room struct {
	Users  []string // Users in the room and their reported status (ready or not)
	Status bool     // status is the room status, either started or not
	Admin  string   // Admin is the user who created the room
	Code   string   // Code is the room code
	last   time.Time
}

func (r *Room) addUser(id string) {
	for _, user := range r.Users {
		if user == id {
			return
		}
	}
	r.Users = append(r.Users, id)
}

func (r *Room) removeUser(id string) {
	for i, user := range r.Users {
		if user == id {
			r.Users = append(r.Users[:i], r.Users[i+1:]...)
			break
		}
	}
}

func (r *Room) expired() bool {
	return time.Since(r.last) > 120*time.Second
}

var rooms = make(map[string]Room)

func cleanupRooms() int {
	count := 0
	for id, room := range rooms {
		if room.expired() {
			delete(rooms, id)
			count++
		}
	}
	return count
}

func touchRoom(id string) {
	if room, ok := rooms[id]; ok {
		room.last = time.Now()
		rooms[id] = room
	}
}

// generateRoomCode generates a 5 character string consisting of 0-9 and a-z
func generateRoomCode() string {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyz"
	code := make([]byte, 5)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

// acquireRoomCode attempts to generate a new room code and compares against existing rooms.
func acquireRoomCode() (string, error) {
	// This is kind of hacky, but whatever.
	for i := 0; i < 100; i++ {
		code := generateRoomCode()
		if _, ok := rooms[code]; !ok {
			return code, nil
		}
	}
	return "", ErrNoRoomCode
}

// ErrNoRoomCode indicates that there was a failure in acquiring a room code.
var ErrNoRoomCode = fmt.Errorf("no room code available")
