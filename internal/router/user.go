package router

func (c *container) RegisterUserRoute() {
	r := c.router.Group("/user")
	{
		r.Get("/me", c.AuthMiddleware(c.handler.user.Me)...)
	}
}
