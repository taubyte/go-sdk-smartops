package function

import (
	"github.com/taubyte/go-sdk-smartops/common"
	httpFunction "github.com/taubyte/go-sdk-smartops/resource/function/http"
	p2pFunction "github.com/taubyte/go-sdk-smartops/resource/function/p2p"
	pubsubFunction "github.com/taubyte/go-sdk-smartops/resource/function/pubsub"
)

func Convert(r common.Resource) Function {
	return Function{r}
}

func (r Function) Http() (*httpFunction.Function, error) {
	return httpFunction.Convert(r)
}

func (r Function) P2P() (*p2pFunction.Function, error) {
	return p2pFunction.Convert(r)
}

func (r Function) PubSub() (*pubsubFunction.Function, error) {
	return pubsubFunction.Convert(r)
}
