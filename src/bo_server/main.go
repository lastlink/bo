package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/mux"

	"boardInfo"
	"gameState"
)

var activeGame *gameState.Game

func init() {
	activeGame = gameState.NewGame([]string{
		"Player 1",
		"Player 2",
		"Player 3",
		"Player 4",
		"Player 5",
		"Player 6",
	})
}

func getboardInfo(w http.ResponseWriter, r *http.Request) {
	if buf, err := boardInfo.JsonMap(); err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf)
	}
}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	if buf, err := json.Marshal(activeGame.Players); err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf)
	}
}

func getCompanies(w http.ResponseWriter, r *http.Request) {
	if buf, err := json.Marshal(activeGame.Companies); err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf)
	}
}

func main() {
	staticFS := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo}
	r := mux.NewRouter()
	r.HandleFunc("/board_info", getboardInfo)
	r.HandleFunc("/companies", getCompanies)
	r.HandleFunc("/players", getPlayers)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(staticFS)))

	fmt.Println("Hello World!")
	http.ListenAndServe(":8000", r)
}
