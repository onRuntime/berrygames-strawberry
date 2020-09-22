package data

import "context"

type Queue struct {
	ServerType string
	Players    []string
}

func (q *Queue) run() {}

func (q *Queue) addPlayer(playerName string) {
	q.Players = append(q.Players, playerName)
}

func (q *Queue) removePlayer(playerName string) {
}


func (d *Data) getQueue(serverType string) *Queue {
	queue, ok := d.queues[serverType]
	if !ok {
		return nil
	}
	return queue
}

func (d *Data) getPlayerQueue(playerName string) *Queue {
	return nil
}

func (d *Data) IsPlayerInQueue(playerName string) (bool, *Queue) {
	return false, nil
}

func (d *Data) AddPlayerToQueue(serverType, playerName string) {
	queue := d.getQueue(serverType)
	queue.addPlayer(playerName)
	d.redis.Publish(context.Background(), "", "")
}

func (d *Data) RemovePlayerFromQueue(playerName string) {
	queue := d.getPlayerQueue(playerName)
	queue.removePlayer(playerName)
	d.redis.Publish(context.Background(), "", "")
}
