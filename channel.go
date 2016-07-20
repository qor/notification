package notification

import "github.com/qor/qor"

type ChannelInterface interface {
	Send(message *Message, context *qor.Context)
}
