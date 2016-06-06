package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gocraft/web"
	"github.com/satori/go.uuid"
)

type registerPlayerJSON struct {
	Name string `json:"name"`
}

func (game *gameserver) registerPlayer(rw web.ResponseWriter, req *web.Request) {
	reqJSON := new(registerPlayerJSON)
	if err := getJSON(req, reqJSON); err != nil {
		responseError(rw, err)
		return
	}
	player, err := game.player.Register(reqJSON.Name)
	if err != nil {
		responseError(rw, err)
		return
	}
	response(rw, "Player has been successfully registered", http.StatusCreated, player)
}

func (game *gameserver) loadPlayer(rw web.ResponseWriter, req *web.Request) {
	id, err := getID(req)
	if err != nil {
		responseError(rw, err)
		return
	}
	player, err := game.player.Load(id)
	if err != nil {
		responseError(rw, err)
		return
	}
	response(rw, "Player found", http.StatusOK, player)
}

func (game *gameserver) searchPlayers(rw web.ResponseWriter, req *web.Request) {
	//TODO
}

func (game *gameserver) activatePlayer(rw web.ResponseWriter, req *web.Request) {
	id, err := getID(req)
	if err != nil {
		responseError(rw, err)
		return
	}
	if err := game.player.Activate(id); err != nil {
		responseError(rw, err)
		return
	}
	response(rw, "Player activated", http.StatusOK, nil)
}

func (game *gameserver) deactivatePlayer(rw web.ResponseWriter, req *web.Request) {
	id, err := getID(req)
	if err != nil {
		responseError(rw, err)
		return
	}
	if err := game.player.Deactivate(id); err != nil {
		responseError(rw, err)
		return
	}
	response(rw, "Player deactivated", http.StatusOK, nil)
}

func (game *gameserver) banPlayer(rw web.ResponseWriter, req *web.Request) {
	id, err := getID(req)
	if err != nil {
		responseError(rw, err)
		return
	}
	if err := game.player.Ban(id); err != nil {
		responseError(rw, err)
		return
	}
	response(rw, "Player banned", http.StatusOK, nil)
}

var (
	errGetJSON = errors.New("Can't decode JSON. It's propably not valid.")
)

func getJSON(req *web.Request, target interface{}) error {
	if err := json.NewDecoder(req.Body).Decode(target); err != nil {
		return errGetJSON
	}
	return nil
}

var (
	errGetID = errors.New("Can't parse id to UUID. ID is propably not a valid UUID.")
)

func getID(req *web.Request) (uuid.UUID, error) {
	id, err := uuid.FromString(req.PathParams["id"])
	if err != nil {
		return uuid.Nil, errGetID
	}
	return id, nil
}
