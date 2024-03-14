package lucy

import (
	"guowei.com/mmo_game_actor/core/chanrpc"
	"guowei.com/mmo_game_actor/core/module"
	"guowei.com/mmo_game_actor/example/iproto"
)

type Module struct {
	skeleton module.ISkeleton
}

func (m *Module) OnInit() error {
	m.ChanSrv().Register((*iproto.TestReq)(nil), m.onHandlerTestReq)
	return nil
}

func (m *Module) OnDestroy() {
	//TODO implement me
	panic("implement me")
}

func (m *Module) Run() {
	//TODO implement me
	panic("implement me")
}

func (m *Module) ChanSrv() chanrpc.IServer {
	return m.skeleton.Server()
}

func (m *Module) onHandlerTestReq(req *chanrpc.ReqCtx) {
	//TODO implement me
	panic("implement me")
}

func NewModule(clientChanSize, serverChanSize int) *Module {
	return &Module{
		skeleton: module.NewSkeleton(clientChanSize, serverChanSize),
	}
}
