package data

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/onRuntime/berrygames-strawberry/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

type Data struct {
	sync.Mutex

	db          *gorm.DB
	redis       *redis.Client

	// - Stored data
	servers     map[model.ServerType][]*model.Server
	// Stored by Name [PROXY/HUB/SKYWARS/...]
	serverTypes map[string]*model.ServerType
	// Stored by player Name
	players     map[string]*model.Player

	queues      map[string]*Queue
}

func New() *Data {
	return &Data{
		servers:     make(map[model.ServerType][]*model.Server),
		serverTypes: make(map[string]*model.ServerType),
		players:     make(map[string]*model.Player),
		queues:      make(map[string]*Queue),
	}
}

func (d *Data) DBConnect(addr, port, user, password, database, params string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", user, password, addr, port, database, params)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return err
	}
	d.db = db

	return nil
}

func (d *Data) RedisConnect(addr, port, password string, database int) error {
	d.redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", addr, port),
		Password: password,
		DB:       database,
	})
	return d.redis.Ping(context.Background()).Err()
}

func (d *Data) Init() error {
	// - Migrates tables
	if err := d.db.AutoMigrate(&model.ServerType{}, &model.Server{}, &model.Rank{}, &model.Player{}); err != nil {
		return err
	}

	// - Initialize cached data
	var players []model.Player
	d.db.Find(&players)
	for _, player := range players {
		d.players[player.PlayerName] = &player
	}
	//TODO: GORM query

	// - Initialize queues
	for _, serverType := range d.serverTypes {
		queue := &Queue{serverType.Name, []string{}}
		d.queues[serverType.Name] = queue
		go queue.run()
	}

	return nil
}

func (d *Data) Close() {
	//TODO: this
}
