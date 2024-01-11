package router

func (c *container) RegisterUserRoute() {
	userRouter := c.router.Group("/user")
	{
		userRouter.Get("/me", c.AuthMiddleware(c.handler.user.Me)...)
	}
}
