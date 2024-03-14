package main

import (
	"guowei.com/mmo_game_actor/core"
	"guowei.com/mmo_game_actor/example/jimmy"
	"guowei.com/mmo_game_actor/example/lucy"
)

const (
	ServerChanSize = 100
	ClientChanSize = 100
)

func main() {
	// 创建 Actor 模块`Jimmy` 和 `Lucy` ,模块内需要实现 Actor 的消息传递方式
	lucyModule := lucy.NewModule(ClientChanSize, ServerChanSize)
	jimmyModule := jimmy.NewModule(ClientChanSize, ServerChanSize)

	// 启动项目
	core.DefaultApp().Start(lucyModule, jimmyModule)

	// 异步消息
}
