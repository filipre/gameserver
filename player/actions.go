package player

import (
	"errors"
	"time"

	"github.com/satori/go.uuid"
)

var (
	ErrNotImplemented = errors.New("Functionality not implemented yet.")
)

var (
	ErrDBExistName = errors.New("Could not check if name exists in database. There might be a problem right now.")
	ErrDBInsert    = errors.New("Could not insert player in database. There might be a problem right now.")
	ErrDBUpdate    = errors.New("Could not update player.")
	ErrDBFind      = errors.New("Could not find player. There might be a problem with the database right now.")
)

type Repository interface {
	ExistName(string) (bool, error)
	// Find([]uuid.UUID, []string, []Status) ([]*Player, error)
	FindByID(uuid.UUID) (*Player, error)
	// FindByName(string) (*Player, error)
	Insert(*Player) error
	UpdateStatus(uuid.UUID, Status) error
}

type Context struct {
	Repository
}

func NewContext(repo Repository) *Context {
	return &Context{repo}
}

var (
	ErrPlayerExists = errors.New("Could not create player because there is already a player with the given name.")
)

func (c *Context) Register(name string) (*Player, error) {
	exists, err := c.ExistName(name)
	if err != nil {
		return nil, ErrDBExistName
	}
	if exists {
		return nil, ErrPlayerExists
	}
	player, err2 := New(uuid.NewV4(), name, Joined, time.Now())
	if err2 != nil {
		return nil, err2
	}
	if err3 := c.Insert(player); err3 != nil {
		return nil, ErrDBInsert
	}
	return player, nil
}

var (
	ErrNotFound = errors.New("Could not find player in the System. Please check your ID.")
)

func (c *Context) Load(playerID uuid.UUID) (*Player, error) {
	player, err := c.FindByID(playerID)
	if err != nil {
		return nil, ErrDBFind
	}
	if player == nil {
		return nil, ErrNotFound
	}
	return player, nil
}

func (c *Context) Search(name []string, status []Status) ([]*Player, error) {
	return nil, ErrNotImplemented
}

func (c *Context) Activate(playerID uuid.UUID) error {
	return c.changeStatus(playerID, Activated)
}

func (c *Context) Deactivate(playerID uuid.UUID) error {
	return c.changeStatus(playerID, Deactivated)
}

func (c *Context) Ban(playerID uuid.UUID) error {
	return c.changeStatus(playerID, Banned)
}

func (c *Context) changeStatus(playerID uuid.UUID, newStatus Status) error {
	player, err := c.Load(playerID)
	if err != nil {
		return err
	}
	if player.Status == newStatus {
		return nil
	}
	if err := c.UpdateStatus(playerID, newStatus); err != nil {
		return ErrDBUpdate // TODO: log error to stderr: https://golang.org/pkg/log/#Logger
	}
	return nil
}
