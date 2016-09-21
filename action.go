package notification

import (
	"errors"
	"fmt"

	"github.com/qor/admin"
	"github.com/qor/qor/utils"
)

func (notification *Notification) Action(action *Action) error {
	if a := notification.GetAction(action.Name); a != nil {
		message := fmt.Sprintf("Action %v already registered", action.Name)
		fmt.Println(message)
		return errors.New(message)
	}

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
	return nil
}

func (notification *Notification) GetAction(name string) *Action {
	for _, action := range notification.Actions {
		if utils.ToParamString(action.Name) == utils.ToParamString(name) {
			return action
		}
	}
	return nil
}

type ActionArgument struct {
	Message  *QorNotification
	Context  *admin.Context
	Argument interface{}
}

type Action struct {
	Name         string
	Label        string
	Method       string
	MessageTypes []string
	Resource     *admin.Resource
	Visible      func(data *QorNotification, context *admin.Context) bool
	URL          func(data *QorNotification, context *admin.Context) string
	Handle       func(*ActionArgument) error
	Undo         func(*ActionArgument) error
}

// ToParam used to register routes for actions
func (action Action) ToParam() string {
	return utils.ToParamString(action.Name)
}

func (action Action) HasMessageType(t string) bool {
	for _, mt := range action.MessageTypes {
		if mt == t {
			return true
		}
	}
	if len(action.MessageTypes) == 0 {
		return true
	}
	return false
}
