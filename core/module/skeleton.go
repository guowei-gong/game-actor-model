package module

import "guowei.com/mmo_game_actor/core/chanrpc"

// ISkeleton 骨头架
type ISkeleton interface {
	Run()
	Server() chanrpc.IServer
}

// skeleton 模块基础框架
type skeleton struct {
	chanSrv *chanrpc.Server
	chanCli *chanrpc.Client
}

func (s *skeleton) Server() chanrpc.IServer {
	return s.chanSrv
}

// Run 启动初始化
func (s *skeleton) Run() {
	for {
		select {
		case ackCtx := <-s.chanCli.ChanAck():
			s.chanCli.Exec(ackCtx)
		case reqCtx := <-s.chanSrv.ChanReq():
			s.chanSrv.Exec(reqCtx)
		}
	}
}

func NewSkeleton(clientChanSize, serverChanSize int) ISkeleton {
	return &skeleton{
		chanSrv: chanrpc.NewServer(serverChanSize),
		chanCli: chanrpc.NewClient(clientChanSize),
	}
}
