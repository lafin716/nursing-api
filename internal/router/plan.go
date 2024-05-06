package router

func (c *container) RegisterPlanRoute() {
	r := c.router.Group("/plan")
	{
		r.Get("/", c.AuthMiddleware(c.handler.plan.GetByDate)...)
		r.Get("/month", c.AuthMiddleware(c.handler.plan.GetByMonth)...)
		r.Post("/", c.AuthMiddleware(c.handler.plan.Add)...)
		r.Delete("/:id", c.AuthMiddleware(c.handler.plan.Delete)...)
		r.Get("/summary", c.AuthMiddleware(c.handler.plan.Summary)...)
		r.Post("/take", c.AuthMiddleware(c.handler.plan.Take)...)
		r.Post("/take/pill", c.AuthMiddleware(c.handler.plan.PillToggle)...)
		r.Post("/memo", c.AuthMiddleware(c.handler.plan.UpdateMemo)...)
	}
}
