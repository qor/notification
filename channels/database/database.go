package database

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	"github.com/qor/notification"
	"github.com/qor/qor"
)

type Config struct {
	DB *gorm.DB
}

func New(config *Config) *Database {
	if config.DB != nil {
		config.DB.AutoMigrate(&notification.QorNotification{})
	} else {
		fmt.Println("Need to have gorm DB in the configuration in order to run migrations")
	}
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

func (database *Database) GetNotifications(user interface{}, results *notification.NotificationsResult, _ *notification.Notification, context *qor.Context) error {
	var to = database.getUserID(user, context)
	var db = context.GetDB()

	// get unresolved notifications
	if err := db.Find(&results.Notifications, fmt.Sprintf("%v = ? AND %v IS NULL", db.Dialect().Quote("to"), db.Dialect().Quote("resolved_at")), to).Error; err != nil {
		return err
	}

	// get resolved notifications
	return db.Find(&results.Resolved, fmt.Sprintf("%v = ? AND %v IS NOT NULL", db.Dialect().Quote("to"), db.Dialect().Quote("resolved_at")), to).Error
}

func (database *Database) GetNotification(user interface{}, notificationID string, _ *notification.Notification, context *qor.Context) (*notification.QorNotification, error) {
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
