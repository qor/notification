package notification

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
)

type Message struct {
	From        interface{}
	To          interface{}
	Title       string
	Body        string
	MessageType string
	Data        []byte
	ResolvedAt  *time.Time
}

type QorNotification struct {
	gorm.Model
	From        string
	To          string
	Title       string
	Body        string `sql:"size:65532"`
	MessageType string
	Data        []byte `sql:"size:65532"`
	ResolvedAt  *time.Time
}

func (qorNotification QorNotification) IsResolved() bool {
	return qorNotification.ResolvedAt != nil
}

func (qorNotification *QorNotification) Actions(context *admin.Context) (actions []*Action) {
	var globalActions []*Action
	if n := context.Get("Notification"); n != nil {
		if notification, ok := n.(*Notification); ok {
			for _, action := range notification.Actions {
				if action.HasMessageType(qorNotification.MessageType) {
					if action.Visible != nil {
						if !action.Visible(qorNotification, context) {
							continue
						}
					}

					if len(action.MessageTypes) == 0 {
						globalActions = append(globalActions, action)
					} else {
						actions = append(actions, action)
					}
				}
			}
		}
	}

	return append(actions, globalActions...)
}
