package core

import (
	deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"
	feedback_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/feedback"
	montior_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/montior"
	user_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/user"
)

type CoreAdapter struct {
	usersServices    user_ports.IUsersServices
	feedbackServices feedback_ports.IFeedbackServices
	depsServices     deps_ports.IDepsServices
	monitorServices  montior_ports.IMonitorServices
}

func NewCoreAdapter(usersServices user_ports.IUsersServices, feedbackServices feedback_ports.IFeedbackServices, depsServices deps_ports.IDepsServices, monitorServices montior_ports.IMonitorServices) *CoreAdapter {
	return &CoreAdapter{
		usersServices:    usersServices,
		feedbackServices: feedbackServices,
		depsServices:     depsServices,
		monitorServices:  monitorServices,
	}
}

func (c *CoreAdapter) UsersServices() user_ports.IUsersServices {
	return c.usersServices
}

func (c *CoreAdapter) FeedbackServices() feedback_ports.IFeedbackServices {
	return c.feedbackServices
}

func (c *CoreAdapter) DepsServices() deps_ports.IDepsServices {
	return c.depsServices
}

func (c *CoreAdapter) MonitorServices() montior_ports.IMonitorServices {
	return c.monitorServices
}
