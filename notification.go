package notification

import (
	"path"

	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/roles"
)

type Notification struct {
	Config   *Config
	Channels []ChannelInterface
	Actions  []*Action
}

func New(config *Config) *Notification {
	return &Notification{Config: config}
}

func (notification *Notification) Send(message *Message, context *qor.Context) {
	for _, channel := range notification.Channels {
		channel.Send(message, context)
	}
}

func (notification *Notification) ConfigureQorResource(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		Admin := res.GetAdmin()
		router := Admin.GetRouter()
		notificationController := controller{Notification: notification}

		router.Get("!notifications", notificationController.List, admin.RouteConfig{
			PermissionMode: roles.Read,
			Resource:       res,
		})

		for _, action := range notification.Actions {
			actionController := controller{Notification: notification, action: action}
			router.Get(path.Join("!notifications/:message_id", action.ToParam()), actionController.List, admin.RouteConfig{
				PermissionMode: roles.Update,
				Resource:       res,
			})

			router.Put(path.Join("!notifications/:message_id", action.ToParam()), actionController.Action, admin.RouteConfig{
				PermissionMode: roles.Update,
				Resource:       res,
			})
		}
	}
}
