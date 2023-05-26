package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Note struct {
	ID   uint32 `json:"id"`
	Note string `json:"note"`
}

type Session struct {
	SID string `json:"sid"`
}

var users []User
var notes []Note

func main() {
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/notes", notesHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users = append(users, user)
	w.WriteHeader(http.StatusOK)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, u := range users {
		if u.Email == user.Email && u.Password == user.Password {
			session := Session{SID: generateSessionID()}
			response := struct {
				SID string `json:"sid"`
			}{SID: session.SID}

			jsonResponse(w, http.StatusOK, response)
			return
		}
	}

	w.WriteHeader(http.StatusUnauthorized)
}

func notesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		listNotesHandler(w, r)
	} else if r.Method == http.MethodPost {
		createNoteHandler(w, r)
	} else if r.Method == http.MethodDelete {
		deleteNoteHandler(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func listNotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	sessionID := r.URL.Query().Get("sid")
	if !isValidSession(sessionID) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jsonResponse(w, http.StatusOK, struct {
		Notes []Note `json:"notes"`
	}{Notes: notes})
}

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionID := r.URL.Query().Get("sid")
	if !isValidSession(sessionID) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	note.ID = generateNoteID()
	notes = append(notes, note)

	response := struct {
		ID uint32 `json:"id"`
	}{ID: note.ID}

	jsonResponse(w, http.StatusOK, response)
}

func deleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

