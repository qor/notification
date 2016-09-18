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
	From         string
	To           string
	Title        string
	Body         string `sql:"size:65532"`
	MessageType  string
	AckedAt      *time.Time
	Notification *Notification `sql:"-"`
}

func (qorNotification *QorNotification) Actions(context *admin.Context) (actions []*Action) {
	for _, action := range qorNotification.Notification.Actions {
		if qorNotification.MessageType == action.MessageType {
			if action.Visible != nil {
				if !action.Visible(qorNotification, context) {
					continue
				}
			}

			actions = append(actions, action)
		}
	}

	return
}
