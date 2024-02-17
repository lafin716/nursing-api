package router

func (c *container) RegisterPrescriptionRoute() {
	r := c.router.Group("/prescription")
	{
		r.Get("", c.AuthMiddleware(c.handler.prescription.GetList)...)
		r.Post("", c.AuthMiddleware(c.handler.prescription.Register)...)
		r.Patch("", c.AuthMiddleware(c.handler.prescription.Update)...)
		r.Delete("/:id", c.AuthMiddleware(c.handler.prescription.Delete)...)
		r.Post("/item", c.AuthMiddleware(c.handler.prescription.AddItem)...)
		r.Patch("/item", c.AuthMiddleware(c.handler.prescription.UpdateItem)...)
		r.Delete("/item/:id", c.AuthMiddleware(c.handler.prescription.DeleteItem)...)
	}
}
