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
}

type QorNotification struct {
	gorm.Model
	From        string
	To          string
	Title       string
	Body        string `sql:"size:65532"`
	MessageType string
	ResolvedAt  *time.Time
}

func (qorNotification QorNotification) IsResolved() bool {
	return qorNotification.ResolvedAt != nil
}

func (qorNotification QorNotification) Actions(context *admin.Context) (actions []*Action) {
	if n := context.Get("Notification"); n != nil {
		if notification, ok := n.(*Notification); ok {
			for _, action := range notification.Actions {
				if qorNotification.MessageType == action.MessageType {
					if action.Visible != nil {
						if !action.Visible(qorNotification, context) {
							continue
						}
					}

					actions = append(actions, action)
				}
			}
		}
	}

	if len(actions) == 0 {
		return []*Action{}
	}

	return
}
