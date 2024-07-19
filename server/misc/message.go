package misc

import (
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/ext"
)

type MessageToken string
type MessageController = ext.MessageController[MessageToken]
type MessageClient = ext.MessageClient[MessageToken]
type Message = ext.Message[MessageToken]
type MessageControllerProvide = ext.MessageControllerProvide[MessageToken]
type MessageControllerProviderConf = ext.MessageControllerProviderConf[MessageToken]

func ProvideMessageController() MessageControllerProvide {
	return ext.ProvideMessageController(&MessageControllerProviderConf{
		TokenAccess: func(ctx cjungo.HttpContext) (MessageToken, error) {
			return MessageToken(ctx.GetReqID()), nil
		},
	})
}
