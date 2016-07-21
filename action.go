package notification

import (
	"github.com/qor/admin"
	"github.com/qor/qor/utils"
)

func (notification *Notification) Action(action *Action) {
	notification.Actions = append(notification.Actions, action)
}

type ActionArgument struct {
	Message  *QorNotification
	Context  *admin.Context
	Argument interface{}
}

type Action struct {
	Name     string
	Method   string
	URL      string
	Resource *admin.Resource
	Handle   func(*ActionArgument) error
}

// ToParam used to register routes for actions
func (action Action) ToParam() string {
	return utils.ToParamString(action.Name)
}
