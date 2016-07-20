package notification

import "github.com/qor/admin"

type controller struct {
	Notification *Notification
	action       *Action
}

func (c *controller) List(context *admin.Context) {
	context.Execute("notifications", nil)
}

func (c *controller) Action(context *admin.Context) {
	if context.Request.Method == "GET" {
		context.Execute("action", c.action)
	} else {
	}
}
