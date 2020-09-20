package data

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/onRuntime/berrygames-strawberry/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type Data struct {
	sync.Mutex

	db         *gorm.DB

	// - Stored data
	servers    map[model.ServerType][]*model.Server
	// Stored by Name [PROXY/HUB/SKYWARS/...]
	serverType map[string]*model.ServerType
	// Stored by player Name
	players    map[string]*model.Player
}

func New() *Data {
	return &Data{
		servers:    map[model.ServerType][]*model.Server{},
		serverType: map[string]*model.ServerType{},
		players:    map[string]*model.Player{},
	}
}

func (d *Data) Connect(addr, port, user, password, database, params string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", user, password, addr, port, database, params)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	d.db = db

	return nil
}

func (d *Data) Init() error {
	// - Migrates tables
	if err := d.db.AutoMigrate(&model.Server{}, &model.ServerType{}, &model.Player{}); err != nil {
		return err
	}

	// - Initialize cached data

	return nil
}

func (d *Data) Close() {
	//TODO: this
}
