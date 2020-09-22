package data

import (
	"errors"
	"github.com/onRuntime/berrygames-strawberry/model"
)

func (d *Data) GetPlayers() map[string]*model.Player {
	return d.players
}

func (d *Data) GetPlayer(name string) (*model.Player, error) {
	player, ok := d.players[name]
	if !ok {
		return nil, errors.New("")
	}
	return player, nil
}
