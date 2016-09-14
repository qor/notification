package database

import (
	"fmt"
	"reflect"

	"github.com/qor/notification"
	"github.com/qor/qor"
)

type Config struct {
}

func New(config *Config) *Database {
	return &Database{Config: config}
}

type Database struct {
	Config *Config
}

func (database *Database) Send(message *notification.Message, context *qor.Context) error {
	notice := notification.QorNotification{
		From:        database.getUserID(message.From, context),
		To:          database.getUserID(message.To, context),
		Title:       message.Title,
		Body:        message.Body,
		MessageType: message.MessageType,
	}

	return context.GetDB().Save(&notice).Error
}

func (database *Database) GetNotifications(user interface{}, notifications *[]*notification.QorNotification, context *qor.Context) error {
	var newNotifications []*notification.QorNotification
	var to = database.getUserID(user, context)
	var db = context.GetDB()

	err := db.Find(&newNotifications, fmt.Sprintf("%v = ?", db.Dialect().Quote("to")), to).Error
	*notifications = append(*notifications, newNotifications...)
	return err
}

func (database *Database) GetNotification(user interface{}, notificationID string, context *qor.Context) (*notification.QorNotification, error) {
	var (
		notice notification.QorNotification
		to     = database.getUserID(user, context)
		db     = context.GetDB()
	)

	err := db.First(&notice, fmt.Sprintf("%v = ? AND %v = ?", db.Dialect().Quote("to"), db.Dialect().Quote("id")), to, notificationID).Error
	return &notice, err
}

func (database *Database) getUserID(user interface{}, context *qor.Context) string {
	var (
		userID string
		scope  = context.GetDB().NewScope(user)
	)

	if scope.IndirectValue().Kind() == reflect.Struct {
		userID = fmt.Sprint(scope.PrimaryKeyValue())
	} else {
		userID = fmt.Sprint(user)
	}

	return userID
}
