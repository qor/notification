package notification

import (
	"net/http"

	"github.com/qor/admin"
	"github.com/qor/responder"
)

type controller struct {
	Notification *Notification
	action       *Action
}

func (c *controller) List(context *admin.Context) {
	context.Execute("notifications/notifications", map[string]interface{}{
		"Messages": c.Notification.GetNotifications(context.CurrentUser, context.Context),
	})
}

func (c *controller) Action(context *admin.Context) {
	action := c.action
	message := c.Notification.GetNotification(context.CurrentUser, context.ResourceID, context.Context)

	if context.Request.Method == "GET" {
		context.Execute("action", action)
	} else {
		var actionArgument = &ActionArgument{
			Message: message,
			Context: context,
		}

		if action.Resource != nil {
			result := action.Resource.NewStruct()
			action.Resource.Decode(context.Context, result)
			actionArgument.Argument = result
		}

		if err := action.Handle(actionArgument); err == nil {
			notice := string(context.Admin.T(context.Context, "qor_admin.actions.executed_successfully", "Action {{.Name}}: Executed successfully", action))
			responder.With("html", func() {
				context.Flash(notice, "success")
				http.Redirect(context.Writer, context.Request, context.Request.Referer(), http.StatusFound)
			}).With("json", func() {
				context.JSON("OK", map[string]string{"message": notice, "status": "ok"})
			}).Respond(context.Request)
		} else {
			notice := string(context.Admin.T(context.Context, "qor_admin.actions.executed_failed", "Action {{.Name}}: Failed to execute", action))
			context.JSON("OK", map[string]string{"error": notice, "status": "error"})
		}
	}
}
