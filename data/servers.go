package data

import "github.com/onRuntime/berrygames-strawberry/model"

func (d *Data) GetAllServers() []*model.Server {
	var servers []*model.Server
	for _, s := range d.servers {
		servers = append(servers, s...)
	}
	return servers
}

func (d *Data) GetServers() map[model.ServerType][]*model.Server {
	return d.servers
}

func (d *Data) GetServersByType(serverType model.ServerType) []*model.Server {
	return d.servers[serverType]
}

func (d *Data) GetServer(id string) *model.Server {
	return nil
}

func (d *Data) RemoveServer(id string) {
}
