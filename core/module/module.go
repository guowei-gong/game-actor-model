package module

import "guowei.com/mmo_game_actor/core/chanrpc"

type Module interface {
	OnInit() error            // 初始化
	OnDestroy()               // 销毁
	Run()                     // 启动
	ChanSrv() chanrpc.IServer // 消息通道(信箱)
}
