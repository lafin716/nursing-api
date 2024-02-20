package router

func (c *container) RegisterTakeHistoryRoute() {
	r := c.router.Group("/takehistory")
	{
		r.Get("", c.AuthMiddleware(c.handler.takehistory.GetList)...)
		r.Get("/:id", c.AuthMiddleware(c.handler.takehistory.GetDetail)...)
	}
}
