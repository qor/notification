package notification

import (
	"github.com/qor/admin"
	"github.com/qor/qor/utils"
)

func (notification *Notification) Action(action *Action) {
	if action.Label == "" {
		action.Label = utils.HumanizeString(action.Name)
	}

	if action.Method == "" {
		if action.URL != nil {
			action.Method = "GET"
		} else {
			action.Method = "PUT"
		}
	}

	if action.Resource != nil && action.Handle == nil {
		utils.ExitWithMsg("No Handler registered for action")
	}

	notification.Actions = append(notification.Actions, action)
}

type ActionArgument struct {
	Message  *QorNotification
	Context  *admin.Context
	Argument interface{}
}

type Action struct {
	Name        string
	Label       string
	Method      string
	MessageType string
	Resource    *admin.Resource
	Visible     func(data interface{}, context *admin.Context) bool
	URL         func(data interface{}, context *admin.Context) string
	Handle      func(*ActionArgument) error
}

// ToParam used to register routes for actions
func (action Action) ToParam() string {
	return utils.ToParamString(action.Name)
}
