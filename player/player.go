package player

import (
	"errors"
	"regexp"
	"time"

	"github.com/satori/go.uuid"
)

type Player struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Status  Status    `json:"status"`
	Changed time.Time `json:"changed"`
}

type Status string

const (
	Joined      Status = "joined"
	Activated   Status = "activated"
	Deactivated Status = "deactivated"
	Banned      Status = "banned"
)

func New(id uuid.UUID, name string, status Status, changed time.Time) (*Player, error) {
	p := &Player{id, name, status, changed}
	if err := p.Validate(); err != nil {
		return nil, err
	}
	return p, nil
}

type checker interface {
	check() error
}

type allowedName struct {
	*Player
}

var (
	ErrBadName = errors.New("Provided name is not valid")
)

func (v *allowedName) check() error {
	if !regexp.MustCompile(`^[\w ]{3,30}$`).MatchString(v.Name) {
		return ErrBadName
	}
	return nil
}

func (p *Player) Validate() error {
	testFor := []checker{&allowedName{p}}
	for _, test := range testFor {
		if test.check() != nil {
			return test.check()
		}
	}
	return nil
}
