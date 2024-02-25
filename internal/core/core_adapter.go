package core

import (
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	roadmap_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/roadmap"
	social_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/social"
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"
)

type CoreAdapter struct {
	usersServices   user_ports.IUsersServices
	depsServices    deps_ports.IDepsServices
	socialServices  social_ports.ISocialServices
	roadmapServices roadmap_ports.IRoadmapServices
}

func NewCoreAdapter(
	usersServices user_ports.IUsersServices,
	depsServices deps_ports.IDepsServices,
	socialService social_ports.ISocialServices,
	roadmapServices roadmap_ports.IRoadmapServices,
) *CoreAdapter {
	return &CoreAdapter{
		usersServices:   usersServices,
		depsServices:    depsServices,
		socialServices:  socialService,
		roadmapServices: roadmapServices,
	}
}

func (c *CoreAdapter) UsersServices() user_ports.IUsersServices {
	return c.usersServices
}

func (c *CoreAdapter) DepsServices() deps_ports.IDepsServices {
	return c.depsServices
}
func (c *CoreAdapter) SocialServices() social_ports.ISocialServices {
	return c.socialServices
}

func (c *CoreAdapter) RoadmapServices() roadmap_ports.IRoadmapServices {
	return c.roadmapServices
}
