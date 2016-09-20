package notification

import (
	"errors"
	"fmt"
	"strings"

	"github.com/qor/admin"
	"github.com/qor/qor/utils"
)

func (notification *Notification) Action(action *Action) error {
	if a := notification.GetAction(action.Name); a != nil && a.MessageType == action.MessageType {
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
		if action.Name == name {
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
	if action.MessageType == "" {
		return utils.ToParamString(action.Name)
	}
	return strings.Join([]string{utils.ToParamString(action.MessageType), utils.ToParamString(action.Name)}, "/")
}
