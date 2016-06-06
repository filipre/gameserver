package main

import (
	"net/http"
	"os"

	"./player"

	"github.com/gocraft/web"
)

type gameserver struct {
	// db DatabaseCRUD
	player *player.Context
	// league league.Context
	// member member.Context
}

func main() {
	game := &gameserver{}

	db, err := newConnection(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	game.player = player.NewContext(newPlayerDB(db))
	// game.league = league.NewContext(LeagueDB(db))
	// game.member = member.NewContext(MemberDB(db))

	router := web.New(*game)
	router.Get("/player", game.searchPlayers)
	router.Get("/player/:id", game.loadPlayer)
	router.Post("/player", game.registerPlayer)
	router.Post("/player/:id/activate", game.activatePlayer)
	router.Post("/player/:id/deactivate", game.deactivatePlayer)
	router.Post("/player/:id/ban", game.banPlayer)

	http.ListenAndServe("localhost:8888", router)
}
