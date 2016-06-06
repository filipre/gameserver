package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// type repository interface {
// 	player.Repository
// 	// LeagueCRUD
// 	// MemberCRUD
// }

func newConnection(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// func (db *gameserverDB) Close() error {
// 	return db.Close()
// }
