package user_services

import (
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"
)

type usersServices struct {
	userService user_ports.IUsersService
}

func NewUsersServices(usersRepositories user_ports.IUsersRepositories, deps deps_ports.IDepsServices) user_ports.IUsersServices {
	return &usersServices{
		userService: newUsersService(usersRepositories, deps),
	}
}

func (s *usersServices) UsersService() user_ports.IUsersService {
	return s.userService
}
