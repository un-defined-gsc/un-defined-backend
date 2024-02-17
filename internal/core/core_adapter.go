package core

import (
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"
)

type CoreAdapter struct {
	usersServices user_ports.IUsersServices
	depsServices  deps_ports.IDepsServices
}

func NewCoreAdapter(usersServices user_ports.IUsersServices, depsServices deps_ports.IDepsServices) *CoreAdapter {
	return &CoreAdapter{
		usersServices: usersServices,
		depsServices:  depsServices,
	}
}

func (c *CoreAdapter) UsersServices() user_ports.IUsersServices {
	return c.usersServices
}

func (c *CoreAdapter) DepsServices() deps_ports.IDepsServices {
	return c.depsServices
}
