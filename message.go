package notification

import (
	"time"

	"github.com/jinzhu/gorm"
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
	AckedAt     *time.Time
}
