package notification

import "github.com/qor/qor"

func (notification *Notification) RegisterChannel(channel ChannelInterface) {
	notification.Channels = append(notification.Channels, channel)
}

type ChannelInterface interface {
	Send(message *Message, context *qor.Context)
}
