package misc

import (
	"github.com/cjungo/cjungo"
	"github.com/cjungo/cjungo/ext"
)

type ShellMessageToken string
type ShellMessage = ext.Message[ShellMessageToken]
type ShellMessageClient = ext.MessageClient[ShellMessageToken]
type ShellMessageController = ext.MessageController[ShellMessageToken]
type ShellMessageControllerProvide = ext.MessageControllerProvide[ShellMessageToken]
type ShellMessageControllerProviderConf = ext.MessageControllerProviderConf[ShellMessageToken]

func ProvideShellMessageController() ShellMessageControllerProvide {
	return ext.ProvideMessageController(&ShellMessageControllerProviderConf{
		TokenAccess: func(ctx cjungo.HttpContext) (ShellMessageToken, error) {
			return ShellMessageToken(ctx.GetReqID()), nil
		},
		OnRecv: func(
			controller *ShellMessageController,
			client *ShellMessageClient,
			msg *ShellMessage,
		) error {
			return nil
		},
	})
}
