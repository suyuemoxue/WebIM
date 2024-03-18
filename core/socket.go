package core

import (
	"CopyQQ/service/ws"
)

func InitSocket() {
	go ws.Manager.Start()
}
