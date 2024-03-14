package chanrpc

type Client struct {
	chanAsyncAck chan *AckCtx
}

func NewClient(l int) *Client {
	return &Client{
		chanAsyncAck: make(chan *AckCtx, l),
	}
}

func (c *Client) ChanAck() chan *AckCtx {
	return c.chanAsyncAck
}

func (c *Client) Exec(ackCtx *AckCtx) {
	c.exec(ackCtx)
}

func (c *Client) exec(ackCtx *AckCtx) {

}
