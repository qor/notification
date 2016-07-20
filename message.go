package notification

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	Title       string
	Body        string
	MessageType string
	AckedAt     *time.Time
}
