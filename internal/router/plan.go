package router

func (c *container) RegisterPlanRoute() {
	r := c.router.Group("/plan")
	{
		r.Get("/today", c.AuthMiddleware(c.handler.plan.Today)...)
		r.Get("/summary", c.AuthMiddleware(c.handler.plan.Summary)...)
		r.Post("/take", c.AuthMiddleware(c.handler.plan.TakePlan)...)
		r.Post("/take/pill", c.AuthMiddleware(c.handler.plan.TakePill)...)
		r.Post("/memo", c.AuthMiddleware(c.handler.plan.UpdateMemo)...)
	}
}
