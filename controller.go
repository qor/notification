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
	context.Execute("notifications", map[string]interface{}{
		"Messages": c.Notification.GetMessages(context.Context),
	})
}

func (c *controller) Action(context *admin.Context) {
	action := c.action

	if context.Request.Method == "GET" {
		context.Execute("action", action)
	} else {
		var actionArgument = ActionArgument{
			Context: context,
		}

		if action.Resource != nil {
			result := action.Resource.NewStruct()
			action.Resource.Decode(context.Context, result)
			actionArgument.Argument = result
		}

		if err := action.Handle(&actionArgument); err == nil {
			message := string(context.Admin.T(context.Context, "qor_admin.actions.executed_successfully", "Action {{.Name}}: Executed successfully", action))
			responder.With("html", func() {
				context.Flash(message, "success")
				http.Redirect(context.Writer, context.Request, context.Request.Referer(), http.StatusFound)
			}).With("json", func() {
				context.JSON("OK", map[string]string{"message": message, "status": "ok"})
			}).Respond(context.Request)
		} else {
			message := string(context.Admin.T(context.Context, "qor_admin.actions.executed_failed", "Action {{.Name}}: Failed to execute", action))
			context.JSON("OK", map[string]string{"error": message, "status": "error"})
		}
	}
}
