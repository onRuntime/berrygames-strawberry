package data

import "github.com/onRuntime/berrygames-strawberry/model"

func (d *Data) GetServers() map[model.ServerType][]*model.Server {
	return d.servers
}

func (d *Data) GetAllServers() []*model.Server {
	var servers []*model.Server
	for _, s := range d.servers {
		servers = append(servers, s...)
	}
	return servers
}

func (d *Data) GetServersByType(serverType model.ServerType) []*model.Server {
	return d.servers[serverType]
}

func (d *Data) GetServer(id string) *model.Server {
	return nil
}

func (d *Data) AddServer(server *model.Server) {
	go func() {
		if err := d.db.Create(&server).Error; err != nil {
			panic(err)
		}
	}()

	d.Lock()
	d.servers[server.ServerType] = append(d.servers[server.ServerType], server)
	d.Unlock()
}

func (d *Data) RemoveServer(server *model.Server) {
	//TODO: Condition
	go d.db.Delete(&model.Server{}, "")

	d.Lock()
	d.Unlock()
}
