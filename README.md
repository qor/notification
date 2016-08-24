# QOR Notification

QOR Notification (WIP)

## Usage

```go
Notification := notification.New(&notification.Config{})

// Add to Admin
Admin.NewResource(Notification)

// Register Channels
Notification.RegisterChannel(ChannelInterface interface {
  Send(message *Message, context *qor.Context)
  GetNotifications(user interface{}, notification []*QorNotification, context *qor.Context) error
  GetNotification(user interface{}, notificationID string, context *qor.Context) (*QorNotification, error)
})

// Send Message
Notification.Send(message *Message, context *qor.Context)

notification.Get(context *qor.Context) []Notification

type Notification struct {
  Title string
  Body string
  serializable_meta.SerializableMeta
}

notification.RegisterChannel()
```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
