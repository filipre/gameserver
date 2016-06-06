package main

import (
	"database/sql"
	"fmt"
	"time"

	"./player"
	"github.com/satori/go.uuid"
)

type playerDB struct {
	*sql.DB
}

func newPlayerDB(db *sql.DB) *playerDB {
	return &playerDB{db}
}

func (db *playerDB) ExistName(name string) (bool, error) {
	// TODO
	// stmt, err := `SELECT EXISTS(SELECT 1 FROM player WHERE name = $1)`
	// TODO
	return false, fmt.Errorf("Not implemented")
}

// func (db *playerDB) Find(interface{}) ([]*player.Player, error) {
// 	//TODO define parameters
// 	return nil, fmt.Errorf("Not implemented")
// }

func (db *playerDB) FindByID(id uuid.UUID) (*player.Player, error) {
	sql := `SELECT player.id, player.name, player_status.status, player_status.changed
    FROM player, player_status
    WHERE player.id = $1
    AND player_status.player = $1
    ORDER BY player_status.changed DESC
    LIMIT 1;`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	var playerID uuid.UUID
	var playerName string
	var playerStatus []byte
	var playerChanged time.Time
	if err = stmt.QueryRow(id).Scan(&playerID, &playerName, &playerStatus, &playerChanged); err != nil {
		return nil, err
	}
	player, err := player.New(playerID, playerName, player.Status(playerStatus), playerChanged)
	if err != nil {
		return nil, err
	}
	return player, nil
}

// func (db *playerDB) Count() (int, error) {
// 	row := db.QueryRow("SELECT COUNT(id) FROM player;")
// 	var n int
// 	if err := row.Scan(&n); err != nil {
// 		return -1, err
// 	}
// 	return n, nil
// }

func (db *playerDB) Insert(player *player.Player) error {
	insertPlayer, err := db.Prepare(`INSERT INTO player (id, name) VALUES ($1, $2);`)
	if err != nil {
		return err
	}
	insertPlayerStatus, err := db.Prepare(`INSERT INTO player_status (id, player, status, changed) VALUES ($1, $2, $3, $4);`)
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Stmt(insertPlayer).Exec(player.ID, player.Name); err != nil {
		return err
	}
	if _, err := tx.Stmt(insertPlayerStatus).Exec(uuid.NewV4(), player.ID, []byte(player.Status), player.Changed); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (db *playerDB) UpdateStatus(id uuid.UUID, new player.Status) error {
	stmt, err := db.Prepare(`INSERT INTO player_status (id, player, status, changed) VALUES ($1, $2, $3, $4);`)
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(uuid.NewV4(), id, []byte(new), time.Now()); err != nil {
		return err
	}
	return nil
}
