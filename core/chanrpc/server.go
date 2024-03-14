package chanrpc

// Server 代理服务器
type Server struct {
	chanReq  chan *ReqCtx
	handlers map[uint32]Handler
}

func (s *Server) Register(msg any, f Handler) {
	msgID := MsgID(msg)
	s.handlers[msgID] = f
}

func (s *Server) ChanReq() chan *ReqCtx {
	return s.chanReq
}

func NewServer(l int) *Server {
	s := new(Server)
	s.chanReq = make(chan *ReqCtx, l)
	s.handlers = make(map[uint32]Handler)
	return s
}

// Exec 执行
func (s *Server) Exec(reqCtx *ReqCtx) {
	s.exec(reqCtx)
}

func (s *Server) exec(reqCtx *ReqCtx) {
	handler := s.handlers[reqCtx.id]
	handler(reqCtx)
}
