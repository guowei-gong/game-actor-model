package jimmy

import (
	"fmt"
	"guowei.com/mmo_game_actor/core/chanrpc"
	"guowei.com/mmo_game_actor/core/module"
	"guowei.com/mmo_game_actor/example/iproto"
)

type Module struct {
	skeleton module.ISkeleton
}

func (m *Module) OnInit() error {
	m.ChanSrv().Register((*iproto.TestReq)(nil), m.onHandleTestReq)
	return nil
}

func (m *Module) OnDestroy() {}

func (m *Module) Run() {
	m.skeleton.Run()
}

func (m *Module) ChanSrv() chanrpc.IServer {
	return m.skeleton.Server()
}

func NewModule(clientChanSize, serverChanSize int) *Module {
	return &Module{
		skeleton: module.NewSkeleton(clientChanSize, serverChanSize),
	}
}

func (m *Module) onHandleTestReq(req *chanrpc.ReqCtx) {
	fmt.Println("Jimmy receive msg")
}
