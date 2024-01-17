package router

func (c *container) RegisterMainRoute() {
	c.router.Get("/", c.handler.main.HelloWorld)
}
