package notification

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/qor/qor"
	"github.com/qor/serializable_meta"
)

func Get(context *qor.Context) []Notification {
	return []Notification{}
}

func New(notification *Notification, context *qor.Context) {
	for _, channel := range registeredChannels {
		channel.Send(notification, context)
	}
}

type Notification struct {
	gorm.Model
	Kind    string
	Title   string
	Body    string
	ReadAt  *time.Time
	AckedAt *time.Time
	Actions Actions
}

func (notification *Notification) AddAction(action *Action) {
}

func (notification *Notification) LoadAction() []Action {
	return []Action{}
}

type Actions struct {
	Actions []Action
}

// (*Actions) Scanner, Valuer

type Action struct {
	Name string
	URL  string
	serializable_meta.SerializableMeta
}
