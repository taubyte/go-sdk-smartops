package messaging

import (
	"github.com/taubyte/go-sdk-smartops/common"
	pubsubMessaging "github.com/taubyte/go-sdk-smartops/resource/messaging/pubsub"
	websocketMessaging "github.com/taubyte/go-sdk-smartops/resource/messaging/websocket"
)

func Convert(r common.Resource) Messaging {
	return Messaging{r}
}

func (r Messaging) PubSub() (*pubsubMessaging.Messaging, error) {
	return pubsubMessaging.Convert(r)
}

func (r Messaging) WebSocket() (*websocketMessaging.Messaging, error) {
	return websocketMessaging.Convert(r)
}
