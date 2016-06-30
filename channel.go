package notification

import "github.com/qor/qor"

type ChannelInterface interface {
	Send(notification *Notification, context *qor.Context)
}

var registeredChannels []ChannelInterface

func RegisterChannel(channel ChannelInterface) {
	registeredChannels = append(registeredChannels, channel)
}
