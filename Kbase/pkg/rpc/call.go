package rpc

// Call represents an active rpc
type Call struct {
	Seq           uint64
	ServiceMethod string      // format "<serviece>.<function>"
	Args          interface{} // arguments for the function
	Reply         interface{} // reply from the function
	Error         error       //if error occurs,it will be set
	Done          chan *Call  // Storbes when call is complete
}

func (c *Call) done() {
	c.Done <- c
}
