package router

func (c *container) RegisterAuthRoute() {
	r := c.router.Group("/auth")
	{
		r.Post("/signup", c.handler.auth.SignUp)
		r.Post("/signin", c.handler.auth.SignIn)
		r.Post("/signout", c.AuthMiddleware(c.handler.auth.SignOut)...)
		r.Post("/refresh", c.handler.auth.RefreshToken)
	}
}
