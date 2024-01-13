package router

func (c *container) RegisterPrescriptionRoute() {
	r := c.router.Group("/prescription")
	{
		r.Post("/add", c.AuthMiddleware(c.handler.prescription.Regist)...)
		r.Get("", c.AuthMiddleware(c.handler.prescription.GetList)...)
	}
}
