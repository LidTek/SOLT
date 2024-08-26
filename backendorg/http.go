package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// USERS
type userCreateRequest struct {
	Dial string `json:"dial"`
}
type userCreateResponse struct{}

func httpCreateUser(w http.ResponseWriter, r *http.Request, user string) {
	// Read body as JSON.
	var req userCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Dial == "" {
		http.Error(w, "missing dial", http.StatusBadRequest)
		return
	}

	touchUser(user)
	setUserDial(user, req.Dial)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userCreateResponse{})
}

type userGetResponse struct {
	Dial string `json:"dial"`
}

func httpGetUser(w http.ResponseWriter, r *http.Request, user string) {
	u, ok := users[user]
	if !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userGetResponse{Dial: u.DialURL})
}

// ROOMS

type roomCreateResponse struct {
	Code string `json:"code"`
}

func httpCreateRoom(w http.ResponseWriter, r *http.Request, user string) {
	code, err := acquireRoomCode()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	room := Room{
		Admin: user,
		Code:  code,
		last:  time.Now(),
	}
	room.addUser(user)
	rooms[room.Code] = room

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(roomCreateResponse{Code: room.Code})
}

type roomDestroyResponse struct {
}

func httpDestroyRoom(w http.ResponseWriter, r *http.Request, user string, code string) {
	room, ok := rooms[code]
	if !ok {
		http.Error(w, "room not found", http.StatusNotFound)
		return
	}

	if room.Admin != user {
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return
	}

	delete(rooms, code)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(roomDestroyResponse{})
}

type roomJoinResponse struct {
	Ready bool `json:"status"`
}

func httpJoinRoom(w http.ResponseWriter, r *http.Request, user string, code string) {
	room, ok := rooms[code]
	if !ok {
		http.Error(w, "room not found", http.StatusNotFound)
		return
	}

	room.addUser(user)
	touchRoom(code)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(roomJoinResponse{Ready: len(room.Users) > 1})
}
