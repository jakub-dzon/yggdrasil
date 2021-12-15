package transport

import (
	"github.com/redhatinsights/yggdrasil"
)

type CommandHandler func(command []byte, t Transport)
type DataHandler func(data []byte)

type Transport interface {
	Start() error
	SendData(data yggdrasil.Data) (*yggdrasil.APIresponse, error)
	SendControl(ctrlMsg interface{}) (*yggdrasil.APIresponse, error)
	Disconnect(quiesce uint)
}
