package notification

type controller struct {
	Notification *Notification
}

func (c *controller) List(context *Context) {
	context.Execute("notifications", nil)
}

func (c *controller) Action(context *Context) {
	if context.Request.Method == "GET" {
		context.Execute("action", action)
	} else {
	}
}
