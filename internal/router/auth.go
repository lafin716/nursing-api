package router

func (c *container) RegisterAuthRoute() {
	userRouter := c.router.Group("/auth")
	{
		userRouter.Post("/signup", c.handler.auth.SignUp)
		userRouter.Post("/signin", c.handler.auth.SignIn)
		userRouter.Post("/signout", c.AuthMiddleware(c.handler.auth.SignOut)...)
	}
}
