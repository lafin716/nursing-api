package router

func (c *container) RegisterTimeZoneRoute() {
	r := c.router.Group("/timezone")
	{
		r.Get("", c.AuthMiddleware(c.handler.timezone.GetList)...)
		r.Post("", c.AuthMiddleware(c.handler.timezone.Create)...)
		r.Patch("", c.AuthMiddleware(c.handler.timezone.Update)...)
		r.Delete("/:id", c.AuthMiddleware(c.handler.timezone.Delete)...)
	}
}
