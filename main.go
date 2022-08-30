package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

// Game - Our struct for all Games
type Game struct {
	Pool     string `json:"pool"`
	Team_A   string `json:"team_a"`
	Team_B   string `json:"team_b"`
	Time     string `json:"time"`
	Division string `json:"division"`
}

var Games []Game

func returnAllGames(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Games)
}

func returnGames(w http.ResponseWriter, r *http.Request) {
	key := path.Base(r.URL.Path)
	var s []Game

	for _, Game := range Games {
		if Game.Pool == key {
			s = append(s, Game)
		}
	}

	if len(s) > 0 {
		json.NewEncoder(w).Encode(s)
	} else {
		json.NewEncoder(w).Encode("No pools labeled " + key)
	}

}

func createNewGame(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Game struct
	// append this to our Games array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var Game Game
	json.Unmarshal(reqBody, &Game)
	// update our global Games array to include
	// our new Game
	Games = append(Games, Game)

	json.NewEncoder(w).Encode(Game)
}

func deleteGame(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)

	for index, Game := range Games {
		if Game.Pool == id {
			Games = append(Games[:index], Games[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/games", returnAllGames)
	// myRouter.HandleFunc("/addGame", createNewGame).Methods("POST")
	// myRouter.HandleFunc("/removeGame/{pool}", deleteGame).Methods("DELETE")
	myRouter.HandleFunc("/getGame/{pool}", returnGames)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Games = []Game{
		{Pool: "A", Team_A: "Swps", Team_B: "Storm", Time: "04 Oct 2022 00:12:00 PST", Division: "Open"},
		{Pool: "A", Team_A: "Swps", Team_B: "UVIC", Time: "04 Oct 2022 00:12:30 PST", Division: "Open"},
		{Pool: "A", Team_A: "UVIC", Team_B: "Storm", Time: "04 Oct 2022 00:13:00 PST", Division: "Open"},
		{Pool: "B", Team_A: "North West", Team_B: "UVIC", Time: "04 Oct 2022 00:12:00 PST", Division: "Bantom"},
		{Pool: "B", Team_A: "Fraser River", Team_B: "North West", Time: "04 Oct 2022 00:12:30 PST", Division: "Bantom"},
		{Pool: "B", Team_A: "Swps", Team_B: "North West", Time: "04 Oct 2022 00:13:00 PST", Division: "Bantom"},
		{Pool: "C", Team_A: "Fraser River", Team_B: "OpenA", Time: "04 Oct 2022 00:12:00 PST", Division: "Atom"},
		{Pool: "C", Team_A: "AtomA", Team_B: "OpenA", Time: "04 Oct 2022 00:12:30 PST", Division: "Atom"},
		{Pool: "C", Team_A: "AtomA", Team_B: "Storm", Time: "04 Oct 2022 00:13:00 PST", Division: "Atom"},
		{Pool: "D", Team_A: "Swps", Team_B: "Storm", Time: "04 Oct 2022 00:12:00 PST", Division: "Peewee"},
		{Pool: "D", Team_A: "Swps", Team_B: "UVIC", Time: "04 Oct 2022 00:12:30 PST", Division: "Peewee"},
		{Pool: "D", Team_A: "UVIC", Team_B: "Storm", Time: "04 Oct 2022 00:13:00 PST", Division: "Peewee"},
	}

	handleRequests()
}
