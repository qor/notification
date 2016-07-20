package notification

import "github.com/qor/qor"

type Notification struct {
	Config   *Config
	Channels []ChannelInterface
}

func New(config *Config) *Notification {
	return &Notification{Config: config}
}

func (notification *Notification) RegisterChannel(channel ChannelInterface) {
	notification.Channels = append(notification.Channels, channel)
}

func (notification *Notification) Send(message *Message, context *qor.Context) {
	for _, channel := range notification.Channels {
		channel.Send(message, context)
	}
}
