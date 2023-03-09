package pubsub

import (
	"context"
)

type NatsPubSub interface {
	Publish(ctx context.Context, channel string, data *Message) error
	Subscribe(ctx context.Context, channel string) (ch <-chan *Message, close func())
}
