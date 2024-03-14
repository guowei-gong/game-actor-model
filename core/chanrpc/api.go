package chanrpc

type IServer interface {
	Register(reqMsg any, f Handler)
}
